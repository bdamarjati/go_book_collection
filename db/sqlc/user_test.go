package sqlc

import (
	"context"
	"database/sql"

	// "log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username: sql.NullString{String: "John", Valid: true},
		Role:     sql.NullString{String: "Admin", Valid: true},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	rowsAffected, err := user.RowsAffected()
	require.NoError(t, err)
	require.NotZero(t, rowsAffected)
}
