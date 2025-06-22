package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

func TestStoreProductRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	newProduct, _, _ := createRandomStoreProduct(t)

	foundProduct, err := storeProductTestRepo.GetByID(context.Background(), newProduct.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan produk by ID: %v", err)
	}
	if foundProduct == nil {
		t.Fatal("Produk yang baru dibuat tidak ditemukan")
	}

	if newProduct.ID != foundProduct.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newProduct.ID, foundProduct.ID)
	}
}

func TestStoreProductRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()
	nonExistentID := uuid.New()
	foundProduct, err := storeProductTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundProduct != nil {
		t.Errorf("Diharapkan produk nil, tetapi mendapatkan: %+v", foundProduct)
	}
}

func TestStoreProductRepository_GetByStoreAndMasterProduct(t *testing.T) {
	defer cleanup()

	newProduct, store, _ := createRandomStoreProduct(t)

	foundProduct, err := storeProductTestRepo.GetByStoreAndMasterProduct(context.Background(), store.ID, newProduct.MasterProductID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan produk by StoreID and MasterProductID: %v", err)
	}
	if foundProduct == nil {
		t.Fatal("Produk tidak ditemukan berdasarkan StoreID dan MasterProductID")
	}

	if newProduct.ID != foundProduct.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newProduct.ID, foundProduct.ID)
	}

	// Test for non-existent product code in existing store
	_, err = storeProductTestRepo.GetByStoreAndMasterProduct(context.Background(), store.ID, uuid.New())
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows untuk kode produk tidak ada, tetapi mendapatkan: %v", err)
	}
}

func TestStoreProductRepository_ListByStore(t *testing.T) {
	defer cleanup()

	businessLineID := createTestCompanyAndBusinessLine(t)
	store1 := createRandomStore(t, businessLineID)
	store2 := createRandomStore(t, businessLineID)

	// Buat 2 produk untuk store1
	createRandomProductForStore(t, store1.ID)
	createRandomProductForStore(t, store1.ID)

	// Buat 1 produk untuk store2
	createRandomProductForStore(t, store2.ID)

	products, err := storeProductTestRepo.ListByStore(context.Background(), store1.ID)
	if err != nil {
		t.Fatalf("Gagal melakukan list products by store ID: %v", err)
	}

	if len(products) != 2 {
		t.Errorf("Diharapkan 2 produk, tetapi mendapatkan %d", len(products))
	}
}

func TestStoreProductRepository_Update(t *testing.T) {
	defer cleanup()

	initialProduct, _, _ := createRandomStoreProduct(t)

	initialProduct.Stock = 100            // Update stock
	initialProduct.PurchasePrice = 110.00 // Update purchase price

	err := storeProductTestRepo.Update(context.Background(), initialProduct)
	if err != nil {
		t.Fatalf("Gagal mengupdate produk: %v", err)
	}

	foundProduct, err := storeProductTestRepo.GetByID(context.Background(), initialProduct.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan produk setelah update: %v", err)
	}
	if foundProduct.Stock != 100 {
		t.Errorf("Stock tidak terupdate. Diharapkan '%d', didapatkan '%d'", 100, foundProduct.Stock)
	}
	if foundProduct.PurchasePrice != 110.00 {
		t.Errorf("PurchasePrice tidak terupdate. Diharapkan '%.2f', didapatkan '%.2f'", 110.00, foundProduct.PurchasePrice)
	}
}

func TestStoreProductRepository_Delete(t *testing.T) {
	defer cleanup()

	productToDelete, _, _ := createRandomStoreProduct(t)

	err := storeProductTestRepo.Delete(context.Background(), productToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus produk: %v", err)
	}

	_, err = storeProductTestRepo.GetByID(context.Background(), productToDelete.ID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}

// createRandomProductForStore adalah helper spesifik untuk test list
func createRandomProductForStore(t *testing.T, storeID uuid.UUID) {
	supplier, _ := createRandomSupplier(t)
	masterProduct := createRandomMasterProduct(t, supplier.CompanyID)
	p := &models.StoreProduct{
		ID:               uuid.New(),
		MasterProductID:  masterProduct.ID,
		StoreID:          storeID,
		SupplierID:       uuid.NullUUID{UUID: supplier.ID, Valid: true},
		StoreSpecificSKU: sql.NullString{String: "SKU-" + randomString(6), Valid: true},
		PurchasePrice:    10.00,
		SellingPrice:     20.00,
		Stock:            5,
		CreatedAt:        time.Now(), // Removed Name, Description, Category, UnitOfMeasure, Barcode
		UpdatedAt:        time.Now(),
	}
	err := storeProductTestRepo.Create(context.Background(), p)
	if err != nil {
		t.Fatalf("Gagal membuat produk untuk store spesifik: %v", err)
	}
}
