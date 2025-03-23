package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func ConnectToPostgresDB(cfg DBConfig) DB {

	url := fmt.Sprintf("host=%s port=%s database=%s user=%s password=%s pool_max_conns=10",
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Username,
		cfg.Password,
	)

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal("could not establish database connection, error: %w", err)
	}

	return DB{
		Pool: pool,
	}
}
