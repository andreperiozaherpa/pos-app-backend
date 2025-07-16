package services

import (
	"context"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

// StockTransferService menangani proses transfer stok antar toko/store dalam satu perusahaan.
type StockTransferService interface {
	// CRUD utama
	CreateStockTransfer(ctx context.Context, transfer *models.StockTransfer) (uuid.UUID, error)
	GetStockTransferByID(ctx context.Context, id uuid.UUID) (*models.StockTransfer, error)
	UpdateStockTransfer(ctx context.Context, transfer *models.StockTransfer) error
	DeleteStockTransfer(ctx context.Context, id uuid.UUID) error
	ListStockTransfersByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.StockTransfer, error)

	// Opsional & Custom
	ApproveStockTransfer(ctx context.Context, transferID, approverID uuid.UUID) error
	CancelStockTransfer(ctx context.Context, transferID uuid.UUID, reason string) error
	ExportStockTransfers(ctx context.Context, companyID uuid.UUID, from, to time.Time) ([]byte, error)
	ListStockTransfersByDateRange(ctx context.Context, companyID uuid.UUID, from, to time.Time) ([]*models.StockTransfer, error)
	GetStockTransferHistory(ctx context.Context, transferID uuid.UUID) ([]*models.StockTransferHistory, error)
}
