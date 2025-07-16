package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// StockTransferStatus mendefinisikan status transfer stok internal.
type StockTransferStatus string

// Konstanta untuk StockTransferStatus
const (
	StockTransferStatusPending           StockTransferStatus = "PENDING"
	StockTransferStatusApproved          StockTransferStatus = "APPROVED"
	StockTransferStatusShipped           StockTransferStatus = "SHIPPED"
	StockTransferStatusPartiallyReceived StockTransferStatus = "PARTIALLY_RECEIVED"
	StockTransferStatusReceived          StockTransferStatus = "RECEIVED"
	StockTransferStatusCancelled         StockTransferStatus = "CANCELLED"
)

// StockTransfer merepresentasikan data transfer stok dari tabel 'stock_transfers'.
type StockTransfer struct {
	ID            string    `db:"id" json:"id"`
	FromStoreID   string    `db:"from_store_id" json:"from_store_id"`
	ToStoreID     string    `db:"to_store_id" json:"to_store_id"`
	TransferredAt time.Time `db:"transferred_at" json:"transferred_at"`
	Status        string    `db:"status" json:"status"` // pending, completed, cancelled
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

// InternalStockTransfer merepresentasikan data transfer stok internal dari tabel 'internal_stock_transfers'.
type InternalStockTransfer struct {
	ID                 uuid.UUID                   `db:"id" json:"id"`
	TransferCode       string                      `db:"transfer_code" json:"transfer_code"`
	CompanyID          uuid.UUID                   `db:"company_id" json:"company_id"`
	SourceStoreID      uuid.UUID                   `db:"source_store_id" json:"source_store_id"`
	DestinationStoreID uuid.UUID                   `db:"destination_store_id" json:"destination_store_id"`
	TransferDate       time.Time                   `db:"transfer_date" json:"transfer_date"` // DATE
	Status             StockTransferStatus         `db:"status" json:"status"`               // pending, completed, cancelled
	Notes              sql.NullString              `db:"notes" json:"notes,omitempty"`
	RequestedByUserID  uuid.NullUUID               `db:"requested_by_user_id" json:"requested_by_user_id,omitempty"`
	ApprovedByUserID   uuid.NullUUID               `db:"approved_by_user_id" json:"approved_by_user_id,omitempty"`
	ShippedByUserID    uuid.NullUUID               `db:"shipped_by_user_id" json:"shipped_by_user_id,omitempty"`
	ReceivedByUserID   uuid.NullUUID               `db:"received_by_user_id" json:"received_by_user_id,omitempty"`
	CreatedAt          time.Time                   `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time                   `db:"updated_at" json:"updated_at"`
	Items              []InternalStockTransferItem `db:"-" json:"items,omitempty"` // Untuk menampung item saat mengambil data
}

// InternalStockTransferItem merepresentasikan data item dalam transfer stok internal.
type InternalStockTransferItem struct {
	ID                      uuid.UUID      `db:"id" json:"id"`
	InternalStockTransferID uuid.UUID      `db:"internal_stock_transfer_id" json:"internal_stock_transfer_id"`
	SourceStoreProductID    uuid.UUID      `db:"source_store_product_id" json:"source_store_product_id"` // ID produk di toko sumber
	QuantityRequested       int32          `db:"quantity_requested" json:"quantity_requested"`
	QuantityShipped         int32          `db:"quantity_shipped" json:"quantity_shipped"`
	QuantityReceived        int32          `db:"quantity_received" json:"quantity_received"`
	Notes                   sql.NullString `db:"notes" json:"notes,omitempty"`
	CreatedAt               time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt               time.Time      `db:"updated_at" json:"updated_at"`
}
