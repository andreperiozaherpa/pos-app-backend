package models

import (
	"time"

	"github.com/google/uuid"
)

// EmployeeLeave merepresentasikan data cuti/izin karyawan pada aplikasi POS.
type EmployeeLeave struct {
	ID             uuid.UUID  `db:"id" json:"id"`
	EmployeeUserID uuid.UUID  `db:"employee_user_id" json:"employee_user_id"` // FK ke Employee
	LeaveType      string     `db:"leave_type" json:"leave_type"`             // Jenis cuti (tahunan, sakit, izin, dsb)
	StartDate      time.Time  `db:"start_date" json:"start_date"`
	EndDate        time.Time  `db:"end_date" json:"end_date"`
	TotalDays      int        `db:"total_days" json:"total_days"` // Durasi cuti
	Reason         string     `db:"reason" json:"reason"`
	Status         string     `db:"status" json:"status"`           // Pengajuan, Disetujui, Ditolak, Selesai
	ApprovedBy     *uuid.UUID `db:"approved_by" json:"approved_by"` // User ID yang approve, nullable
	Notes          string     `db:"notes" json:"notes"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at" json:"updated_at"`
}
