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
	ID                   uuid.UUID           `db:"id"`
	StoreID              uuid.UUID           `db:"store_id"`
	SupplierID           uuid.UUID           `db:"supplier_id"`
	OrderDate            time.Time           `db:"order_date"` // DATE
	ExpectedDeliveryDate sql.NullTime        `db:"expected_delivery_date"`
	Status               PurchaseOrderStatus `db:"status"`
	TotalAmount          sql.NullFloat64     `db:"total_amount"` // DECIMAL bisa NULL
	Notes                sql.NullString      `db:"notes"`
	CreatedByUserID      uuid.UUID           `db:"created_by_user_id"`
	CreatedAt            time.Time           `db:"created_at"`
	UpdatedAt            time.Time           `db:"updated_at"`
	Items                []PurchaseOrderItem `db:"-"` // Untuk menampung item saat mengambil data
}

// PurchaseOrderItem merepresentasikan data item dalam pesanan pembelian dari tabel 'purchase_order_items'.
type PurchaseOrderItem struct {
	ID                   uuid.UUID `db:"id"`
	PurchaseOrderID      uuid.UUID `db:"purchase_order_id"`
	MasterProductID      uuid.UUID `db:"master_product_id"` // Merujuk ke master_products.id
	QuantityOrdered      int32     `db:"quantity_ordered"`
	PurchasePricePerUnit float64   `db:"purchase_price_per_unit"`
	QuantityReceived     int32     `db:"quantity_received"`
	Subtotal             float64   `db:"subtotal"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}
