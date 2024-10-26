package infra

import (
	"context"
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type Connection interface {
	Connection() (*sql.DB, error)
	// Close(ctx context.Context) error
}

func setupDB(dbDriver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}
	return db, err
}

func Connect() *sql.DB {
	dbDriver := "postgres"
	dsn := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=cassette sslmode=disable"
	db, err := setupDB(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db
}

func ConnectDB() *pgx.Conn {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=cassette sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	return conn
}
