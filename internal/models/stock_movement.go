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
	ID              uuid.UUID         `db:"id"`
	StoreProductID  uuid.UUID         `db:"store_product_id"` // Merujuk ke store_products.id
	StoreID         uuid.UUID         `db:"store_id"`
	MovementType    StockMovementType `db:"movement_type"`
	QuantityChanged int32             `db:"quantity_changed"` // Bisa positif atau negatif
	MovementDate    time.Time         `db:"movement_date"`    // TIMESTAMPTZ
	ReferenceID     uuid.NullUUID     `db:"reference_id"`     // Bisa NULL
	ReferenceType   sql.NullString    `db:"reference_type"`   // Bisa NULL
	Notes           sql.NullString    `db:"notes"`
	CreatedByUserID uuid.NullUUID     `db:"created_by_user_id"`
	CreatedAt       time.Time         `db:"created_at"`
	// UpdatedAt        time.Time         `db:"updated_at"` // Tambahkan jika ada di DB
}
