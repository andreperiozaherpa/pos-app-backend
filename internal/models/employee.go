package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Employee merepresentasikan data karyawan dari tabel 'employees'.
type Employee struct {
	UserID           uuid.UUID      `db:"user_id"` // Primary Key, juga Foreign Key ke users.id
	CompanyID        uuid.UUID      `db:"company_id"`
	StoreID          uuid.NullUUID  `db:"store_id"`           // Bisa NULL untuk peran sentral
	EmployeeIDNumber sql.NullString `db:"employee_id_number"` // Bisa NULL
	JoinDate         sql.NullTime   `db:"join_date"`          // Menggunakan sql.NullTime karena DATE bisa NULL
	Position         sql.NullString `db:"position"`           // Bisa NULL
	CreatedAt        time.Time      `db:"created_at"`
	UpdatedAt        time.Time      `db:"updated_at"`
}
