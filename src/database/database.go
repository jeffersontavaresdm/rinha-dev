package database

import (
	"context"
	"database/sql"
	"os"
	"time"
)

func OpenDB() *sql.DB {

	// Format: postgres://user:password@host:port/dbname?connect_timeout=5&sslmode=disable
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		panic("DATABASE_URL not defined")
	}

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	return db
}
