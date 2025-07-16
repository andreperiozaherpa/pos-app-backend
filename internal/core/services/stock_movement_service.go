package services

import (
	"context"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

// StockMovementService menangani manajemen mutasi stok produk pada setiap store.
type StockMovementService interface {
	// CRUD utama
	CreateStockMovement(ctx context.Context, movement *models.StockMovement) (uuid.UUID, error)
	ListStockMovementsByStoreProductID(ctx context.Context, storeProductID uuid.UUID) ([]*models.StockMovement, error)

	// Opsional & Custom
	GetStockMovementByID(ctx context.Context, id uuid.UUID) (*models.StockMovement, error)
	ListStockMovementsByDateRange(ctx context.Context, storeID uuid.UUID, from, to time.Time) ([]*models.StockMovement, error)
	ExportStockMovements(ctx context.Context, storeID uuid.UUID, from, to time.Time) ([]byte, error)
	GetStockMovementSummary(ctx context.Context, storeID uuid.UUID, from, to time.Time) (*models.StockMovementSummary, error)
}
