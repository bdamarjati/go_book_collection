-- name: CreateCollection :execresult
INSERT INTO collections (user, name, status) VALUES (?, ?, ?);

-- name: GetCollection :one
SELECT * FROM collections WHERE collection_id = ? LIMIT 1;

-- name: ListCollections :many
SELECT * FROM collections ORDER BY collection_id LIMIT ? OFFSET ?;

-- name: ListCollectionsByUser :many
SELECT * FROM collections WHERE user = ? ORDER BY collection_id LIMIT ? OFFSET ?;

-- name: UpdateCollection :exec
UPDATE collections SET name = ?, status = ? WHERE collection_id = ?;

-- name: DeleteCollection :exec
DELETE FROM collections WHERE collection_id = ?;
