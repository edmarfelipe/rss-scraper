package database

import (
	"database/sql"
	"fmt"

	"github.com/edmarfelipe/rss-scraper/internal/env"
	sqlm "github.com/edmarfelipe/rss-scraper/sql"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func NewConnection() (*Queries, error) {
	conn, err := sql.Open("postgres", env.Config.DBConn())
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	if err := Migrate(conn); err != nil {
		return nil, err
	}

	return New(conn), nil
}

func Migrate(conn *sql.DB) error {
	goose.SetBaseFS(sqlm.Migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("error setting dialect: %w", err)
	}

	if err := goose.Up(conn, "schema"); err != nil {
		return fmt.Errorf("error running migrations: %w", err)
	}

	return nil
}
