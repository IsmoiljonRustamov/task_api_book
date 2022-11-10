package main

import (
	"books/apiy"
	"books/storage"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	PostgresUser = "postgres"
	PostgresPassword = "12345"
	PostgresHost = "localhost"
	PostgresPort = 5432
	PostgresDatabase = "book"
)

func main() {
	connStr := fmt.Sprintf("user = %s password = %s host = %s port = %d dbname = %s sslmode = disable",
	PostgresUser, PostgresPassword, PostgresHost, PostgresPort, PostgresDatabase)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection")
	}
	storage := storage.NewDBManager(db)

	server := apiy.NewServer(storage)

	err = server.Run(":8001")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}

