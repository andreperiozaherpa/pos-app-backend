package models

import (
	"time"

	"github.com/google/uuid"
)

type StockMovementSummary struct {
	ID             uuid.UUID `db:"id" json:"id"`
	StoreProductID uuid.UUID `db:"store_product_id" json:"store_product_id"`
	PeriodStart    time.Time `db:"period_start" json:"period_start"`
	PeriodEnd      time.Time `db:"period_end" json:"period_end"`
	TotalIn        int       `db:"total_in" json:"total_in"`
	TotalOut       int       `db:"total_out" json:"total_out"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}
