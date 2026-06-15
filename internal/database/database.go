package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type Database struct {
	conn *sql.DB
}

func New() (*Database, error) {
	_ = godotenv.Load()

	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DATABASE")
	sslmode := os.Getenv("POSTGRES_SSLMODE")

	if username == "" || password == "" || databaseName == "" || sslmode == "" {
		return nil, fmt.Errorf("missing postgres configuration")
	}

	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, databaseName, sslmode)

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := conn.PingContext(ctx); err != nil {
		_ = conn.Close()
		return nil, err
	}

	db := &Database{conn: conn}
	if err := db.createSchema(ctx); err != nil {
		_ = conn.Close()
		return nil, err
	}

	return db, nil
}

func (db *Database) Close() error {
	return db.conn.Close()
}

func (db *Database) createSchema(ctx context.Context) error {
	const schema = `
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    shift INTEGER NOT NULL,
    original_text TEXT NOT NULL,
    shifted_text TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
`

	_, err := db.conn.ExecContext(ctx, schema)
	return err
}

func (db *Database) SaveMessage(ctx context.Context, shift int, original, shifted string) error {
	_, err := db.conn.ExecContext(ctx,
		"INSERT INTO messages (shift, original_text, shifted_text) VALUES ($1, $2, $3)",
		shift,
		original,
		shifted,
	)
	return err
}
