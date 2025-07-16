package models

import (
	"database/sql"
)

// Role merepresentasikan data peran dari tabel 'roles'.
type Role struct {
	ID          int32          `db:"id" json:"id"`                             // SERIAL di PostgreSQL (PK)
	Name        string         `db:"name" json:"name"`                         // Nama peran, unik dan wajib
	Description sql.NullString `db:"description" json:"description,omitempty"` // Bisa NULL
}
