package database

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
)

const (
	maxRetries        = 60
	connMaxIdleTime   = 30 * time.Second
	maxIdleConns      = 30
	maxOpenConns      = 30
	connectionTimeout = 5 * time.Second
	retryDelay        = 1 * time.Second
)

// NewDatabaseConnection establishes a new PostgreSQL database connection with retry logic.
func NewDatabaseConnection(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Configure database connection pool settings
	db.SetConnMaxIdleTime(connMaxIdleTime)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)

	// Attempt to connect with retries
	if err := connectWithRetries(db); err != nil {
		return nil, err
	}

	slog.Info("Successfully connected to the database")
	return db, nil
}

// connectWithRetries tries to establish a connection to the database with retries.
func connectWithRetries(db *sql.DB) error {
	for i := 0; i < maxRetries; i++ {
		// Create a new context with a timeout for each retry
		ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
		defer cancel() // Cancel the context after each iteration

		// Try pinging the database
		if err := db.PingContext(ctx); err == nil {
			return nil
		}

		// Log the retry attempt number
		slog.Warn("Failed to connect to the database, retrying")
		time.Sleep(retryDelay)
	}

	return fmt.Errorf("couldn't connect to the database after %d retries", maxRetries)
}
