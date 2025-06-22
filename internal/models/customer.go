package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Customer merepresentasikan data pelanggan dari tabel 'customers'.
type Customer struct {
	UserID           uuid.UUID      `db:"user_id"` // Primary Key, juga Foreign Key ke users.id
	CompanyID        uuid.UUID      `db:"company_id"`
	MembershipNumber sql.NullString `db:"membership_number"` // Bisa NULL (untuk non-member)
	JoinDate         sql.NullTime   `db:"join_date"`         // Bisa NULL
	Points           int32          `db:"points"`            // INTEGER default 0
	Tier             sql.NullString `db:"tier"`              // Bisa NULL
	CreatedAt        time.Time      `db:"created_at"`
	UpdatedAt        time.Time      `db:"updated_at"`

	// Relasi (digunakan untuk join/preload data, bukan kolom di tabel 'customers')
	User *User `json:"user,omitempty" db:"-"`
}
