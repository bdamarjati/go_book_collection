-- name: CreateBook :execresult
INSERT INTO books (collection_id, title, author, language, year_published, ISBN) VALUES (?, ?, ?, ?, ?, ?);

-- name: GetBook :one
SELECT * FROM books WHERE book_id = ? LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books ORDER BY book_id LIMIT ? OFFSET ?;

-- name: ListBooksByCollection :many
SELECT * FROM books WHERE collection_id = ? ORDER BY book_id LIMIT ? OFFSET ?;

-- name: ListBooksByUser :many
SELECT books.* FROM ((books 
INNER JOIN collections ON books.collection_id = collections.collection_id)
INNER JOIN users ON collections.user_id = users.id) 
WHERE users.id = ? ORDER BY books.book_id LIMIT ? OFFSET ?;

-- name: UpdateBook :exec
UPDATE books SET title = ?, author = ?, language = ?, year_published = ?, ISBN = ? WHERE book_id = ?;

-- name: DeleteBook :exec
DELETE FROM books WHERE book_id = ?;
