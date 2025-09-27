package business

import (
	"errors"
	"log"

	"github.com/yeungon/gossr/internal/module/articles/domain"
)

var ErrInvalidName = errors.New("invalid item name")

type ArticleService struct {
	repo   ArticleRepository
	logger *log.Logger
}

func NewArticleService(r ArticleRepository, logger *log.Logger) *ArticleService {
	return &ArticleService{repo: r, logger: logger}
}

func (s *ArticleService) GetArticle(id int64) (*domain.Article, error) {
	s.logger.Println("Fetching article", id)
	return s.repo.GetByID(id)
}
