package models

import (
	"time"

	"github.com/google/uuid"
)

type CustomerActivityReport struct {
	ID                uuid.UUID  `db:"id" json:"id"`
	CustomerUserID    uuid.UUID  `db:"customer_user_id" json:"customer_user_id"`
	CompanyID         uuid.UUID  `db:"company_id" json:"company_id"`
	ActivityDate      time.Time  `db:"activity_date" json:"activity_date"`
	TotalTransactions int        `db:"total_transactions" json:"total_transactions"`
	TotalAmount       float64    `db:"total_amount" json:"total_amount"`
	PointsEarned      int        `db:"points_earned" json:"points_earned"`
	LastTransactionAt *time.Time `db:"last_transaction_at" json:"last_transaction_at,omitempty"`
	CreatedAt         time.Time  `db:"created_at" json:"created_at"`
}
