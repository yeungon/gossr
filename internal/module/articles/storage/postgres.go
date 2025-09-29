package storage

import (
	"context"

	"github.com/yeungon/gossr/internal/module/articles/business"
	"github.com/yeungon/gossr/internal/module/articles/domain"
	articles "github.com/yeungon/gossr/internal/module/articles/mapper"
	"github.com/yeungon/gossr/internal/module/articles/sqlc"
)

// ArticlePostgres implements business.ArticleRepository using sqlc + Postgres.
type ArticlePostgres struct {
	q *sqlc.Queries
}

// NewArticlePostgres returns a Postgres-backed repository.
func NewArticlePostgres(q *sqlc.Queries) business.ArticleRepository {
	return &ArticlePostgres{q: q}
}

// --- Repository methods ---

func (r *ArticlePostgres) GetByID(id int64) (*domain.Article, error) {
	row, err := r.q.GetArticleByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	article := articles.FromSQLCToDomain(row)
	return &article, nil
}

// Example: List all articles
func (r *ArticlePostgres) ListAll() ([]domain.Article, error) {
	rows, err := r.q.ListArticles(context.Background())
	if err != nil {
		return nil, err
	}
	articleList := make([]domain.Article, len(rows))
	for i, row := range rows {
		articleList[i] = articles.FromSQLCToDomain(row)
	}
	return articleList, nil
}

func (r *ArticlePostgres) Create(a domain.Article) (*domain.Article, error) {
	params := articles.ToSQLCInsertParams(a)

	created, err := r.q.InsertArticle(context.Background(), params)
	if err != nil {
		return nil, err
	}
	article := articles.FromSQLCToDomain(created)
	return &article, nil
}
