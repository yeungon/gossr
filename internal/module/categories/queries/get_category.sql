-- name: GetCategoryByID :one
SELECT id, name, created_at
FROM categories
WHERE id = $1;

-- name: ListCategories :many
SELECT id, name, created_at
FROM categories
ORDER BY id;

-- name: InsertCategory :one
INSERT INTO categories (name, created_at)
VALUES ($1, $2)
RETURNING id, name, created_at;
