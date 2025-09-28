package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/yeungon/gossr/internal/module/categories/business"
	"github.com/yeungon/gossr/internal/module/categories/domain"
	"github.com/yeungon/gossr/internal/module/categories/sqlc"
)

type CategoryPostgres struct {
	q *sqlc.Queries
}

func NewCategoryPostgres(q *sqlc.Queries) business.CategoryRepository {
	return &CategoryPostgres{q: q}
}

func toDomain(c sqlc.Category) domain.Category {
	return domain.Category{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt.Time,
	}
}

func (r *CategoryPostgres) GetByID(id int64) (*domain.Category, error) {
	row, err := r.q.GetCategoryByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	cat := toDomain(row)
	return &cat, nil
}

func (r *CategoryPostgres) ListAll() ([]domain.Category, error) {
	rows, err := r.q.ListCategories(context.Background())
	if err != nil {
		return nil, err
	}
	cats := make([]domain.Category, len(rows))
	for i, row := range rows {
		cats[i] = toDomain(row)
	}
	return cats, nil
}

func (r *CategoryPostgres) Create(c domain.Category) (*domain.Category, error) {
	params := sqlc.InsertCategoryParams{
		Name: c.Name,
		CreatedAt: pgtype.Timestamptz{
			Time:  c.CreatedAt,
			Valid: true,
		},
	}

	created, err := r.q.InsertCategory(context.Background(), params)
	if err != nil {
		return nil, err
	}
	cat := toDomain(created)
	return &cat, nil
}
