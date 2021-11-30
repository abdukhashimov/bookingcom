package services_test

import (
	"abdukhashimov/mybron.uz/config"
	"abdukhashimov/mybron.uz/pkg/jwt"
	"abdukhashimov/mybron.uz/pkg/logger"
	"abdukhashimov/mybron.uz/services"
	"abdukhashimov/mybron.uz/storage/sqlc"
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	postgresConn *sql.DB
	err          error
	queries      *sqlc.Queries
	cfg          *config.Config
	svs          *services.Services
	log          logger.Logger
)

func TestMain(m *testing.M) {
	cfg = config.NewConfig()
	log = logger.New("debug", "test-my-bron.uz")
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
		log.Fatal("error while connecting to postgres", logger.Error(err))
	}

	queries = sqlc.New(postgresConn)
	svs = services.NewServices(queries, jwt.NewJwt(cfg, log))
	os.Exit(m.Run())
}
