package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// PurchaseOrderStatus mendefinisikan status pesanan pembelian.
type PurchaseOrderStatus string

// Konstanta untuk PurchaseOrderStatus
const (
	POStatusPending           PurchaseOrderStatus = "PENDING"
	POStatusOrdered           PurchaseOrderStatus = "ORDERED"
	POStatusPartiallyReceived PurchaseOrderStatus = "PARTIALLY_RECEIVED"
	POStatusReceived          PurchaseOrderStatus = "RECEIVED"
	POStatusCancelled         PurchaseOrderStatus = "CANCELLED"
)

// PurchaseOrder merepresentasikan data pesanan pembelian dari tabel 'purchase_orders'.
type PurchaseOrder struct {
	ID                   uuid.UUID           `db:"id" json:"id"`
	StoreID              uuid.UUID           `db:"store_id" json:"store_id"`
	SupplierID           uuid.UUID           `db:"supplier_id" json:"supplier_id"`
	OrderDate            time.Time           `db:"order_date" json:"order_date"`
	ExpectedDeliveryDate sql.NullTime        `db:"expected_delivery_date" json:"expected_delivery_date,omitempty"`
	Status               PurchaseOrderStatus `db:"status" json:"status"`
	TotalAmount          sql.NullFloat64     `db:"total_amount" json:"total_amount,omitempty"`
	Notes                sql.NullString      `db:"notes" json:"notes,omitempty"`
	CreatedByUserID      uuid.UUID           `db:"created_by_user_id" json:"created_by_user_id"`
	CreatedAt            time.Time           `db:"created_at" json:"created_at"`
	UpdatedAt            time.Time           `db:"updated_at" json:"updated_at"`
	Items                []PurchaseOrderItem `db:"-" json:"items,omitempty"` // Untuk menampung item saat mengambil data
}

// PurchaseOrderItem merepresentasikan data item dalam pesanan pembelian dari tabel 'purchase_order_items'.
type PurchaseOrderItem struct {
	ID                   uuid.UUID `db:"id" json:"id"`
	PurchaseOrderID      uuid.UUID `db:"purchase_order_id" json:"purchase_order_id"`
	MasterProductID      uuid.UUID `db:"master_product_id" json:"master_product_id"`
	QuantityOrdered      int32     `db:"quantity_ordered" json:"quantity_ordered"`
	PurchasePricePerUnit float64   `db:"purchase_price_per_unit" json:"purchase_price_per_unit"`
	QuantityReceived     int32     `db:"quantity_received" json:"quantity_received"`
	Subtotal             float64   `db:"subtotal" json:"subtotal"`
	CreatedAt            time.Time `db:"created_at" json:"created_at"`
	UpdatedAt            time.Time `db:"updated_at" json:"updated_at"`
}
