package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
    _ "github.com/lib/pq"

)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:postgres@localhost:5432/go_bank_app_db?sslmode=disable"
)

var testQueries *Queries

var testDB *sql.DB

func TestMain(m *testing.M){
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil{
		log.Fatal("DB Connection Fail: ", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}