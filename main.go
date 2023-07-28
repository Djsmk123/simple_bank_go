package main

import (
	"database/sql"
	"log"
	"simple_bank/db/api"
	db "simple_bank/db/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://postgres:1235789@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}
