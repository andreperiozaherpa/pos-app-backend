package postgres

import (
	"database/sql"
	"errors"
	"strings"
)

// ErrNotFound adalah error standar ketika data tidak ditemukan.
var ErrNotFound = sql.ErrNoRows

// IsDuplicateKeyError memeriksa apakah error merupakan duplicate key violation di Postgres.
func IsDuplicateKeyError(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	// Pesan error Postgres duplicate key biasanya mengandung "duplicate key"
	return strings.Contains(msg, "duplicate key")
}

// WrapError memberikan pesan error yang lebih informatif.
func WrapError(operation string, err error) error {
	if err == nil {
		return nil
	}
	return errors.New(operation + ": " + err.Error())
}
