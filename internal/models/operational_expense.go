package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// OperationalExpense merepresentasikan data pengeluaran operasional dari tabel 'operational_expenses'.
type OperationalExpense struct {
	ID              uuid.UUID      `db:"id"`
	CompanyID       uuid.UUID      `db:"company_id"`
	StoreID         uuid.NullUUID  `db:"store_id"` // Bisa NULL
	ExpenseDate     time.Time      `db:"expense_date"`
	Category        string         `db:"category"`
	Description     sql.NullString `db:"description"`
	Amount          float64        `db:"amount"`
	CreatedByUserID uuid.NullUUID  `db:"created_by_user_id"` // Bisa NULL
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
}
