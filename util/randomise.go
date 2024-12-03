package util

import (
	"database/sql"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomUser() sql.NullString {
	return sql.NullString{
		String: randomString(6),
		Valid:  true,
	}
}

func RandomRole() sql.NullString {
	roles := []string{"Admin", "User", "Superadmin"}
	n := len(roles)
	return sql.NullString{
		String: roles[rand.Intn(n)],
		Valid:  true,
	}
}
