package models

import (
	"time"

	"github.com/google/uuid"
)

// ExpenseReport merepresentasikan hasil laporan pengeluaran (expense) untuk keperluan rekap, dashboard, atau export.
type ExpenseReport struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	CompanyID   uuid.UUID  `db:"company_id" json:"company_id"`
	StoreID     *uuid.UUID `db:"store_id" json:"store_id"`         // Nullable jika laporan per perusahaan
	PeriodStart time.Time  `db:"period_start" json:"period_start"` // Awal periode laporan
	PeriodEnd   time.Time  `db:"period_end" json:"period_end"`     // Akhir periode laporan
	Category    string     `db:"category" json:"category"`         // Kategori expense (gaji, listrik, dll)
	TotalAmount float64    `db:"total_amount" json:"total_amount"` // Total pengeluaran pada periode tsb
	Detail      string     `db:"detail" json:"detail"`             // Optional: JSON/CSV detail rekap (jika diperlukan)
	GeneratedAt time.Time  `db:"generated_at" json:"generated_at"`
	GeneratedBy uuid.UUID  `db:"generated_by" json:"generated_by"` // User yang generate laporan
}
