package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomStockMovement adalah helper untuk membuat stock movement baru.
func createRandomStockMovement(t *testing.T) (*models.StockMovement, *models.StoreProduct) {
	// 1. Buat dependensi
	storeProduct, store, _ := createRandomStoreProduct(t)
	user := createRandomEmployee(t)

	// 2. Buat stock movement
	movement := &models.StockMovement{
		ID:              uuid.New(),
		StoreProductID:  storeProduct.ID,
		MovementType:    models.MovementTypeSale,
		StoreID:         store.ID,
		QuantityChanged: -2, // Penjualan mengurangi stok
		MovementDate:    time.Now(),
		ReferenceID:     uuid.NullUUID{UUID: uuid.New(), Valid: true}, // Misal ID transaksi
		ReferenceType:   sql.NullString{String: "TRANSACTION_ITEM", Valid: true},
		Notes:           sql.NullString{String: "Sale transaction", Valid: true},
		CreatedByUserID: uuid.NullUUID{UUID: user.UserID, Valid: true},
		CreatedAt:       time.Now(),
	}

	// 3. Panggil repository untuk membuat movement
	err := stockMovementTestRepo.Create(context.Background(), movement)
	if err != nil {
		t.Fatalf("Gagal membuat stock movement random untuk test: %v", err)
	}

	return movement, storeProduct
}

func TestStockMovementRepository_CreateAndListByStoreProduct(t *testing.T) {
	defer cleanup()

	newMovement, storeProduct := createRandomStockMovement(t)

	// Buat movement lain untuk produk yang berbeda, untuk memastikan list hanya mengembalikan yang benar
	createRandomStockMovement(t)

	movements, err := stockMovementTestRepo.ListByStoreProduct(context.Background(), storeProduct.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan list stock movement by StoreProductID: %v", err)
	}

	if len(movements) != 1 {
		t.Fatalf("Diharapkan 1 movement, tetapi mendapatkan %d", len(movements))
	}

	foundMovement := movements[0]
	if newMovement.ID != foundMovement.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newMovement.ID, foundMovement.ID)
	}
	if newMovement.MovementType != foundMovement.MovementType {
		t.Errorf("MovementType tidak cocok. Diharapkan '%s', didapatkan '%s'", newMovement.MovementType, foundMovement.MovementType)
	}
	if newMovement.QuantityChanged != foundMovement.QuantityChanged {
		t.Errorf("QuantityChanged tidak cocok. Diharapkan %d, didapatkan %d", newMovement.QuantityChanged, foundMovement.QuantityChanged)
	}
}

func TestStockMovementRepository_ListByStoreProduct_Empty(t *testing.T) {
	defer cleanup()

	// Buat produk tapi jangan buat movement untuknya
	storeProduct, _, _ := createRandomStoreProduct(t)

	movements, err := stockMovementTestRepo.ListByStoreProduct(context.Background(), storeProduct.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan list stock movement by StoreProductID: %v", err)
	}

	if len(movements) != 0 {
		t.Errorf("Diharapkan 0 movements, tetapi mendapatkan %d", len(movements))
	}
}
