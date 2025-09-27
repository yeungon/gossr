package storage

import (
	"context"

	"github.com/yeungon/gossr/internal/infra/db/sqlc"
	"github.com/yeungon/gossr/internal/module/order/business"
	"github.com/yeungon/gossr/internal/module/order/domain"
)

type PostgresRepo struct {
	q *sqlc.Queries
}

func NewPostgresRepo(q *sqlc.Queries) business.OrderRepository {
	return &PostgresRepo{q: q}
}

func (r *PostgresRepo) Save(order *domain.Order) error {
	row, err := r.q.CreateOrder(context.Background(), int32(order.ItemID), int32(order.Qty))
	if err != nil {
		return err
	}
	order.ID = int(row.ID)
	order.ItemID = int(row.ItemID)
	order.Qty = int(row.Qty)
	return nil
}
