package business

import (
	"errors"

	itembiz "github.com/yeungon/gossr/internal/module/categories/business"
	"github.com/yeungon/gossr/internal/module/order/domain"
)

type Service struct {
	repo    OrderRepository
	itemSvc *itembiz.Service
}

func NewService(r OrderRepository, itemSvc *itembiz.Service) *Service {
	return &Service{repo: r, itemSvc: itemSvc}
}

func (s *Service) CreateOrder(order *domain.Order) error {
	if order.Qty <= 0 {
		return ErrInvalidOrder
	}
	_, err := s.itemSvc.GetItem(order.ItemID)
	if err != nil {
		return errors.New("item not found")
	}
	return s.repo.Save(order)
}
