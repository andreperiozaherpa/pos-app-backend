package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomTransaction adalah helper untuk membuat transaksi baru dengan item.
func createRandomTransaction(t *testing.T) *models.Transaction {
	// 1. Buat semua dependensi
	cashier := createRandomEmployee(t)
	customer := createRandomCustomer(t)
	storeProduct, _, _ := createRandomStoreProduct(t)

	// 2. Buat item transaksi
	item1 := models.TransactionItem{
		ID:                         uuid.New(),
		StoreProductID:             storeProduct.ID,
		Quantity:                   2,
		PricePerUnitAtTransaction:  storeProduct.SellingPrice,
		ItemSubtotalBeforeDiscount: storeProduct.SellingPrice * 2,
		ItemDiscountAmount:         0,
		ItemSubtotalAfterDiscount:  storeProduct.SellingPrice * 2,
		AppliedTaxRateID:           sql.NullInt32{},
		AppliedTaxRatePercentage:   sql.NullFloat64{},
		TaxAmountForItem:           0,
		ItemFinalTotal:             storeProduct.SellingPrice * 2,
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
	}

	// 3. Buat header transaksi
	finalTotal := item1.ItemFinalTotal
	transaction := &models.Transaction{
		ID:                             uuid.New(),
		TransactionCode:                "TRX-" + randomString(10),
		StoreID:                        cashier.StoreID.UUID,
		CashierEmployeeUserID:          cashier.UserID,
		CustomerUserID:                 uuid.NullUUID{UUID: customer.UserID, Valid: true},
		ActiveShiftID:                  uuid.NullUUID{},
		TransactionDate:                time.Now(),
		SubtotalAmount:                 finalTotal,
		TotalItemDiscountAmount:        0,
		SubtotalAfterItemDiscounts:     finalTotal,
		TransactionLevelDiscountAmount: 0,
		TaxableAmount:                  finalTotal,
		TotalTaxAmount:                 0,
		FinalTotalAmount:               finalTotal,
		ReceivedAmount:                 finalTotal + 10000,
		ChangeAmount:                   10000,
		PaymentMethod:                  sql.NullString{String: "CASH", Valid: true},
		Notes:                          sql.NullString{},
		CreatedAt:                      time.Now(),
		UpdatedAt:                      time.Now(),
		Items:                          []models.TransactionItem{item1},
	}

	// 4. Panggil repository untuk membuat transaksi
	err := transactionTestRepo.Create(context.Background(), transaction)
	if err != nil {
		t.Fatalf("Gagal membuat transaksi random untuk test: %v", err)
	}

	return transaction
}

func TestTransactionRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	newTx := createRandomTransaction(t)

	foundTx, err := transactionTestRepo.GetByID(context.Background(), newTx.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan transaksi by ID: %v", err)
	}
	if foundTx == nil {
		t.Fatal("Transaksi yang baru dibuat tidak ditemukan")
	}

	if newTx.ID != foundTx.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newTx.ID, foundTx.ID)
	}
	if newTx.TransactionCode != foundTx.TransactionCode {
		t.Errorf("TransactionCode tidak cocok. Diharapkan '%s', didapatkan '%s'", newTx.TransactionCode, foundTx.TransactionCode)
	}
	if len(foundTx.Items) != 1 {
		t.Errorf("Jumlah item tidak cocok. Diharapkan 1, didapatkan %d", len(foundTx.Items))
	}
	if foundTx.Items[0].StoreProductID != newTx.Items[0].StoreProductID {
		t.Errorf("StoreProductID pada item tidak cocok. Diharapkan '%s', didapatkan '%s'", newTx.Items[0].StoreProductID, foundTx.Items[0].StoreProductID)
	}
}

func TestTransactionRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()
	nonExistentID := uuid.New()
	foundTx, err := transactionTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundTx != nil {
		t.Errorf("Diharapkan transaksi nil, tetapi mendapatkan: %+v", foundTx)
	}
}
