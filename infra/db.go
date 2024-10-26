package infra

import (
	"context"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

func (i *Infrastructure) ConnectDB() (*pgx.Conn, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=db port=5432 user=postgres password=postgres dbname=cassette sslmode=disable")
	if err != nil {
		return nil, err
	}

	return conn, err
}
func (i *Infrastructure) CloseDB() error {
	sqlDB, err := i.ConnectDB()
	if err != nil {
		return err
	}
	sqlDB.Close(context.Background())
	return nil
}
