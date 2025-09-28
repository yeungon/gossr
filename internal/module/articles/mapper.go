package articles

import (
	"github.com/yeungon/gossr/internal/module/articles/domain"
	"github.com/yeungon/gossr/internal/module/articles/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

// FromSQLC converts a sqlc.Article to a domain.Article
func FromSQLC(a sqlc.Article) domain.Article {
	return domain.Article{
		ID:        a.ID,
		Title:     a.Title,
		Content:   a.Content,
		CreatedAt: a.CreatedAt.Time, // pgtype -> time.Time
	}
}

// ToSQLC converts a domain.Article to a sqlc.Article
func ToSQLC(a domain.Article) sqlc.Article {
	return sqlc.Article{
		ID:      a.ID,
		Title:   a.Title,
		Content: a.Content,
		// pgtype.Timestamptz has Valid + Time
		CreatedAt: pgtype.Timestamptz{
			Time:  a.CreatedAt,
			Valid: true,
		},
	}
}

// Batch helpers for slices
func FromSQLCList(list []sqlc.Article) []domain.Article {
	out := make([]domain.Article, len(list))
	for i, a := range list {
		out[i] = FromSQLC(a)
	}
	return out
}
