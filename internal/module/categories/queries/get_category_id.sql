-- name: GetCategoryByID :one
SELECT id, name, created_at
FROM categories
WHERE id = $1;