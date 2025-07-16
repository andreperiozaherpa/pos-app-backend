package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentInfo struct {
	ID            uuid.UUID `db:"id" json:"id"`
	TransactionID uuid.UUID `db:"transaction_id" json:"transaction_id"`
	PaymentMethod string    `db:"payment_method" json:"payment_method"`
	Amount        float64   `db:"amount" json:"amount"`
	PaymentDate   time.Time `db:"payment_date" json:"payment_date"`
}
