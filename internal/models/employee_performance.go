package models

import (
	"time"

	"github.com/google/uuid"
)

// EmployeePerformanceSummary merepresentasikan ringkasan performa/kinerja karyawan dalam periode tertentu.
type EmployeePerformanceSummary struct {
	ID             uuid.UUID `db:"id" json:"id"`
	EmployeeUserID uuid.UUID `db:"employee_user_id" json:"employee_user_id"`
	PeriodStart    time.Time `db:"period_start" json:"period_start"`     // Awal periode penilaian (misal: awal bulan)
	PeriodEnd      time.Time `db:"period_end" json:"period_end"`         // Akhir periode penilaian
	TotalShifts    int       `db:"total_shifts" json:"total_shifts"`     // Jumlah shift
	TotalPresence  int       `db:"total_presence" json:"total_presence"` // Total hadir
	TotalLate      int       `db:"total_late" json:"total_late"`         // Jumlah keterlambatan
	TotalAbsence   int       `db:"total_absence" json:"total_absence"`   // Total absen/tidak hadir
	SalesAchieved  float64   `db:"sales_achieved" json:"sales_achieved"` // Total penjualan dicapai
	Note           string    `db:"note" json:"note"`                     // Catatan/catatan manajer
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}
