package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Supplier merepresentasikan data pemasok dari tabel 'suppliers'.
type Supplier struct {
	ID            uuid.UUID      `db:"id"`
	CompanyID     uuid.UUID      `db:"company_id"`
	Name          string         `db:"name"`
	ContactPerson sql.NullString `db:"contact_person"` // Bisa NULL
	Email         sql.NullString `db:"email"`          // Bisa NULL
	PhoneNumber   sql.NullString `db:"phone_number"`   // Bisa NULL
	Address       sql.NullString `db:"address"`        // Bisa NULL
	CreatedAt     time.Time      `db:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at"`
}
