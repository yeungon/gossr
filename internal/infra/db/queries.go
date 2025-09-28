package db

import (
	articlesqlc "github.com/yeungon/gossr/internal/module/articles/sqlc"
	categoriesqlc "github.com/yeungon/gossr/internal/module/categories/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Queries struct {
	Articles   *articlesqlc.Queries
	Categories *categoriesqlc.Queries
}

func NewQueries(pool *pgxpool.Pool) *Queries {
	return &Queries{
		Articles:   articlesqlc.New(pool),
		Categories: categoriesqlc.New(pool),
	}
}
