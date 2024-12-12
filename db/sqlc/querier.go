// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateBook(ctx context.Context, arg CreateBookParams) (sql.Result, error)
	CreateCollection(ctx context.Context, arg CreateCollectionParams) (sql.Result, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error)
	DeleteBook(ctx context.Context, bookID int32) error
	DeleteCollection(ctx context.Context, collectionID int32) error
	DeleteUser(ctx context.Context, username string) error
	GetBook(ctx context.Context, bookID int32) (Book, error)
	GetCollection(ctx context.Context, collectionID int32) (Collection, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListBooks(ctx context.Context, arg ListBooksParams) ([]Book, error)
	ListBooksByCollection(ctx context.Context, arg ListBooksByCollectionParams) ([]Book, error)
	ListBooksByUser(ctx context.Context, arg ListBooksByUserParams) ([]Book, error)
	ListCollections(ctx context.Context, arg ListCollectionsParams) ([]Collection, error)
	ListCollectionsByUser(ctx context.Context, arg ListCollectionsByUserParams) ([]Collection, error)
	UpdateBook(ctx context.Context, arg UpdateBookParams) error
	UpdateCollection(ctx context.Context, arg UpdateCollectionParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
