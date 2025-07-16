package models

import (
	"time"

	"github.com/google/uuid"
)

type StockReport struct {
	ID              uuid.UUID `db:"id" json:"id"`
	StoreID         uuid.UUID `db:"store_id" json:"store_id"`
	ReportDate      time.Time `db:"report_date" json:"report_date"`
	TotalStockValue float64   `db:"total_stock_value" json:"total_stock_value"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}
