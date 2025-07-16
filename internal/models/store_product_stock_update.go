package models

import (
	"time"

	"github.com/google/uuid"
)

type StoreProductStockUpdate struct {
	ID               uuid.UUID `db:"id" json:"id"`
	StoreProductID   uuid.UUID `db:"store_product_id" json:"store_product_id"`
	AdjustmentType   string    `db:"adjustment_type" json:"adjustment_type"`
	QuantityBefore   int       `db:"quantity_before" json:"quantity_before"`
	QuantityAfter    int       `db:"quantity_after" json:"quantity_after"`
	AdjustedByUserID uuid.UUID `db:"adjusted_by_user_id" json:"adjusted_by_user_id"`
	Reason           string    `db:"reason" json:"reason"`
	AdjustmentDate   time.Time `db:"adjustment_date" json:"adjustment_date"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}
