package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// DiscountService menangani logika bisnis diskon dan promo.
type DiscountService interface {
	// CRUD utama
	CreateDiscount(ctx context.Context, discount *models.Discount) (uuid.UUID, error)
	GetDiscountByID(ctx context.Context, id uuid.UUID) (*models.Discount, error)
	UpdateDiscount(ctx context.Context, discount *models.Discount) error
	DeleteDiscount(ctx context.Context, id uuid.UUID) error
	ListDiscountsByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Discount, error)

	// Opsional/custom
	AssignDiscountToProduct(ctx context.Context, discountID, productID uuid.UUID) error
	BulkUpdateDiscounts(ctx context.Context, discounts []*models.Discount) error
	FindActiveDiscounts(ctx context.Context, companyID uuid.UUID) ([]*models.Discount, error)
	CheckDiscountEligibility(ctx context.Context, discountID, customerID uuid.UUID, amount float64) (bool, error)
	RemoveDiscountFromProduct(ctx context.Context, discountID, productID uuid.UUID) error
	ListProductsByDiscountID(ctx context.Context, discountID uuid.UUID) ([]*models.MasterProduct, error)
	ExportDiscounts(ctx context.Context, companyID uuid.UUID) ([]byte, error)
	ArchiveDiscount(ctx context.Context, id uuid.UUID) error
	RestoreDiscount(ctx context.Context, id uuid.UUID) error
}
