package models

import (
	"time"

	"github.com/google/uuid"
)

type SalesReport struct {
	ID         uuid.UUID `db:"id" json:"id"`
	StoreID    uuid.UUID `db:"store_id" json:"store_id"`
	ReportDate time.Time `db:"report_date" json:"report_date"` // gunakan tipe time.Time agar mudah proses range & filter
	TotalSales float64   `db:"total_sales" json:"total_sales"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
