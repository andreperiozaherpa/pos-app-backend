package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// OperationalExpense merepresentasikan data pengeluaran operasional dari tabel 'operational_expenses'.
type OperationalExpense struct {
	ID              uuid.UUID      `db:"id" json:"id"`
	CompanyID       uuid.UUID      `db:"company_id" json:"company_id"`
	StoreID         uuid.NullUUID  `db:"store_id" json:"store_id,omitempty"` // Bisa NULL
	ExpenseDate     time.Time      `db:"expense_date" json:"expense_date"`
	Category        string         `db:"category" json:"category"`
	Description     sql.NullString `db:"description" json:"description,omitempty"`
	Amount          float64        `db:"amount" json:"amount"`
	CreatedByUserID uuid.NullUUID  `db:"created_by_user_id" json:"created_by_user_id,omitempty"` // Bisa NULL
	CreatedAt       time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at" json:"updated_at"`
}
