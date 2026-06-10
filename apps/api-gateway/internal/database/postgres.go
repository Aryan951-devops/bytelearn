// Package database sets up and manages the database connection.
package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB is the shared connection pool used to talk to the database.
var DB *pgxpool.Pool

// ConnectDB creates a new connection pool using the provided database URL.
func ConnectDB(databaseURL string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal("Failed to create database pool:", err)
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = dbpool

	log.Println("Database connected successfully")
}
