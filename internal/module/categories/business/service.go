package business

import (
	"github.com/yeungon/gossr/internal/module/categories/domain"
)

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(r CategoryRepository) *CategoryService {
	return &CategoryService{repo: r}
}

func (s *CategoryService) GetCategory(id int64) (*domain.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) ListCategories() ([]domain.Category, error) {
	return s.repo.ListAll()
}

func (s *CategoryService) CreateCategory(c domain.Category) (*domain.Category, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return s.repo.Create(c)
}
