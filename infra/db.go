package infra

import (
	"context"

	sqlc "github.com/DeNA-Autumn-Hackathon2024-b/backend/db/sqlc_gen"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

func ConnectDB() (*pgx.Conn, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=db port=5432 user=postgres password=postgres dbname=cassette sslmode=disable")
	if err != nil {
		return nil, err
	}

	return conn, err
}
func (i *Infrastructure) CloseDB() error {
	sqlDB, err := ConnectDB()
	if err != nil {
		return err
	}
	sqlDB.Close(context.Background())
	return nil
}

func (i *Infrastructure) NewDB() *sqlc.Queries {
	sqlDB, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	return sqlc.New(sqlDB)
}
