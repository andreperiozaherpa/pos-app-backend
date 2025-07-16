package models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionSummary struct {
	ID                uuid.UUID `db:"id" json:"id"`
	StoreID           uuid.UUID `db:"store_id" json:"store_id"`
	TotalTransactions int       `db:"total_transactions" json:"total_transactions"`
	TotalRevenue      float64   `db:"total_revenue" json:"total_revenue"`
	PeriodStart       time.Time `db:"period_start" json:"period_start"`
	PeriodEnd         time.Time `db:"period_end" json:"period_end"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
}
