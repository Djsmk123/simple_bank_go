package main

import (
	"database/sql"
	"log"
	"simple_bank/db/api"
	db "simple_bank/db/db/sqlc"
	"simple_bank/db/utils"
)

func main() {
	config, err := utils.LoadConfiguration(".")
	if err != nil {
		log.Fatal("Failed to load configuration", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}
