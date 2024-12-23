// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: book.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createBook = `-- name: CreateBook :execresult
INSERT INTO books (collection_id, title, author, language, year_published, ISBN) VALUES (?, ?, ?, ?, ?, ?)
`

type CreateBookParams struct {
	CollectionID  int32          `json:"collection_id"`
	Title         sql.NullString `json:"title"`
	Author        sql.NullString `json:"author"`
	Language      sql.NullString `json:"language"`
	YearPublished sql.NullInt32  `json:"year_published"`
	Isbn          sql.NullString `json:"isbn"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createBook,
		arg.CollectionID,
		arg.Title,
		arg.Author,
		arg.Language,
		arg.YearPublished,
		arg.Isbn,
	)
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books WHERE book_id = ?
`

func (q *Queries) DeleteBook(ctx context.Context, bookID int32) error {
	_, err := q.db.ExecContext(ctx, deleteBook, bookID)
	return err
}

const getBook = `-- name: GetBook :one
SELECT book_id, collection_id, title, author, language, year_published, isbn, created_at FROM books WHERE book_id = ? LIMIT 1
`

func (q *Queries) GetBook(ctx context.Context, bookID int32) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, bookID)
	var i Book
	err := row.Scan(
		&i.BookID,
		&i.CollectionID,
		&i.Title,
		&i.Author,
		&i.Language,
		&i.YearPublished,
		&i.Isbn,
		&i.CreatedAt,
	)
	return i, err
}

const listBooks = `-- name: ListBooks :many
SELECT book_id, collection_id, title, author, language, year_published, isbn, created_at FROM books ORDER BY book_id LIMIT ? OFFSET ?
`

type ListBooksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBooks(ctx context.Context, arg ListBooksParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.CollectionID,
			&i.Title,
			&i.Author,
			&i.Language,
			&i.YearPublished,
			&i.Isbn,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBooksByCollection = `-- name: ListBooksByCollection :many
SELECT book_id, collection_id, title, author, language, year_published, isbn, created_at FROM books WHERE collection_id = ? ORDER BY book_id LIMIT ? OFFSET ?
`

type ListBooksByCollectionParams struct {
	CollectionID int32 `json:"collection_id"`
	Limit        int32 `json:"limit"`
	Offset       int32 `json:"offset"`
}

func (q *Queries) ListBooksByCollection(ctx context.Context, arg ListBooksByCollectionParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooksByCollection, arg.CollectionID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.CollectionID,
			&i.Title,
			&i.Author,
			&i.Language,
			&i.YearPublished,
			&i.Isbn,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBooksByUser = `-- name: ListBooksByUser :many
SELECT books.book_id, books.collection_id, books.title, books.author, books.language, books.year_published, books.isbn, books.created_at FROM ((books 
INNER JOIN collections ON books.collection_id = collections.collection_id)
INNER JOIN users ON collections.user = users.username) 
WHERE users.username = ? ORDER BY books.book_id LIMIT ? OFFSET ?
`

type ListBooksByUserParams struct {
	Username string `json:"username"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

func (q *Queries) ListBooksByUser(ctx context.Context, arg ListBooksByUserParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooksByUser, arg.Username, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.CollectionID,
			&i.Title,
			&i.Author,
			&i.Language,
			&i.YearPublished,
			&i.Isbn,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBook = `-- name: UpdateBook :exec
UPDATE books SET title = ?, author = ?, language = ?, year_published = ?, ISBN = ? WHERE book_id = ?
`

type UpdateBookParams struct {
	Title         sql.NullString `json:"title"`
	Author        sql.NullString `json:"author"`
	Language      sql.NullString `json:"language"`
	YearPublished sql.NullInt32  `json:"year_published"`
	Isbn          sql.NullString `json:"isbn"`
	BookID        int32          `json:"book_id"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	_, err := q.db.ExecContext(ctx, updateBook,
		arg.Title,
		arg.Author,
		arg.Language,
		arg.YearPublished,
		arg.Isbn,
		arg.BookID,
	)
	return err
}
