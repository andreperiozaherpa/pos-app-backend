package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomBusinessLine adalah helper untuk membuat business line baru.
// Ini membutuhkan company yang sudah ada, jadi kita panggil createRandomCompany dulu.
func createRandomBusinessLine(t *testing.T) (*models.BusinessLine, *models.Company) {
	company := createRandomCompany(t)

	bl := &models.BusinessLine{
		ID:          uuid.New(),
		CompanyID:   company.ID,
		Name:        "Business Line " + uuid.NewString(),
		Description: sql.NullString{String: "Test Description", Valid: true},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := businessLineTestRepo.Create(context.Background(), bl)
	if err != nil {
		t.Fatalf("Gagal membuat business line random untuk test: %v", err)
	}

	return bl, company
}

func TestBusinessLineRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	newBL, _ := createRandomBusinessLine(t)

	foundBL, err := businessLineTestRepo.GetByID(context.Background(), newBL.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan business line by ID: %v", err)
	}
	if foundBL == nil {
		t.Fatal("Business line yang baru dibuat tidak ditemukan")
	}

	if newBL.ID != foundBL.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newBL.ID, foundBL.ID)
	}
	if newBL.Name != foundBL.Name {
		t.Errorf("Name tidak cocok. Diharapkan '%s', didapatkan '%s'", newBL.Name, foundBL.Name)
	}
}

func TestBusinessLineRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()
	nonExistentID := uuid.New()
	foundBL, err := businessLineTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundBL != nil {
		t.Errorf("Diharapkan business line nil, tetapi mendapatkan: %+v", foundBL)
	}
}

func TestBusinessLineRepository_ListByCompanyID(t *testing.T) {
	defer cleanup()

	// Buat 2 business line untuk company 1
	_, company1 := createRandomBusinessLine(t)
	createRandomBusinessLineForCompany(t, company1.ID)

	// Buat 1 business line untuk company 2
	_, _ = createRandomBusinessLine(t)

	// Test: Ambil business line untuk company 1
	lines, err := businessLineTestRepo.ListByCompanyID(context.Background(), company1.ID)
	if err != nil {
		t.Fatalf("Gagal list business lines by company ID: %v", err)
	}

	if len(lines) != 2 {
		t.Errorf("Diharapkan 2 business lines, tetapi mendapatkan %d", len(lines))
	}
}

func TestBusinessLineRepository_Update(t *testing.T) {
	defer cleanup()

	initialBL, _ := createRandomBusinessLine(t)

	updatedName := "Updated Business Line Name"
	initialBL.Name = updatedName
	initialBL.UpdatedAt = time.Now()

	err := businessLineTestRepo.Update(context.Background(), initialBL)
	if err != nil {
		t.Fatalf("Gagal mengupdate business line: %v", err)
	}

	foundBL, err := businessLineTestRepo.GetByID(context.Background(), initialBL.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan business line setelah update: %v", err)
	}
	if foundBL.Name != updatedName {
		t.Errorf("Name tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedName, foundBL.Name)
	}
}

func TestBusinessLineRepository_Delete(t *testing.T) {
	defer cleanup()

	blToDelete, _ := createRandomBusinessLine(t)

	err := businessLineTestRepo.Delete(context.Background(), blToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus business line: %v", err)
	}

	_, err = businessLineTestRepo.GetByID(context.Background(), blToDelete.ID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}

// createRandomBusinessLineForCompany adalah helper spesifik untuk test list
func createRandomBusinessLineForCompany(t *testing.T, companyID uuid.UUID) {
	bl := &models.BusinessLine{ID: uuid.New(), CompanyID: companyID, Name: "Another BL", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := businessLineTestRepo.Create(context.Background(), bl)
	if err != nil {
		t.Fatalf("Gagal membuat business line untuk company spesifik: %v", err)
	}
}
