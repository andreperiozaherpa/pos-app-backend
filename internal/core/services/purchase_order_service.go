package services

import (
	"context"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

// PurchaseOrderService menangani logika bisnis untuk Purchase Order (PO) barang ke supplier.
type PurchaseOrderService interface {
	// CRUD utama
	CreatePurchaseOrder(ctx context.Context, po *models.PurchaseOrder) (uuid.UUID, error)
	GetPurchaseOrderByID(ctx context.Context, id uuid.UUID) (*models.PurchaseOrder, error)
	UpdatePurchaseOrder(ctx context.Context, po *models.PurchaseOrder) error
	DeletePurchaseOrder(ctx context.Context, id uuid.UUID) error
	ListPurchaseOrdersByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.PurchaseOrder, error)

	// Opsional & Custom
	ApprovePurchaseOrder(ctx context.Context, poID uuid.UUID, approverID uuid.UUID) error
	ReceivePurchaseOrder(ctx context.Context, poID uuid.UUID, receiverID uuid.UUID) error
	CancelPurchaseOrder(ctx context.Context, poID uuid.UUID, reason string) error
	ListPurchaseOrdersBySupplierID(ctx context.Context, supplierID uuid.UUID) ([]*models.PurchaseOrder, error)
	GeneratePurchaseOrderReport(ctx context.Context, storeID uuid.UUID, from, to time.Time) ([]byte, error)
	ExportPurchaseOrders(ctx context.Context, storeID uuid.UUID, from, to time.Time) ([]byte, error)
	ListPurchaseOrdersByDateRange(ctx context.Context, storeID uuid.UUID, from, to time.Time) ([]*models.PurchaseOrder, error)
	GetPurchaseOrderHistory(ctx context.Context, poID uuid.UUID) ([]*models.PurchaseOrderHistory, error)
	NotifyPurchaseOrderStatusChange(ctx context.Context, poID uuid.UUID, status string) error
}
