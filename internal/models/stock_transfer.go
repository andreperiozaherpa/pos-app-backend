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

// InternalStockTransfer merepresentasikan data transfer stok internal dari tabel 'internal_stock_transfers'.
type InternalStockTransfer struct {
	ID                 uuid.UUID                   `db:"id"`
	TransferCode       string                      `db:"transfer_code"`
	CompanyID          uuid.UUID                   `db:"company_id"`
	SourceStoreID      uuid.UUID                   `db:"source_store_id"`
	DestinationStoreID uuid.UUID                   `db:"destination_store_id"`
	TransferDate       time.Time                   `db:"transfer_date"` // DATE
	Status             StockTransferStatus         `db:"status"`
	Notes              sql.NullString              `db:"notes"`
	RequestedByUserID  uuid.NullUUID               `db:"requested_by_user_id"`
	ApprovedByUserID   uuid.NullUUID               `db:"approved_by_user_id"`
	ShippedByUserID    uuid.NullUUID               `db:"shipped_by_user_id"`
	ReceivedByUserID   uuid.NullUUID               `db:"received_by_user_id"`
	CreatedAt          time.Time                   `db:"created_at"`
	UpdatedAt          time.Time                   `db:"updated_at"`
	Items              []InternalStockTransferItem `db:"-"` // Untuk menampung item saat mengambil data
}

// InternalStockTransferItem merepresentasikan data item dalam transfer stok internal.
type InternalStockTransferItem struct {
	ID                      uuid.UUID      `db:"id"`
	InternalStockTransferID uuid.UUID      `db:"internal_stock_transfer_id"`
	SourceStoreProductID    uuid.UUID      `db:"source_store_product_id"` // ID produk di toko sumber
	QuantityRequested       int32          `db:"quantity_requested"`
	QuantityShipped         int32          `db:"quantity_shipped"`
	QuantityReceived        int32          `db:"quantity_received"`
	Notes                   sql.NullString `db:"notes"`
	CreatedAt               time.Time      `db:"created_at"`
	UpdatedAt               time.Time      `db:"updated_at"`
}
