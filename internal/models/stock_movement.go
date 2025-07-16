package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// StockMovementType mendefinisikan tipe pergerakan stok.
type StockMovementType string

// Konstanta untuk StockMovementType
const (
	MovementTypeSale             StockMovementType = "SALE"
	MovementTypePurchaseReceipt  StockMovementType = "PURCHASE_RECEIPT"
	MovementTypeReturnToSupplier StockMovementType = "RETURN_TO_SUPPLIER"
	MovementTypeCustomerReturn   StockMovementType = "CUSTOMER_RETURN"
	MovementTypeAdjustmentIn     StockMovementType = "ADJUSTMENT_IN"
	MovementTypeAdjustmentOut    StockMovementType = "ADJUSTMENT_OUT"
	MovementTypeTransferOut      StockMovementType = "TRANSFER_OUT"
	MovementTypeTransferIn       StockMovementType = "TRANSFER_IN"
)

// StockMovement merepresentasikan data pergerakan stok dari tabel 'stock_movements'.
type StockMovement struct {
	ID              uuid.UUID         `db:"id" json:"id"`
	StoreProductID  uuid.UUID         `db:"store_product_id" json:"store_product_id"`
	StoreID         uuid.UUID         `db:"store_id" json:"store_id"`
	MovementType    StockMovementType `db:"movement_type" json:"movement_type"`
	QuantityChanged int32             `db:"quantity_changed" json:"quantity_changed"`
	MovementDate    time.Time         `db:"movement_date" json:"movement_date"`
	ReferenceID     uuid.NullUUID     `db:"reference_id" json:"reference_id,omitempty"`
	ReferenceType   sql.NullString    `db:"reference_type" json:"reference_type,omitempty"`
	Notes           sql.NullString    `db:"notes" json:"notes,omitempty"`
	CreatedByUserID uuid.NullUUID     `db:"created_by_user_id" json:"created_by_user_id,omitempty"`
	CreatedAt       time.Time         `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time         `db:"updated_at" json:"updated_at"`
}
