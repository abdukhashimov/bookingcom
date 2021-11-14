package handler_test

import (
	"abdukhashimov/mybron.uz/storage/sqlc"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	postgresConn *sql.DB
	err          error
	queries      *sqlc.Queries
)

func TestMain(m *testing.M) {
	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		"localhost",
		5432,
		"postgres",
		"postgres",
		"booking",
		"disable",
	)

	postgresConn, err = sql.Open("postgres", conStr)

	if err != nil {
		log.Fatal(err)
	}

	queries = sqlc.New(postgresConn)
	os.Exit(m.Run())
}
