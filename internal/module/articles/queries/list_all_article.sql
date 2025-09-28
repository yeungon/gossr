
-- name: ListArticles :many
SELECT id, title, content, created_at
FROM articles
ORDER BY id ASC;