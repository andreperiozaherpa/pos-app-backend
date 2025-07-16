package models

import (
	"time"

	"github.com/google/uuid"
)

type StockTransferHistory struct {
	ID              uuid.UUID `db:"id" json:"id"`
	StockTransferID uuid.UUID `db:"stock_transfer_id" json:"stock_transfer_id"`
	Action          string    `db:"action" json:"action"`
	ActionDate      time.Time `db:"action_date" json:"action_date"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}
