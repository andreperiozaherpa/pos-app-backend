package models

import (
	"database/sql"
)

// Permission merepresentasikan data izin (hak akses) dari tabel 'permissions'.
type Permission struct {
	ID          int32          `db:"id"` // SERIAL di PostgreSQL
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"` // Bisa NULL
	GroupName   sql.NullString `db:"group_name"`  // Bisa NULL
}
