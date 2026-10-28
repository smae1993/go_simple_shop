-- name: ListPublishedPosts :many
SELECT * FROM posts
WHERE is_published = 1
ORDER BY created_at DESC
LIMIT ?;

-- name: ListAllPostsAdmin :many
SELECT * FROM posts
ORDER BY created_at DESC;

-- name: GetPostBySlug :one
SELECT * FROM posts
WHERE slug = ? AND is_published = 1 LIMIT 1;

-- name: GetPostByID :one
SELECT * FROM posts
WHERE id = ? LIMIT 1;

-- name: CreatePost :one
INSERT INTO posts (title, slug, summary, content, is_published)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdatePost :one
UPDATE posts
SET title = ?, slug = ?, summary = ?, content = ?, is_published = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = ?;