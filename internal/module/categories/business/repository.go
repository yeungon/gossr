package business

import "github.com/yeungon/gossr/internal/module/categories/domain"

type CategoryRepository interface {
	GetByID(id int64) (*domain.Category, error)
	ListAll() ([]domain.Category, error)
	Create(c domain.Category) (*domain.Category, error)
}
