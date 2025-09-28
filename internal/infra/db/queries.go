package db

import (
	articlesqlc "github.com/yeungon/gossr/internal/module/articles/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Queries struct {
	Articles *articlesqlc.Queries
}

func NewQueries(pool *pgxpool.Pool) *Queries {
	return &Queries{
		Articles: articlesqlc.New(pool),
	}
}
