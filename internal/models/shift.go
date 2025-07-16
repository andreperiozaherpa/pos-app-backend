package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Shift merepresentasikan data shift karyawan dari tabel 'shifts'.
type Shift struct {
	ID              uuid.UUID      `db:"id" json:"id"`
	EmployeeUserID  uuid.UUID      `db:"employee_user_id" json:"employee_user_id"`
	StoreID         uuid.UUID      `db:"store_id" json:"store_id"`
	ShiftDate       time.Time      `db:"shift_date" json:"shift_date"`                           // DATE bisa dimapping ke time.Time (hanya bagian tanggal yang relevan)
	StartTime       time.Time      `db:"start_time" json:"start_time"`                           // TIME lebih baik dimapping ke time.Time
	EndTime         time.Time      `db:"end_time" json:"end_time"`                               // TIME lebih baik dimapping ke time.Time
	ActualCheckIn   sql.NullTime   `db:"actual_check_in" json:"actual_check_in,omitempty"`       // TIMESTAMPTZ bisa NULL
	ActualCheckOut  sql.NullTime   `db:"actual_check_out" json:"actual_check_out,omitempty"`     // TIMESTAMPTZ bisa NULL
	Notes           sql.NullString `db:"notes" json:"notes,omitempty"`                           // TEXT bisa NULL
	CreatedByUserID uuid.NullUUID  `db:"created_by_user_id" json:"created_by_user_id,omitempty"` // UUID bisa NULL
	CreatedAt       time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at" json:"updated_at"`
}
