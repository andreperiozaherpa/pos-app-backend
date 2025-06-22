package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
)

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
