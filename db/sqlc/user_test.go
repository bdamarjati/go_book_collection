package sqlc

import (
	"context"
	"database/sql"

	"testing"

	"github.com/bdamarjati/go_book_collection/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomUser(),
		Role:     util.RandomRole(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	rowsAffected, err := user.RowsAffected()
	require.NoError(t, err)
	require.NotZero(t, rowsAffected)

	return User{
		Username: arg.Username,
		Role:     arg.Role,
	}
}

func getUserByUsername(username string) (User, error) {
	user, err := testQueries.GetUser(context.Background(), username)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := getUserByUsername(user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Role, user2.Role)
	require.NotEmpty(t, user2.CreatedAt)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		Username: user1.Username,
		Role:     util.RandomRole(),
	}

	err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)

	user2, err := getUserByUsername(user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.Role, user2.Role)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.Username)
	require.NoError(t, err)

	user2, err := getUserByUsername(user1.Username)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
