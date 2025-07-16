package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Customer merepresentasikan data pelanggan dari tabel 'customers'.
//
// PK: UserID (juga FK ke users.id)
type Customer struct {
	UserID           uuid.UUID      `db:"user_id" json:"user_id"`                               // Primary Key, juga Foreign Key ke users.id
	CompanyID        uuid.UUID      `db:"company_id" json:"company_id"`                         // FK ke companies.id
	MembershipNumber sql.NullString `db:"membership_number" json:"membership_number,omitempty"` // Bisa NULL (untuk non-member)
	JoinDate         sql.NullTime   `db:"join_date" json:"join_date,omitempty"`                 // Bisa NULL
	Points           int32          `db:"points" json:"points"`                                 // INTEGER default 0
	Tier             sql.NullString `db:"tier" json:"tier,omitempty"`                           // Bisa NULL (tier keanggotaan)
	CreatedAt        time.Time      `db:"created_at" json:"created_at"`                         // Timestamp buat
	UpdatedAt        time.Time      `db:"updated_at" json:"updated_at"`                         // Timestamp update

	// Relasi (digunakan untuk join/preload data, bukan kolom di tabel 'customers')
	User *User `json:"user,omitempty" db:"-"`
}
