package repository_test

import (
	"context"
	"testing"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomTransactionDiscount is a helper to create a random transaction discount.
// createRandomAppliedTransactionDiscount is a helper to create a random applied transaction discount.
// It requires an existing transaction and discount.
func createRandomAppliedTransactionDiscount(t *testing.T, transactionID uuid.UUID, discountID uuid.UUID) *models.AppliedTransactionDiscount {
	td := &models.AppliedTransactionDiscount{
		TransactionID:                      transactionID,
		DiscountID:                         discountID,
		AppliedDiscountAmountOnTransaction: 15000.00, // Example amount
	}

	err := appliedTransactionDiscountTestRepo.Create(context.Background(), []models.AppliedTransactionDiscount{*td})
	if err != nil {
		t.Fatalf("Failed to create random transaction discount for test: %v", err)
	}
	return td
}

func TestAppliedTransactionDiscountRepository_CreateAndListByTransactionID(t *testing.T) {
	defer cleanup()

	// 1. Create dependencies: a transaction and a discount
	transaction := createRandomTransaction(t) // This helper creates a full transaction
	discount, _ := createRandomDiscount(t)    // This helper creates a discount

	// 2. Create a transaction discount
	td1 := models.AppliedTransactionDiscount{
		TransactionID:                      transaction.ID,
		DiscountID:                         discount.ID,
		AppliedDiscountAmountOnTransaction: 10000.00,
	}

	// 3. Create another discount and apply it to the same transaction
	discount2, _ := createRandomDiscount(t)
	td2 := models.AppliedTransactionDiscount{
		TransactionID:                      transaction.ID,
		DiscountID:                         discount2.ID,
		AppliedDiscountAmountOnTransaction: 5000.00,
	}

	// 4. Use the repository to create multiple transaction discounts
	err := appliedTransactionDiscountTestRepo.Create(context.Background(), []models.AppliedTransactionDiscount{td1, td2})
	if err != nil {
		t.Fatalf("Failed to create transaction discounts: %v", err)
	}

	// 5. List transaction discounts by transaction ID
	foundDiscounts, err := appliedTransactionDiscountTestRepo.ListByTransactionID(context.Background(), transaction.ID)
	if err != nil {
		t.Fatalf("Failed to list transaction discounts by transaction ID: %v", err)
	}

	if len(foundDiscounts) != 2 {
		t.Errorf("Expected 2 applied transaction discounts, got %d", len(foundDiscounts))
	}

	// Verify the data
	// Since order is not guaranteed, we check for existence and values
	foundMap := make(map[uuid.UUID]*models.AppliedTransactionDiscount)
	for _, td := range foundDiscounts {
		foundMap[td.DiscountID] = td
	}

	if _, ok := foundMap[td1.DiscountID]; !ok {
		t.Errorf("AppliedTransactionDiscount 1 not found in listed discounts")
	}
	if foundMap[td1.DiscountID].AppliedDiscountAmountOnTransaction != td1.AppliedDiscountAmountOnTransaction {
		t.Errorf("AppliedDiscountAmountOnTransaction for td1 mismatch. Expected %f, got %f", td1.AppliedDiscountAmountOnTransaction, foundMap[td1.DiscountID].AppliedDiscountAmountOnTransaction)
	}

	if _, ok := foundMap[td2.DiscountID]; !ok {
		t.Errorf("AppliedTransactionDiscount 2 not found in listed discounts")
	}
	if foundMap[td2.DiscountID].AppliedDiscountAmountOnTransaction != td2.AppliedDiscountAmountOnTransaction {
		t.Errorf("AppliedDiscountAmountOnTransaction for td2 mismatch. Expected %f, got %f", td2.AppliedDiscountAmountOnTransaction, foundMap[td2.DiscountID].AppliedDiscountAmountOnTransaction)
	}
}

func TestAppliedTransactionDiscountRepository_ListByTransactionID_Empty(t *testing.T) {
	defer cleanup()

	// Create a transaction but no discounts for it
	transaction := createRandomTransaction(t)

	foundDiscounts, err := appliedTransactionDiscountTestRepo.ListByTransactionID(context.Background(), transaction.ID)
	if err != nil {
		t.Fatalf("Failed to list transaction discounts by transaction ID: %v", err)
	}

	if len(foundDiscounts) != 0 {
		t.Errorf("Expected 0 applied transaction discounts, got %d", len(foundDiscounts))
	}
}

func TestAppliedTransactionDiscountRepository_DeleteByTransactionID(t *testing.T) {
	defer cleanup()

	// 1. Create dependencies
	transaction := createRandomTransaction(t)
	discount1, _ := createRandomDiscount(t)
	discount2, _ := createRandomDiscount(t)

	// 2. Create multiple transaction discounts for the same transaction
	tds := []models.AppliedTransactionDiscount{
		{TransactionID: transaction.ID, DiscountID: discount1.ID, AppliedDiscountAmountOnTransaction: 1000},
		{TransactionID: transaction.ID, DiscountID: discount2.ID, AppliedDiscountAmountOnTransaction: 2000},
	}
	err := appliedTransactionDiscountTestRepo.Create(context.Background(), tds)
	if err != nil {
		t.Fatalf("Failed to create transaction discounts for deletion test: %v", err)
	}

	// 3. Delete by transaction ID
	err = appliedTransactionDiscountTestRepo.DeleteByTransactionID(context.Background(), transaction.ID)
	if err != nil {
		t.Fatalf("Failed to delete transaction discounts by transaction ID: %v", err)
	}

	// 4. Verify they are deleted
	remainingDiscounts, err := appliedTransactionDiscountTestRepo.ListByTransactionID(context.Background(), transaction.ID)
	if err != nil {
		t.Fatalf("Failed to list remaining transaction discounts after deletion: %v", err)
	}
	if len(remainingDiscounts) != 0 {
		t.Errorf("Expected 0 applied transaction discounts after deletion, got %d", len(remainingDiscounts))
	}
}
