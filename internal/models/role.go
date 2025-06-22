package models

import (
	"database/sql"
)

// Role merepresentasikan data peran dari tabel 'roles'.
type Role struct {
	ID          int32          `db:"id"` // SERIAL di PostgreSQL biasanya dimapping ke int32 atau int64
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"` // Bisa NULL
}
