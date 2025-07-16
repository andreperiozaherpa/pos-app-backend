package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// StoreProductService menangani logika bisnis produk pada level store (toko).
type StoreProductService interface {
	// CRUD Utama
	CreateStoreProduct(ctx context.Context, product *models.StoreProduct) (uuid.UUID, error)
	GetStoreProductByID(ctx context.Context, id uuid.UUID) (*models.StoreProduct, error)
	UpdateStoreProduct(ctx context.Context, product *models.StoreProduct) error
	DeleteStoreProduct(ctx context.Context, id uuid.UUID) error
	ListStoreProductsByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.StoreProduct, error)

	// Method Opsional/Custom
	UpdateStoreProductStock(ctx context.Context, id uuid.UUID, newStock int) error
	SearchStoreProducts(ctx context.Context, query string, storeID uuid.UUID) ([]*models.StoreProduct, error)
	BulkUpdateStoreProductStock(ctx context.Context, updates []*models.StoreProductStockUpdate) error
	ArchiveStoreProduct(ctx context.Context, id uuid.UUID) error
	RestoreStoreProduct(ctx context.Context, id uuid.UUID) error
	ListStoreProductMovements(ctx context.Context, id uuid.UUID) ([]*models.StockMovement, error)
	ExportStoreProducts(ctx context.Context, storeID uuid.UUID) ([]byte, error)
}
