package business

import (
	"errors"

	"github.com/yeungon/gossr/internal/module/articles/domain"
)

var ErrInvalidName = errors.New("invalid item name")

type ArticleService struct {
	repository ArticleRepository
}

func NewArticleService(r ArticleRepository) *ArticleService {
	return &ArticleService{repository: r}
}

func (s *ArticleService) GetArticle(id int64) (*domain.Article, error) {
	return s.repository.GetByID(id)
}
