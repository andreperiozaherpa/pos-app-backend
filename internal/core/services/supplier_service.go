package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// SupplierService mendefinisikan kontrak use case untuk manajemen supplier di aplikasi POS.
type SupplierService interface {
	// CreateSupplier membuat supplier baru.
	CreateSupplier(ctx context.Context, supplier *models.Supplier) (uuid.UUID, error)

	// GetSupplierByID mengambil supplier berdasarkan ID.
	GetSupplierByID(ctx context.Context, id uuid.UUID) (*models.Supplier, error)

	// UpdateSupplier memperbarui data supplier.
	UpdateSupplier(ctx context.Context, supplier *models.Supplier) error

	// DeleteSupplier menghapus supplier berdasarkan ID.
	DeleteSupplier(ctx context.Context, id uuid.UUID) error

	// ListSuppliersByCompanyID mengambil semua supplier pada company tertentu.
	ListSuppliersByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Supplier, error)

	// FindSupplierByPhoneOrEmail mencari supplier berdasarkan no hp/email (opsional).
	FindSupplierByPhoneOrEmail(ctx context.Context, phoneOrEmail string) (*models.Supplier, error)

	// DeactivateSupplier menonaktifkan supplier (opsional, soft delete).
	DeactivateSupplier(ctx context.Context, id uuid.UUID) error

	// SearchSuppliers pencarian supplier berdasarkan nama/kategori/kota (opsional).
	SearchSuppliers(ctx context.Context, query string, companyID uuid.UUID) ([]*models.Supplier, error)

	// BulkImportSuppliers import supplier dari file excel/CSV (opsional).
	BulkImportSuppliers(ctx context.Context, suppliers []*models.Supplier) error

	// GetSupplierOutstandingPOs menampilkan PO yang belum selesai/lunas dari supplier (custom).
	GetSupplierOutstandingPOs(ctx context.Context, supplierID uuid.UUID) ([]*models.PurchaseOrder, error)

	// UpdateSupplierStatus set supplier menjadi blacklist/prioritas (custom).
	UpdateSupplierStatus(ctx context.Context, supplierID uuid.UUID, status string) error

	// ArchiveSupplier mengarsipkan supplier (opsional, bukan delete).
	ArchiveSupplier(ctx context.Context, supplierID uuid.UUID) error

	// RestoreArchivedSupplier mengembalikan supplier dari arsip (opsional).
	RestoreArchivedSupplier(ctx context.Context, supplierID uuid.UUID) error

	// ApproveSupplier menyetujui supplier baru yang mendaftar (opsional).
	ApproveSupplier(ctx context.Context, supplierID uuid.UUID) error

	// GetSupplierTransactions mengambil daftar transaksi dengan supplier tersebut (opsional).
	GetSupplierTransactions(ctx context.Context, supplierID uuid.UUID) ([]*models.Transaction, error)

	// ExportSuppliers mengekspor data supplier ke file excel/CSV (opsional).
	ExportSuppliers(ctx context.Context, companyID uuid.UUID) ([]byte, error)

	// ListSupplierProducts menampilkan daftar produk yang disediakan supplier (opsional).
	ListSupplierProducts(ctx context.Context, supplierID uuid.UUID) ([]*models.MasterProduct, error)

	// BulkUpdateSupplierStatus mengubah status banyak supplier sekaligus (opsional).
	BulkUpdateSupplierStatus(ctx context.Context, supplierIDs []uuid.UUID, status string) error

	// GetSupplierContactHistory mengambil riwayat komunikasi dengan supplier (call/email/wa) (opsional).
	GetSupplierContactHistory(ctx context.Context, supplierID uuid.UUID) ([]*models.ContactHistory, error)
}
