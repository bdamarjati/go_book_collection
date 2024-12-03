-- name: CreateUser :execresult
INSERT INTO users (username, role) VALUES (?, ?);

-- name: GetUser :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id LIMIT ? OFFSET ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;

-- name: UpdateUser :exec
UPDATE users SET username = ?, role = ? WHERE id = ?;
