package models

import (
	"time"

	"github.com/google/uuid"
)

// EmployeeAttendance merepresentasikan data absensi harian karyawan pada aplikasi POS.
type EmployeeAttendance struct {
	ID               uuid.UUID  `db:"id" json:"id"`
	EmployeeUserID   uuid.UUID  `db:"employee_user_id" json:"employee_user_id"` // FK ke Employee
	StoreID          uuid.UUID  `db:"store_id" json:"store_id"`                 // FK ke Store
	ShiftID          uuid.UUID  `db:"shift_id" json:"shift_id"`                 // FK ke Shift (opsional)
	AttendanceDate   time.Time  `db:"attendance_date" json:"attendance_date"`
	CheckInTime      *time.Time `db:"check_in_time" json:"check_in_time"`         // Nullable, absen masuk
	CheckOutTime     *time.Time `db:"check_out_time" json:"check_out_time"`       // Nullable, absen keluar
	AttendanceStatus string     `db:"attendance_status" json:"attendance_status"` // hadir, izin, sakit, dll
	Notes            string     `db:"notes" json:"notes"`
	CreatedAt        time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at" json:"updated_at"`
}
