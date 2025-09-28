
-- name: InsertArticle :one
INSERT INTO articles (title, content, created_at)
VALUES ($1, $2, $3)
RETURNING id, title, content, created_at;
