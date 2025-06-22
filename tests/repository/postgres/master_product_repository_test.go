package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomMasterProductHelper adalah helper untuk membuat master product baru.
// Ini berbeda dengan createRandomMasterProduct di product_repository_test.go
// karena ini menggunakan repository yang sedang diuji.
func createRandomMasterProductHelper(t *testing.T) (*models.MasterProduct, *models.Company) {
	company := createRandomCompany(t)
	taxRate := createRandomTaxRate(t, company.ID) // Buat tax rate yang valid untuk company ini

	mp := &models.MasterProduct{
		ID:                uuid.New(),
		CompanyID:         company.ID,
		MasterProductCode: "MPC-" + randomString(6),
		Name:              "Master Product " + randomString(8),
		Description:       sql.NullString{String: "Description for master product", Valid: true},
		Category:          sql.NullString{String: "Electronics", Valid: true},
		UnitOfMeasure:     sql.NullString{String: "pcs", Valid: true},
		Barcode:           sql.NullString{String: "BARC-" + randomString(12), Valid: true}, // Perbaiki Barcode
		DefaultTaxRateID:  sql.NullInt32{Int32: taxRate.ID, Valid: true},                   // Gunakan ID tax rate yang baru dibuat
		ImageURL:          sql.NullString{String: "http://example.com/image.jpg", Valid: true},
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	err := masterProductTestRepo.Create(context.Background(), mp)
	if err != nil {
		t.Fatalf("Gagal membuat master product random untuk test: %v", err)
	}

	return mp, company
}

func TestMasterProductRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	newMP, _ := createRandomMasterProductHelper(t)

	foundMP, err := masterProductTestRepo.GetByID(context.Background(), newMP.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan master product by ID: %v", err)
	}
	if foundMP == nil {
		t.Fatal("Master product yang baru dibuat tidak ditemukan")
	}

	if newMP.ID != foundMP.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newMP.ID, foundMP.ID)
	}
	if newMP.Name != foundMP.Name {
		t.Errorf("Name tidak cocok. Diharapkan '%s', didapatkan '%s'", newMP.Name, foundMP.Name)
	}
	if newMP.MasterProductCode != foundMP.MasterProductCode {
		t.Errorf("MasterProductCode tidak cocok. Diharapkan '%s', didapatkan '%s'", newMP.MasterProductCode, foundMP.MasterProductCode)
	}
}

func TestMasterProductRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()
	nonExistentID := uuid.New()
	foundMP, err := masterProductTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundMP != nil {
		t.Errorf("Diharapkan master product nil, tetapi mendapatkan: %+v", foundMP)
	}
}

func TestMasterProductRepository_GetByCompanyAndCode(t *testing.T) {
	defer cleanup()

	newMP, company := createRandomMasterProductHelper(t)

	foundMP, err := masterProductTestRepo.GetByCompanyAndCode(context.Background(), company.ID, newMP.MasterProductCode)
	if err != nil {
		t.Fatalf("Gagal mendapatkan master product by CompanyID and Code: %v", err)
	}
	if foundMP == nil {
		t.Fatal("Master product tidak ditemukan berdasarkan CompanyID dan Code")
	}

	if newMP.ID != foundMP.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newMP.ID, foundMP.ID)
	}
	if newMP.MasterProductCode != foundMP.MasterProductCode {
		t.Errorf("MasterProductCode tidak cocok. Diharapkan '%s', didapatkan '%s'", newMP.MasterProductCode, foundMP.MasterProductCode)
	}

	// Test for non-existent code in existing company
	_, err = masterProductTestRepo.GetByCompanyAndCode(context.Background(), company.ID, "NONEXISTENTCODE")
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows untuk kode produk tidak ada, tetapi mendapatkan: %v", err)
	}
}

func TestMasterProductRepository_ListByCompanyID(t *testing.T) {
	defer cleanup()

	// Buat 2 master product untuk company 1
	_, company1 := createRandomMasterProductHelper(t)
	createRandomMasterProductForCompany(t, company1.ID)

	// Buat 1 master product untuk company 2
	_, _ = createRandomMasterProductHelper(t)

	// Test: Ambil master product untuk company 1
	masterProducts, err := masterProductTestRepo.ListByCompanyID(context.Background(), company1.ID)
	if err != nil {
		t.Fatalf("Gagal list master products by company ID: %v", err)
	}

	if len(masterProducts) != 2 {
		t.Errorf("Diharapkan 2 master products, tetapi mendapatkan %d", len(masterProducts))
	}
}

func TestMasterProductRepository_Update(t *testing.T) {
	defer cleanup()

	initialMP, _ := createRandomMasterProductHelper(t)

	updatedName := "Updated Master Product Name"
	initialMP.Name = updatedName
	initialMP.UpdatedAt = time.Now()

	err := masterProductTestRepo.Update(context.Background(), initialMP)
	if err != nil {
		t.Fatalf("Gagal mengupdate master product: %v", err)
	}

	foundMP, err := masterProductTestRepo.GetByID(context.Background(), initialMP.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan master product setelah update: %v", err)
	}
	if foundMP.Name != updatedName {
		t.Errorf("Name tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedName, foundMP.Name)
	}
}

func TestMasterProductRepository_Delete(t *testing.T) {
	defer cleanup()

	mpToDelete, _ := createRandomMasterProductHelper(t)

	err := masterProductTestRepo.Delete(context.Background(), mpToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus master product: %v", err)
	}

	_, err = masterProductTestRepo.GetByID(context.Background(), mpToDelete.ID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}

// createRandomMasterProductForCompany adalah helper spesifik untuk test list
func createRandomMasterProductForCompany(t *testing.T, companyID uuid.UUID) {
	mp := &models.MasterProduct{
		ID:                uuid.New(),
		CompanyID:         companyID,
		MasterProductCode: "MPC-" + randomString(6),
		Name:              "Another Master Product",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	err := masterProductTestRepo.Create(context.Background(), mp)
	if err != nil {
		t.Fatalf("Gagal membuat master product untuk company spesifik: %v", err)
	}
}
