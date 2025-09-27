package business

import "github.com/yeungon/gossr/internal/module/articles/domain"

type ArticleRepository interface {
	GetByID(id int64) (*domain.Article, error)
}
