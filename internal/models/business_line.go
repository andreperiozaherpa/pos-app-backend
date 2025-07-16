package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// BusinessLine merepresentasikan data lini usaha dari tabel 'business_lines'.
type BusinessLine struct {
	ID          uuid.UUID      `db:"id" json:"id"`
	CompanyID   uuid.UUID      `db:"company_id" json:"company_id"`
	Name        string         `db:"name" json:"name"`
	Description sql.NullString `db:"description" json:"description,omitempty"` // Kolom description bisa NULL
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
}
