package postgresql

import (
	"context"
	"app/config"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	codeTestRepo *codeRepo
	brandTestRepo *brandRepo
	categoryTestRepo *categoryRepo
	customerTestRepo *customerRepo
	orderTestRepo *orderRepo
	productTestRepo *productRepo
	staffTestRepo *staffRepo
	stockTestRepo *stockRepo
	storeTestRepo *storeRepo
)

func TestMain(m *testing.M) {
	cfg := config.Load()

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		panic(err)
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(pool)
	}

	codeTestRepo = NewCodeRepo(pool)

	os.Exit(m.Run())
}
