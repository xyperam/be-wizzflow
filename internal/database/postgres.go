package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() (*pgxpool.Pool, error) {
	dsn := "postgres://user:password@localhost:5433/wizzflow?sslmode=disable"
	fmt.Printf("Mencoba konek dengan DSN: '%s'\n", dsn)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	// check connection
	err = pool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSql")
	return pool, nil
}
