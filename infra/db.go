package infra

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

func (i *Infrastructure) ConnectDB() *pgx.Conn {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=db port=5432 user=postgres password=postgres dbname=cassette sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	return conn
}
