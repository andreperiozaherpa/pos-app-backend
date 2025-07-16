package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Supplier merepresentasikan data pemasok dari tabel 'suppliers'.
type Supplier struct {
	ID            uuid.UUID      `db:"id" json:"id"`
	CompanyID     uuid.UUID      `db:"company_id" json:"company_id"`
	Name          string         `db:"name" json:"name"`
	ContactPerson sql.NullString `db:"contact_person" json:"contact_person,omitempty"` // Bisa NULL
	Email         sql.NullString `db:"email" json:"email,omitempty"`                   // Bisa NULL
	PhoneNumber   sql.NullString `db:"phone_number" json:"phone_number,omitempty"`     // Bisa NULL
	Address       sql.NullString `db:"address" json:"address,omitempty"`               // Bisa NULL
	CreatedAt     time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at" json:"updated_at"`
}
