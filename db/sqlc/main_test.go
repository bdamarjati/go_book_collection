package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/bdamarjati/go_book_collection/util"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.MySqlSource)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
