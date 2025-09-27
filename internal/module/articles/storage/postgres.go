package storage

import (
	"context"

	"github.com/yeungon/gossr/internal/module/articles/business"
	"github.com/yeungon/gossr/internal/module/articles/domain"
	"github.com/yeungon/gossr/internal/module/articles/sqlc"
)

type ArticlePostgres struct {
	q *sqlc.Queries
}

func NewArticlePostgres(q *sqlc.Queries) business.ArticleRepository {
	return &ArticlePostgres{q: q}
}

func (r *ArticlePostgres) GetByID(id int64) (*domain.Article, error) {
	row, err := r.q.GetArticleByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &domain.Article{
		ID:        row.ID,
		Title:     row.Title,
		Content:   row.Content,
		CreatedAt: row.CreatedAt,
	}, nil
}
