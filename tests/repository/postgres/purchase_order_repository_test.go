package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomPurchaseOrder adalah helper untuk membuat purchase order baru dengan item.
func createRandomPurchaseOrder(t *testing.T) *models.PurchaseOrder {
	// 1. Buat semua dependensi
	store := createRandomStore(t, createTestCompanyAndBusinessLine(t))
	supplier, _ := createRandomSupplier(t)
	masterProduct := createRandomMasterProduct(t, supplier.CompanyID)
	user := createRandomEmployee(t) // User yang membuat PO

	// 2. Buat item purchase order
	item1 := models.PurchaseOrderItem{
		ID:                   uuid.New(),
		MasterProductID:      masterProduct.ID,
		QuantityOrdered:      10,
		PurchasePricePerUnit: 50.00,
		QuantityReceived:     0,
		Subtotal:             500.00,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	// 3. Buat header purchase order
	totalAmount := item1.Subtotal
	po := &models.PurchaseOrder{
		ID:                   uuid.New(),
		StoreID:              store.ID,
		SupplierID:           supplier.ID,
		OrderDate:            time.Now(),
		ExpectedDeliveryDate: sql.NullTime{Time: time.Now().AddDate(0, 0, 7), Valid: true},
		Status:               models.POStatusPending,
		TotalAmount:          sql.NullFloat64{Float64: totalAmount, Valid: true},
		Notes:                sql.NullString{String: "Initial order", Valid: true},
		CreatedByUserID:      user.UserID,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
		Items:                []models.PurchaseOrderItem{item1},
	}

	// 4. Panggil repository untuk membuat purchase order
	err := purchaseOrderTestRepo.Create(context.Background(), po)
	if err != nil {
		t.Fatalf("Gagal membuat purchase order random untuk test: %v", err)
	}

	return po
}

func TestPurchaseOrderRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	newPO := createRandomPurchaseOrder(t)

	foundPO, err := purchaseOrderTestRepo.GetByID(context.Background(), newPO.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan purchase order by ID: %v", err)
	}
	if foundPO == nil {
		t.Fatal("Purchase order yang baru dibuat tidak ditemukan")
	}

	if newPO.ID != foundPO.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newPO.ID, foundPO.ID)
	}
	if newPO.StoreID != foundPO.StoreID {
		t.Errorf("StoreID tidak cocok. Diharapkan '%s', didapatkan '%s'", newPO.StoreID, foundPO.StoreID)
	}
	if len(foundPO.Items) != 1 {
		t.Errorf("Jumlah item tidak cocok. Diharapkan 1, didapatkan %d", len(foundPO.Items))
	}
	if foundPO.Items[0].MasterProductID != newPO.Items[0].MasterProductID {
		t.Errorf("MasterProductID pada item tidak cocok. Diharapkan '%s', didapatkan '%s'", newPO.Items[0].MasterProductID, foundPO.Items[0].MasterProductID)
	}
}

func TestPurchaseOrderRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()
	nonExistentID := uuid.New()
	foundPO, err := purchaseOrderTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundPO != nil {
		t.Errorf("Diharapkan purchase order nil, tetapi mendapatkan: %+v", foundPO)
	}
}

func TestPurchaseOrderRepository_ListByStoreID(t *testing.T) {
	defer cleanup()

	// Buat PO untuk store1
	po1 := createRandomPurchaseOrder(t)
	store1ID := po1.StoreID

	// Buat PO lain untuk store1
	po2 := createRandomPurchaseOrder(t)
	po2.StoreID = store1ID // Pastikan PO ini juga untuk store1
	err := purchaseOrderTestRepo.Update(context.Background(), po2)
	if err != nil {
		t.Fatalf("Gagal mengupdate PO2 untuk store1: %v", err)
	}

	// Buat PO untuk store lain
	createRandomPurchaseOrder(t)

	pos, err := purchaseOrderTestRepo.ListByStoreID(context.Background(), store1ID)
	if err != nil {
		t.Fatalf("Gagal list purchase orders by store ID: %v", err)
	}

	if len(pos) != 2 {
		t.Errorf("Diharapkan 2 purchase orders, tetapi mendapatkan %d", len(pos))
	}
}

func TestPurchaseOrderRepository_Update(t *testing.T) {
	defer cleanup()

	initialPO := createRandomPurchaseOrder(t)

	updatedStatus := models.POStatusOrdered
	updatedNotes := sql.NullString{String: "Order has been placed with supplier.", Valid: true}
	initialPO.Status = updatedStatus
	initialPO.Notes = updatedNotes
	initialPO.UpdatedAt = time.Now()

	err := purchaseOrderTestRepo.Update(context.Background(), initialPO)
	if err != nil {
		t.Fatalf("Gagal mengupdate purchase order: %v", err)
	}

	foundPO, err := purchaseOrderTestRepo.GetByID(context.Background(), initialPO.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan purchase order setelah update: %v", err)
	}
	if foundPO.Status != updatedStatus {
		t.Errorf("Status tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedStatus, foundPO.Status)
	}
	if foundPO.Notes.String != updatedNotes.String {
		t.Errorf("Notes tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedNotes.String, foundPO.Notes.String)
	}
}

func TestPurchaseOrderRepository_Delete(t *testing.T) {
	defer cleanup()

	poToDelete := createRandomPurchaseOrder(t)

	err := purchaseOrderTestRepo.Delete(context.Background(), poToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus purchase order: %v", err)
	}

	_, err = purchaseOrderTestRepo.GetByID(context.Background(), poToDelete.ID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}
