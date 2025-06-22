package postgres

import (
	"context"
	"database/sql"
)

// DBExecutor defines the common methods for *sql.DB and *sql.Tx.
// This interface allows repositories to work with either a direct DB connection
// or a transaction, enabling the Unit of Work pattern.
type DBExecutor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}
