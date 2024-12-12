-- name: CreateUser :execresult
INSERT INTO users (username, password, role) VALUES (?, ?, ?);

-- name: GetUser :one
SELECT * FROM users WHERE username = ? LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users WHERE username = ?;

-- name: UpdateUser :exec
UPDATE users SET password = ?, role = ? WHERE username = ?;
