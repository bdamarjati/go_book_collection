package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/bdamarjati/go_book_collection/util"
	"github.com/go-sql-driver/mysql"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	cfg := mysql.Config{
		User:                 config.DBUser,
		Passwd:               config.DBPassword,
		Net:                  config.DBNet,
		Addr:                 config.DBAddr,
		DBName:               config.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	conn, err := sql.Open(config.DBDriver, cfg.FormatDSN())
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
