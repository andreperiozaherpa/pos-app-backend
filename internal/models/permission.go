package models

import (
	"database/sql"
)

// Permission merepresentasikan data izin (hak akses) dari tabel 'permissions'.
type Permission struct {
	ID          int32          `db:"id" json:"id"`                             // SERIAL di PostgreSQL
	Name        string         `db:"name" json:"name"`                         // Nama permission
	Description sql.NullString `db:"description" json:"description,omitempty"` // Bisa NULL
	GroupName   sql.NullString `db:"group_name" json:"group_name,omitempty"`   // Bisa NULL
}
