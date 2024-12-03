package sqlc

import (
	"context"
	"database/sql"
	"testing"

	"github.com/bdamarjati/go_book_collection/util"
	"github.com/stretchr/testify/require"
)

func createRandomCollection(t *testing.T) Collection {
	user := createRandomUser(t)
	require.NotEmpty(t, user)

	arg := CreateCollectionParams{
		UserID: user.ID,
		Name:   util.RandomCollection(),
		Status: util.RandomCollectionStatus(),
	}

	collection, err := testQueries.CreateCollection(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, collection)

	rowsAffected, err := collection.RowsAffected()
	require.NoError(t, err)
	require.NotZero(t, rowsAffected)

	lastId, err := collection.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, lastId)

	return Collection{
		CollectionID: int32(lastId),
		UserID:       user.ID,
		Name:         arg.Name,
		Status:       arg.Status,
	}
}

func getCollectionById(id int32) (Collection, error) {
	collection, err := testQueries.GetCollection(context.Background(), id)
	if err != nil {
		return Collection{}, err
	}
	return collection, nil
}

func TestCreateCollection(t *testing.T) {
	createRandomCollection(t)
}

func TestGetCollection(t *testing.T) {
	collection1 := createRandomCollection(t)
	collection2, err := getCollectionById(collection1.CollectionID)

	require.NoError(t, err)
	require.NotEmpty(t, collection2)

	require.Equal(t, collection1.CollectionID, collection2.CollectionID)
	require.Equal(t, collection1.Name, collection2.Name)
	require.Equal(t, collection1.Status, collection2.Status)
	require.NotEmpty(t, collection2.CreatedAt)
}

func TestUpdateCollection(t *testing.T) {
	collection1 := createRandomCollection(t)

	arg := UpdateCollectionParams{
		CollectionID: collection1.CollectionID,
		Name:         util.RandomCollection(),
		Status:       util.RandomCollectionStatus(),
	}

	err := testQueries.UpdateCollection(context.Background(), arg)
	require.NoError(t, err)

	collection2, err := getCollectionById(arg.CollectionID)
	require.NoError(t, err)
	require.NotEmpty(t, collection2)

	require.Equal(t, collection1.CollectionID, collection2.CollectionID)
	require.Equal(t, arg.Name, collection2.Name)
	require.Equal(t, arg.Status, collection2.Status)
}

func TestDeleteCollection(t *testing.T) {
	collection1 := createRandomCollection(t)
	err := testQueries.DeleteCollection(context.Background(), collection1.CollectionID)
	require.NoError(t, err)

	collection2, err := getCollectionById(collection1.CollectionID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, collection2)
}

func TestListCollections(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomCollection(t)
	}

	arg := ListCollectionsParams{
		Limit:  2,
		Offset: 3,
	}

	collections, err := testQueries.ListCollections(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, collections, 2)
	for _, collection := range collections {
		require.NotEmpty(t, collection)
	}
}
