package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomSupplier adalah helper untuk membuat supplier baru.
func createRandomSupplier(t *testing.T) (*models.Supplier, *models.Company) {
	company := createRandomCompany(t)

	supplier := &models.Supplier{
		ID:            uuid.New(),
		CompanyID:     company.ID,
		Name:          "Supplier " + randomString(8),
		ContactPerson: sql.NullString{String: "John Doe", Valid: true},
		Email:         sql.NullString{String: "supplier@example.com", Valid: true},
		PhoneNumber:   sql.NullString{String: "111-222-3333", Valid: true},
		Address:       sql.NullString{String: "123 Supplier Lane", Valid: true},
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err := supplierTestRepo.Create(context.Background(), supplier)
	if err != nil {
		t.Fatalf("Gagal membuat supplier random untuk test: %v", err)
	}

	return supplier, company
}

func TestSupplierRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	newSupplier, _ := createRandomSupplier(t)

	foundSupplier, err := supplierTestRepo.GetByID(context.Background(), newSupplier.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan supplier by ID: %v", err)
	}
	if foundSupplier == nil {
		t.Fatal("Supplier yang baru dibuat tidak ditemukan")
	}

	if newSupplier.ID != foundSupplier.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newSupplier.ID, foundSupplier.ID)
	}
	if newSupplier.Name != foundSupplier.Name {
		t.Errorf("Name tidak cocok. Diharapkan '%s', didapatkan '%s'", newSupplier.Name, foundSupplier.Name)
	}
}

func TestSupplierRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()
	nonExistentID := uuid.New()
	foundSupplier, err := supplierTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundSupplier != nil {
		t.Errorf("Diharapkan supplier nil, tetapi mendapatkan: %+v", foundSupplier)
	}
}

func TestSupplierRepository_ListByCompanyID(t *testing.T) {
	defer cleanup()

	// Buat 2 supplier untuk company 1
	_, company1 := createRandomSupplier(t)
	createRandomSupplierForCompany(t, company1.ID)

	// Buat 1 supplier untuk company 2
	_, _ = createRandomSupplier(t)

	// Test: Ambil supplier untuk company 1
	suppliers, err := supplierTestRepo.ListByCompanyID(context.Background(), company1.ID)
	if err != nil {
		t.Fatalf("Gagal list suppliers by company ID: %v", err)
	}

	if len(suppliers) != 2 {
		t.Errorf("Diharapkan 2 suppliers, tetapi mendapatkan %d", len(suppliers))
	}
}

func TestSupplierRepository_Update(t *testing.T) {
	defer cleanup()

	initialSupplier, _ := createRandomSupplier(t)

	updatedName := "Updated Supplier Name"
	initialSupplier.Name = updatedName
	initialSupplier.UpdatedAt = time.Now()

	err := supplierTestRepo.Update(context.Background(), initialSupplier)
	if err != nil {
		t.Fatalf("Gagal mengupdate supplier: %v", err)
	}

	foundSupplier, err := supplierTestRepo.GetByID(context.Background(), initialSupplier.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan supplier setelah update: %v", err)
	}
	if foundSupplier.Name != updatedName {
		t.Errorf("Name tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedName, foundSupplier.Name)
	}
}

func TestSupplierRepository_Delete(t *testing.T) {
	defer cleanup()

	supplierToDelete, _ := createRandomSupplier(t)

	err := supplierTestRepo.Delete(context.Background(), supplierToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus supplier: %v", err)
	}

	_, err = supplierTestRepo.GetByID(context.Background(), supplierToDelete.ID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}

// createRandomSupplierForCompany adalah helper spesifik untuk test list
func createRandomSupplierForCompany(t *testing.T, companyID uuid.UUID) {
	s := &models.Supplier{ID: uuid.New(), CompanyID: companyID, Name: "Another Supplier", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := supplierTestRepo.Create(context.Background(), s)
	if err != nil {
		t.Fatalf("Gagal membuat supplier untuk company spesifik: %v", err)
	}
}
