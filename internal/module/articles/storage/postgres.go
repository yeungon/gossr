package storage

import (
	"context"

	"github.com/yeungon/gossr/internal/module/articles/business"
	"github.com/yeungon/gossr/internal/module/articles/domain"
	"github.com/yeungon/gossr/internal/module/articles/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

// ArticlePostgres implements business.ArticleRepository using sqlc + Postgres.
type ArticlePostgres struct {
	q *sqlc.Queries
}

// NewArticlePostgres returns a Postgres-backed repository.
func NewArticlePostgres(q *sqlc.Queries) business.ArticleRepository {
	return &ArticlePostgres{q: q}
}

// --- Mapper helpers (kept private to this package) ---

func toDomain(a sqlc.Article) domain.Article {
	return domain.Article{
		ID:        a.ID,
		Title:     a.Title,
		Content:   a.Content,
		CreatedAt: a.CreatedAt.Time, // pgtype -> time.Time
	}
}

// --- Repository methods ---

func (r *ArticlePostgres) GetByID(id int64) (*domain.Article, error) {
	row, err := r.q.GetArticleByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	article := toDomain(row)
	return &article, nil
}

// Example: List all articles
func (r *ArticlePostgres) ListAll() ([]domain.Article, error) {
	rows, err := r.q.ListArticles(context.Background())
	if err != nil {
		return nil, err
	}
	articles := make([]domain.Article, len(rows))
	for i, row := range rows {
		articles[i] = toDomain(row)
	}
	return articles, nil
}

func (r *ArticlePostgres) Create(a domain.Article) (*domain.Article, error) {
	params := sqlc.InsertArticleParams{
		Title:   a.Title,
		Content: a.Content,
		CreatedAt: pgtype.Timestamptz{
			Time:  a.CreatedAt,
			Valid: true,
		},
	}

	created, err := r.q.InsertArticle(context.Background(), params)
	if err != nil {
		return nil, err
	}
	article := toDomain(created)
	return &article, nil
}
