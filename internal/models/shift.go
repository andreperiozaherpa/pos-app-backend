package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Shift merepresentasikan data shift karyawan dari tabel 'shifts'.
type Shift struct {
	ID              uuid.UUID      `db:"id"`
	EmployeeUserID  uuid.UUID      `db:"employee_user_id"`
	StoreID         uuid.UUID      `db:"store_id"`
	ShiftDate       time.Time      `db:"shift_date"`         // DATE bisa dimapping ke time.Time (hanya bagian tanggal yang relevan)
	StartTime       time.Time      `db:"start_time"`         // TIME lebih baik dimapping ke time.Time
	EndTime         time.Time      `db:"end_time"`           // TIME lebih baik dimapping ke time.Time
	ActualCheckIn   sql.NullTime   `db:"actual_check_in"`    // TIMESTAMPTZ bisa NULL
	ActualCheckOut  sql.NullTime   `db:"actual_check_out"`   // TIMESTAMPTZ bisa NULL
	Notes           sql.NullString `db:"notes"`              // TEXT bisa NULL
	CreatedByUserID uuid.NullUUID  `db:"created_by_user_id"` // UUID bisa NULL
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
}
