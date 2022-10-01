package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/kntajus/enumarray/db"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), `INSERT INTO choices (id, fruits) VALUES (1, '{"apple", "kiwi"}') ON CONFLICT DO NOTHING;`)
	if err != nil {
		os.Exit(2)
	}

	store := db.New(conn)
	_, _ = store.GetChoice(context.Background(), 1)
}
