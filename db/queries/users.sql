-- name: ListUsers :many
SELECT * FROM users
ORDER BY id DESC;

-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (name, mobile, role)
VALUES (?, ?, ?)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET name = ?, mobile = ?, role = ?
WHERE id = ?
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;