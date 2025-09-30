package configs

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() (*pgxpool.Pool, error) {
	user := os.Getenv("PG_USER")
	pwd := os.Getenv("PG_PWD")
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	db := os.Getenv("PG_DB")

	connstring := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s", user, pwd, host, port, db,
	)

	return pgxpool.New(context.Background(), connstring)
}

func PingDB(p *pgxpool.Pool) error {
	return p.Ping(context.Background())
}
