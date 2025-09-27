package business

import (
	"errors"

	"github.com/yeungon/gossr/internal/module/categories/domain"
)

var ErrInvalidOrder = errors.New("invalid order")

type OrderRepository interface {
	Save(order *domain.Order) error
}
