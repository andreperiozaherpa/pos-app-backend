package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// BusinessLine merepresentasikan data lini usaha dari tabel 'business_lines'.
type BusinessLine struct {
	ID          uuid.UUID      `db:"id"`
	CompanyID   uuid.UUID      `db:"company_id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"` // Kolom description bisa NULL
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
}
