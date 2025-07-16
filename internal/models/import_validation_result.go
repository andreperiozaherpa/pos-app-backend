package models

import (
	"time"
)

type ImportValidationResult struct {
	IsValid   bool      `db:"is_valid" json:"is_valid"`
	Errors    []string  `db:"errors" json:"errors"`
	CheckedAt time.Time `db:"checked_at" json:"checked_at"`
}
