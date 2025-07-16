package models

import (
	"time"

	"github.com/google/uuid"
)

type ProfitLossReport struct {
	ID           uuid.UUID `db:"id" json:"id"`
	CompanyID    uuid.UUID `db:"company_id" json:"company_id"`
	ReportDate   time.Time `db:"report_date" json:"report_date"`
	TotalRevenue float64   `db:"total_revenue" json:"total_revenue"`
	TotalExpense float64   `db:"total_expense" json:"total_expense"`
	NetProfit    float64   `db:"net_profit" json:"net_profit"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
