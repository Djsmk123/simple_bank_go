package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/djsmk123/simplebank/utils"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := utils.LoadConfiguration("../..")
	if err != nil {
		log.Fatal("Failed to load configuration", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Something went wrong opening database", err)
	}
	testQueries = New(testDB)
	m.Run()
	os.Exit(m.Run())

}
