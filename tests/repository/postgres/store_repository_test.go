package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestStoreRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	businessLineID := createTestCompanyAndBusinessLine(t)
	newStore := createRandomStore(t, businessLineID)

	foundStore, err := storeTestRepo.GetByID(context.Background(), newStore.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan store by ID: %v", err)
	}
	if foundStore == nil {
		t.Fatal("Store yang baru dibuat tidak ditemukan")
	}

	if newStore.ID != foundStore.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newStore.ID, foundStore.ID)
	}
	if newStore.Name != foundStore.Name {
		t.Errorf("Name tidak cocok. Diharapkan '%s', didapatkan '%s'", newStore.Name, foundStore.Name)
	}
	if newStore.BusinessLineID != foundStore.BusinessLineID {
		t.Errorf("BusinessLineID tidak cocok. Diharapkan '%s', didapatkan '%s'", newStore.BusinessLineID, foundStore.BusinessLineID)
	}
}

func TestStoreRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()

	nonExistentID := uuid.New()
	foundStore, err := storeTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundStore != nil {
		t.Errorf("Diharapkan store nil, tetapi mendapatkan: %+v", foundStore)
	}
}

func TestStoreRepository_Update(t *testing.T) {
	defer cleanup()

	businessLineID := createTestCompanyAndBusinessLine(t)
	initialStore := createRandomStore(t, businessLineID)

	updatedName := "Updated Store Name"
	updatedAddress := "999 Updated Ave"
	initialStore.Name = updatedName
	initialStore.Address = sql.NullString{String: updatedAddress, Valid: true}
	initialStore.UpdatedAt = time.Now()

	err := storeTestRepo.Update(context.Background(), initialStore)
	if err != nil {
		t.Fatalf("Gagal mengupdate store: %v", err)
	}

	foundStore, err := storeTestRepo.GetByID(context.Background(), initialStore.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan store setelah update: %v", err)
	}

	if foundStore.Name != updatedName {
		t.Errorf("Name tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedName, foundStore.Name)
	}
	if foundStore.Address.String != updatedAddress {
		t.Errorf("Address tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedAddress, foundStore.Address.String)
	}
}

func TestStoreRepository_Delete(t *testing.T) {
	defer cleanup()

	businessLineID := createTestCompanyAndBusinessLine(t)
	storeToDelete := createRandomStore(t, businessLineID)

	err := storeTestRepo.Delete(context.Background(), storeToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus store: %v", err)
	}

	_, err = storeTestRepo.GetByID(context.Background(), storeToDelete.ID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}

func TestStoreRepository_ListByBusinessLine(t *testing.T) {
	defer cleanup()

	blID1 := createTestCompanyAndBusinessLine(t)
	blID2 := createTestCompanyAndBusinessLine(t)

	createRandomStore(t, blID1)
	createRandomStore(t, blID1)
	createRandomStore(t, blID2)

	stores, err := storeTestRepo.ListByBusinessLine(context.Background(), blID1)
	if err != nil {
		t.Fatalf("Gagal melakukan list stores by business line: %v", err)
	}

	if len(stores) != 2 {
		t.Errorf("Diharapkan 2 store, tetapi mendapatkan %d", len(stores))
	}
}
