package models

import (
	"time"
)

type CompanyFinancialSummary struct {
	CompanyID    string    `json:"company_id" db:"company_id"`
	PeriodStart  time.Time `json:"period_start" db:"period_start"`
	PeriodEnd    time.Time `json:"period_end" db:"period_end"`
	TotalRevenue float64   `json:"total_revenue" db:"total_revenue"`
	TotalExpense float64   `json:"total_expense" db:"total_expense"`
	NetProfit    float64   `json:"net_profit" db:"net_profit"`
	LastUpdated  time.Time `json:"last_updated" db:"last_updated"`
}
