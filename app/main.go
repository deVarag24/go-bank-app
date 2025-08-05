package main

import (
	"database/sql"
	"log"
	"my-go-app/app/api"
	db "my-go-app/app/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:postgres@localhost:5432/go_bank_app_db?sslmode=disable"
    serverAddress = "localhost:3000"
)

func main() {
    conn, err := sql.Open(dbDriver, dbSource)
    if err != nil {
        log.Fatal("Failed to open DB:", err)
    }
    defer conn.Close()

    if err := conn.Ping(); err != nil {
        log.Fatal("Failed to connect to db:", err)
    }

    store := db.NewStore(conn)
    server := api.NewServer(store)

    err = server.StartServer(serverAddress)
    if err != nil {
        log.Fatal("Failed to connect to server:", err)
    }

}
