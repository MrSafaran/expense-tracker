package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnection(databaseURL string) (*pgxpool.Pool, error) {

	dbpool, err := pgxpool.New(
		context.Background(),
		databaseURL,
	)

	if err != nil {
		return nil, err
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL")

	return dbpool, nil
}
