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

func RandomUser() string {
	return randomString(6)
}

func RandomRole() sql.NullString {
	roles := []string{"Admin", "User", "Superadmin"}
	n := len(roles)
	return sql.NullString{
		String: roles[rand.Intn(n)],
		Valid:  true,
	}
}

func RandomCollection() sql.NullString {
	return sql.NullString{
		String: randomString(8),
		Valid:  true,
	}
}

func RandomCollectionStatus() sql.NullInt32 {
	statuses := []int32{0, 1}
	n := len(statuses)
	return sql.NullInt32{
		Int32: statuses[rand.Intn(n)],
		Valid: true,
	}
}

func RandomText(n int) sql.NullString {
	return sql.NullString{
		String: randomString(n),
		Valid:  true,
	}
}

func RandomLanguage() sql.NullString {
	languages := []string{"EN", "FR", "DE", "ID", "JP", "CN"}
	n := len(languages)
	return sql.NullString{
		String: languages[rand.Intn(n)],
		Valid:  true,
	}
}

func RandomYear() sql.NullInt32 {
	year := rand.Intn(30) + 1990
	return sql.NullInt32{
		Int32: int32(year),
		Valid: true,
	}
}
