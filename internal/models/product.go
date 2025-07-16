package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// StoreProduct merepresentasikan data produk spesifik toko dari tabel 'store_products'.
type StoreProduct struct {
	ID                uuid.UUID       `db:"id" json:"id"`
	MasterProductID   uuid.UUID       `db:"master_product_id" json:"master_product_id"`
	StoreID           uuid.UUID       `db:"store_id" json:"store_id"`
	SupplierID        uuid.NullUUID   `db:"supplier_id" json:"supplier_id,omitempty"`
	StoreSpecificSKU  sql.NullString  `db:"store_specific_sku" json:"store_specific_sku,omitempty"`
	PurchasePrice     float64         `db:"purchase_price" json:"purchase_price"`
	SellingPrice      float64         `db:"selling_price" json:"selling_price"`
	WholesalePrice    sql.NullFloat64 `db:"wholesale_price" json:"wholesale_price,omitempty"`
	Stock             int32           `db:"stock" json:"stock"`
	MinimumStockLevel sql.NullInt32   `db:"minimum_stock_level" json:"minimum_stock_level,omitempty"`
	ExpiryDate        sql.NullTime    `db:"expiry_date" json:"expiry_date,omitempty"`
	CreatedAt         time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time       `db:"updated_at" json:"updated_at"`
}
