package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// StoreProduct merepresentasikan data produk spesifik toko dari tabel 'store_products'.
type StoreProduct struct {
	ID                uuid.UUID       `db:"id"`
	MasterProductID   uuid.UUID       `db:"master_product_id"`
	StoreID           uuid.UUID       `db:"store_id"`
	SupplierID        uuid.NullUUID   `db:"supplier_id"`
	StoreSpecificSKU  sql.NullString  `db:"store_specific_sku"`
	PurchasePrice     float64         `db:"purchase_price"`
	SellingPrice      float64         `db:"selling_price"`
	WholesalePrice    sql.NullFloat64 `db:"wholesale_price"`
	Stock             int32           `db:"stock"`
	MinimumStockLevel sql.NullInt32   `db:"minimum_stock_level"`
	ExpiryDate        sql.NullTime    `db:"expiry_date"`
	CreatedAt         time.Time       `db:"created_at"`
	UpdatedAt         time.Time       `db:"updated_at"`
}
