-- name: GetArticleByID :one
SELECT id, title, content, created_at
FROM articles
WHERE id = $1;

