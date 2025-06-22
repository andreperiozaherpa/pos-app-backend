package repository_test

import (
	"context"
	"testing"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

func TestAppliedItemDiscountRepository_CreateAndListByTransactionItemID(t *testing.T) {
	defer cleanup()

	// 1. Create dependencies: a transaction with an item, and two discounts
	transaction := createRandomTransaction(t)
	if len(transaction.Items) == 0 {
		t.Fatal("Helper createRandomTransaction tidak menghasilkan item transaksi")
	}
	transactionItem := transaction.Items[0]

	discount1, _ := createRandomDiscount(t)
	discount2, _ := createRandomDiscount(t)

	// 2. Define the discounts to be applied to the item
	discountsToApply := []models.AppliedItemDiscount{
		{
			TransactionItemID:           transactionItem.ID,
			DiscountID:                  discount1.ID,
			AppliedDiscountAmountOnItem: 500.00,
		},
		{
			TransactionItemID:           transactionItem.ID,
			DiscountID:                  discount2.ID,
			AppliedDiscountAmountOnItem: 1000.00,
		},
	}

	// 3. Use the repository to create the applied discounts
	err := appliedItemDiscountTestRepo.Create(context.Background(), discountsToApply)
	if err != nil {
		t.Fatalf("Gagal membuat applied item discounts: %v", err)
	}

	// 4. List the discounts by the transaction item's ID
	foundDiscounts, err := appliedItemDiscountTestRepo.ListByTransactionItemID(context.Background(), transactionItem.ID)
	if err != nil {
		t.Fatalf("Gagal list applied item discounts by transaction item ID: %v", err)
	}

	if len(foundDiscounts) != 2 {
		t.Errorf("Diharapkan 2 applied item discounts, tetapi mendapatkan %d", len(foundDiscounts))
	}

	// Verify the data
	foundMap := make(map[uuid.UUID]*models.AppliedItemDiscount)
	for _, d := range foundDiscounts {
		foundMap[d.DiscountID] = d
	}

	if _, ok := foundMap[discount1.ID]; !ok {
		t.Errorf("Diskon 1 tidak ditemukan di daftar")
	}
	if foundMap[discount1.ID].AppliedDiscountAmountOnItem != 500.00 {
		t.Errorf("AppliedDiscountAmountOnItem untuk diskon 1 tidak cocok. Diharapkan 500.00, didapatkan %f", foundMap[discount1.ID].AppliedDiscountAmountOnItem)
	}
}

func TestAppliedItemDiscountRepository_ListByTransactionItemID_Empty(t *testing.T) {
	defer cleanup()

	// Create a transaction item but don't apply any discounts
	transaction := createRandomTransaction(t)
	transactionItem := transaction.Items[0]

	foundDiscounts, err := appliedItemDiscountTestRepo.ListByTransactionItemID(context.Background(), transactionItem.ID)
	if err != nil {
		t.Fatalf("Gagal list applied item discounts: %v", err)
	}

	if len(foundDiscounts) != 0 {
		t.Errorf("Diharapkan 0 applied item discounts, tetapi mendapatkan %d", len(foundDiscounts))
	}
}

func TestAppliedItemDiscountRepository_DeleteByTransactionItemID(t *testing.T) {
	defer cleanup()

	// 1. Create dependencies
	transaction := createRandomTransaction(t)
	transactionItem := transaction.Items[0]
	discount, _ := createRandomDiscount(t)

	// 2. Apply a discount
	discountsToApply := []models.AppliedItemDiscount{
		{TransactionItemID: transactionItem.ID, DiscountID: discount.ID, AppliedDiscountAmountOnItem: 1000},
	}
	err := appliedItemDiscountTestRepo.Create(context.Background(), discountsToApply)
	if err != nil {
		t.Fatalf("Gagal membuat applied item discount untuk test delete: %v", err)
	}

	// 3. Delete by transaction item ID
	err = appliedItemDiscountTestRepo.DeleteByTransactionItemID(context.Background(), transactionItem.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus applied item discounts: %v", err)
	}

	// 4. Verify they are deleted
	remainingDiscounts, err := appliedItemDiscountTestRepo.ListByTransactionItemID(context.Background(), transactionItem.ID)
	if err != nil {
		t.Fatalf("Gagal list sisa applied item discounts setelah delete: %v", err)
	}
	if len(remainingDiscounts) != 0 {
		t.Errorf("Diharapkan 0 applied item discounts setelah delete, tetapi mendapatkan %d", len(remainingDiscounts))
	}
}
