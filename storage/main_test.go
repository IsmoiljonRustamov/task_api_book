package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	dbMangaer * DBManager

	PostgresUser = "postgres"
	PostgresPassword = "12345"
	PostgresHost = "localhost"
	PostgresPort = 5432
	PostgresDatabase = "book"
)

func TestMain(m *testing.M) {
	connStr := fmt.Sprintf("user = %s password = %s host = %s port = %d dbname= %s sslmode = disable",
	PostgresUser,PostgresPassword,PostgresHost,PostgresPort,PostgresDatabase)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	dbMangaer = NewDBManager(db)
	os.Exit(m.Run())
}