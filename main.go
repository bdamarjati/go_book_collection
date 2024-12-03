package main

import (
	"database/sql"
	"log"

	"github.com/bdamarjati/go_book_collection/api"
	sqlc "github.com/bdamarjati/go_book_collection/db/sqlc"
	"github.com/bdamarjati/go_book_collection/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.MySqlSource)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}

	store := sqlc.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
