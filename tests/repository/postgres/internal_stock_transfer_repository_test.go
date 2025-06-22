package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomInternalStockTransfer adalah helper untuk membuat internal stock transfer baru dengan item.
func createRandomInternalStockTransfer(t *testing.T) *models.InternalStockTransfer {
	// 1. Buat semua dependensi
	company := createRandomCompany(t)
	businessLineID := createTestCompanyAndBusinessLine(t) // This helper creates a company too, might be redundant
	sourceStore := createRandomStore(t, businessLineID)
	destinationStore := createRandomStore(t, businessLineID)
	storeProduct, _, _ := createRandomStoreProduct(t) // This helper creates its own dependencies
	requestedByUser := createRandomEmployee(t)

	// Ensure sourceStore and destinationStore belong to the same company
	// For simplicity in test, we'll just use the company from the first store
	companyID := company.ID

	// 2. Buat item transfer
	item1 := models.InternalStockTransferItem{
		ID:                   uuid.New(),
		SourceStoreProductID: storeProduct.ID,
		QuantityRequested:    5,
		QuantityShipped:      0,
		QuantityReceived:     0,
		Notes:                sql.NullString{String: "Item for transfer", Valid: true},
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	// 3. Buat header transfer
	ist := &models.InternalStockTransfer{
		ID:                 uuid.New(),
		TransferCode:       "IST-" + randomString(8),
		CompanyID:          companyID,
		SourceStoreID:      sourceStore.ID,
		DestinationStoreID: destinationStore.ID,
		TransferDate:       time.Now(),
		Status:             models.StockTransferStatusPending,
		Notes:              sql.NullString{String: "Transfer request", Valid: true},
		RequestedByUserID:  uuid.NullUUID{UUID: requestedByUser.UserID, Valid: true},
		ApprovedByUserID:   uuid.NullUUID{},
		ShippedByUserID:    uuid.NullUUID{},
		ReceivedByUserID:   uuid.NullUUID{},
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		Items:              []models.InternalStockTransferItem{item1},
	}

	// 4. Panggil repository untuk membuat transfer
	err := internalStockTransferTestRepo.Create(context.Background(), ist)
	if err != nil {
		t.Fatalf("Gagal membuat internal stock transfer random untuk test: %v", err)
	}

	return ist
}

func TestInternalStockTransferRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	newIST := createRandomInternalStockTransfer(t)

	foundIST, err := internalStockTransferTestRepo.GetByID(context.Background(), newIST.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan internal stock transfer by ID: %v", err)
	}
	if foundIST == nil {
		t.Fatal("Internal stock transfer yang baru dibuat tidak ditemukan")
	}

	if newIST.ID != foundIST.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newIST.ID, foundIST.ID)
	}
	if newIST.TransferCode != foundIST.TransferCode {
		t.Errorf("TransferCode tidak cocok. Diharapkan '%s', didapatkan '%s'", newIST.TransferCode, foundIST.TransferCode)
	}
	if len(foundIST.Items) != 1 {
		t.Errorf("Jumlah item tidak cocok. Diharapkan 1, didapatkan %d", len(foundIST.Items))
	}
	if foundIST.Items[0].SourceStoreProductID != newIST.Items[0].SourceStoreProductID {
		t.Errorf("SourceStoreProductID pada item tidak cocok. Diharapkan '%s', didapatkan '%s'", newIST.Items[0].SourceStoreProductID, foundIST.Items[0].SourceStoreProductID)
	}
}

func TestInternalStockTransferRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()
	nonExistentID := uuid.New()
	foundIST, err := internalStockTransferTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundIST != nil {
		t.Errorf("Diharapkan internal stock transfer nil, tetapi mendapatkan: %+v", foundIST)
	}
}

func TestInternalStockTransferRepository_ListByCompanyID(t *testing.T) {
	defer cleanup()

	// Buat IST untuk company1
	ist1 := createRandomInternalStockTransfer(t)
	company1ID := ist1.CompanyID

	// Buat IST lain untuk company1
	ist2 := createRandomInternalStockTransfer(t)
	ist2.CompanyID = company1ID // Pastikan IST ini juga untuk company1
	err := internalStockTransferTestRepo.Update(context.Background(), ist2)
	if err != nil {
		t.Fatalf("Gagal mengupdate IST2 untuk company1: %v", err)
	}

	// Buat IST untuk company lain
	createRandomInternalStockTransfer(t)

	ists, err := internalStockTransferTestRepo.ListByCompanyID(context.Background(), company1ID)
	if err != nil {
		t.Fatalf("Gagal list internal stock transfers by company ID: %v", err)
	}

	if len(ists) != 2 {
		t.Errorf("Diharapkan 2 internal stock transfers, tetapi mendapatkan %d", len(ists))
	}
}

func TestInternalStockTransferRepository_Update(t *testing.T) {
	defer cleanup()

	initialIST := createRandomInternalStockTransfer(t)

	updatedStatus := models.StockTransferStatusApproved
	updatedNotes := sql.NullString{String: "Approved by manager.", Valid: true}
	initialIST.Status = updatedStatus
	initialIST.Notes = updatedNotes
	initialIST.UpdatedAt = time.Now()

	err := internalStockTransferTestRepo.Update(context.Background(), initialIST)
	if err != nil {
		t.Fatalf("Gagal mengupdate internal stock transfer: %v", err)
	}

	foundIST, err := internalStockTransferTestRepo.GetByID(context.Background(), initialIST.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan internal stock transfer setelah update: %v", err)
	}
	if foundIST.Status != updatedStatus {
		t.Errorf("Status tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedStatus, foundIST.Status)
	}
	if foundIST.Notes.String != updatedNotes.String {
		t.Errorf("Notes tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedNotes.String, foundIST.Notes.String)
	}
}

func TestInternalStockTransferRepository_Delete(t *testing.T) {
	defer cleanup()

	istToDelete := createRandomInternalStockTransfer(t)

	err := internalStockTransferTestRepo.Delete(context.Background(), istToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus internal stock transfer: %v", err)
	}

	_, err = internalStockTransferTestRepo.GetByID(context.Background(), istToDelete.ID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}
