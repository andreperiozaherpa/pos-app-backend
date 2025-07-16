package services

import (
	"context"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

// TransactionService menangani use case transaksi penjualan (POS).
type TransactionService interface {
	// CRUD Utama
	CreateTransaction(ctx context.Context, tx *models.Transaction) (uuid.UUID, error)
	GetTransactionByID(ctx context.Context, id uuid.UUID) (*models.Transaction, error)
	GetTransactionByCode(ctx context.Context, code string) (*models.Transaction, error)
	ListTransactionsByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.Transaction, error)

	// Method Opsional/Custom
	RefundTransaction(ctx context.Context, id uuid.UUID, amount float64) error
	ExportTransactions(ctx context.Context, storeID uuid.UUID, from, to time.Time) ([]byte, error)
	RecalculateTransactionTotals(ctx context.Context, id uuid.UUID) error
	GetTransactionAuditTrail(ctx context.Context, id uuid.UUID) ([]*models.TransactionAuditLog, error)
	ListTransactionsByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*models.Transaction, error)
	VoidTransaction(ctx context.Context, id uuid.UUID) error
	ApplyDiscountToTransaction(ctx context.Context, transactionID, discountID uuid.UUID) error
	PrintReceipt(ctx context.Context, transactionID uuid.UUID) ([]byte, error)
	ProcessPayment(ctx context.Context, transactionID uuid.UUID, paymentInfo *models.PaymentInfo) error
	ValidateTransactionStock(ctx context.Context, transactionID uuid.UUID) (bool, error)
	ListTransactionsByDateRange(ctx context.Context, storeID uuid.UUID, from, to time.Time) ([]*models.Transaction, error)
	ExportTransactionReceipts(ctx context.Context, storeID uuid.UUID, from, to time.Time) ([]byte, error)
	GetTransactionPaymentStatus(ctx context.Context, transactionID uuid.UUID) (string, error)
	ListTransactionRefunds(ctx context.Context, transactionID uuid.UUID) ([]*models.TransactionRefund, error)
	GetTransactionSummaryByDay(ctx context.Context, storeID uuid.UUID, date time.Time) (*models.TransactionSummary, error)
	NotifyTransactionStatusChange(ctx context.Context, transactionID uuid.UUID, status string) error
}
