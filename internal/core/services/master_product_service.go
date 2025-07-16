package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// MasterProductService mendefinisikan kontrak use case produk pusat (master).
type MasterProductService interface {
	// CRUD Utama
	CreateMasterProduct(ctx context.Context, product *models.MasterProduct) (uuid.UUID, error)
	GetMasterProductByID(ctx context.Context, id uuid.UUID) (*models.MasterProduct, error)
	UpdateMasterProduct(ctx context.Context, product *models.MasterProduct) error
	DeleteMasterProduct(ctx context.Context, id uuid.UUID) error
	ListMasterProductsByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.MasterProduct, error)

	// Method Opsional/Custom
	GetMasterProductHistory(ctx context.Context, id uuid.UUID) ([]*models.MasterProductHistory, error)
	DeactivateMasterProduct(ctx context.Context, id uuid.UUID) error
	SearchMasterProducts(ctx context.Context, query string, companyID uuid.UUID) ([]*models.MasterProduct, error)
	BulkImportMasterProducts(ctx context.Context, products []*models.MasterProduct) error
	GetMasterProductStockLevels(ctx context.Context, id uuid.UUID) (int, error)
	ArchiveMasterProduct(ctx context.Context, id uuid.UUID) error
	RestoreMasterProduct(ctx context.Context, id uuid.UUID) error
	ListMasterProductVariants(ctx context.Context, id uuid.UUID) ([]*models.MasterProduct, error)
	ExportMasterProducts(ctx context.Context, companyID uuid.UUID) ([]byte, error)
	SyncMasterProductWithStoreProducts(ctx context.Context, id uuid.UUID) error
}
