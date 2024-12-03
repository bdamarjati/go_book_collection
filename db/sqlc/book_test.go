package sqlc

import (
	"context"
	"database/sql"
	"testing"

	"github.com/bdamarjati/go_book_collection/util"
	"github.com/stretchr/testify/require"
)

func createRandomBook(t *testing.T) Book {
	collection := createRandomCollection(t)
	require.NotEmpty(t, collection)

	arg := CreateBookParams{
		CollectionID:  collection.CollectionID,
		Title:         util.RandomText(16),
		Author:        util.RandomText(8),
		Isbn:          util.RandomText(16),
		Language:      util.RandomLanguage(),
		YearPublished: util.RandomYear(),
	}

	book, err := testQueries.CreateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book)

	rowsAffected, _ := book.RowsAffected()
	require.NotZero(t, rowsAffected)

	lastId, _ := book.LastInsertId()
	require.NotZero(t, lastId)

	return Book{
		BookID:        int32(lastId),
		CollectionID:  collection.CollectionID,
		Title:         arg.Title,
		Author:        arg.Author,
		Language:      arg.Language,
		YearPublished: arg.YearPublished,
		Isbn:          arg.Isbn,
	}
}

func getBookById(id int32) (Book, error) {
	book, err := testQueries.GetBook(context.Background(), id)
	if err != nil {
		return Book{}, err
	}
	return book, nil
}

func TestCreateBook(t *testing.T) {
	createRandomBook(t)
}

func TestGetBook(t *testing.T) {
	book1 := createRandomBook(t)
	book2, err := getBookById(book1.BookID)

	require.NoError(t, err)
	require.NotEmpty(t, book2)

	require.Equal(t, book1.BookID, book2.BookID)
	require.Equal(t, book1.CollectionID, book2.CollectionID)
	require.Equal(t, book1.Title, book2.Title)
	require.Equal(t, book1.Author, book2.Author)
	require.Equal(t, book1.Language, book2.Language)
	require.Equal(t, book1.YearPublished, book2.YearPublished)
	require.Equal(t, book1.Isbn, book2.Isbn)
	require.NotEmpty(t, book2.CreatedAt)
}

func TestUpdateBook(t *testing.T) {
	book1 := createRandomBook(t)

	arg := UpdateBookParams{
		BookID:        book1.BookID,
		Title:         util.RandomText(16),
		Author:        util.RandomText(8),
		Isbn:          util.RandomText(16),
		Language:      util.RandomLanguage(),
		YearPublished: util.RandomYear(),
	}

	err := testQueries.UpdateBook(context.Background(), arg)
	require.NoError(t, err)

	book2, err := getBookById(arg.BookID)
	require.NoError(t, err)
	require.NotEmpty(t, book2)

	require.Equal(t, book1.BookID, book2.BookID)
	require.Equal(t, arg.Title, book2.Title)
	require.Equal(t, arg.Author, book2.Author)
	require.Equal(t, arg.Language, book2.Language)
	require.Equal(t, arg.YearPublished, book2.YearPublished)
	require.Equal(t, arg.Isbn, book2.Isbn)
}

func TestDeleteBook(t *testing.T) {
	book1 := createRandomBook(t)
	err := testQueries.DeleteBook(context.Background(), book1.BookID)
	require.NoError(t, err)

	book2, err := getBookById(book1.BookID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, book2)
}
