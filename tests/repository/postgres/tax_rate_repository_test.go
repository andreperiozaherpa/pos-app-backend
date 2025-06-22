package repository_test

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomTaxRate adalah helper untuk membuat tax rate baru.
// Helper ini sekarang menggunakan TaxRateRepository.Create.
func createRandomTaxRate(t *testing.T, companyID uuid.UUID) *models.TaxRate {
	tr := &models.TaxRate{
		CompanyID:      companyID,
		Name:           "Tax " + randomString(5),
		RatePercentage: float64(rand.Intn(2000)) / 100, // Random percentage up to 20.00%
		Description:    sql.NullString{String: "Random tax rate", Valid: true},
		IsActive:       true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := taxRateTestRepo.Create(context.Background(), tr)
	if err != nil {
		t.Fatalf("Gagal membuat tax rate random untuk test: %v", err)
	}
	if tr.ID == 0 {
		t.Fatal("TaxRate ID tidak di-populate setelah create")
	}
	return tr
}

func TestTaxRateRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	company := createRandomCompany(t)
	newTaxRate := createRandomTaxRate(t, company.ID)

	foundTaxRate, err := taxRateTestRepo.GetByID(context.Background(), newTaxRate.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan tax rate by ID: %v", err)
	}
	if foundTaxRate == nil {
		t.Fatal("Tax rate yang baru dibuat tidak ditemukan")
	}

	if newTaxRate.ID != foundTaxRate.ID {
		t.Errorf("ID tidak cocok. Diharapkan %d, didapatkan %d", newTaxRate.ID, foundTaxRate.ID)
	}
	if newTaxRate.Name != foundTaxRate.Name {
		t.Errorf("Name tidak cocok. Diharapkan '%s', didapatkan '%s'", newTaxRate.Name, foundTaxRate.Name)
	}
	if newTaxRate.CompanyID != foundTaxRate.CompanyID {
		t.Errorf("CompanyID tidak cocok. Diharapkan '%s', didapatkan '%s'", newTaxRate.CompanyID, foundTaxRate.CompanyID)
	}
}

func TestTaxRateRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()
	nonExistentID := int32(999999)
	foundTaxRate, err := taxRateTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundTaxRate != nil {
		t.Errorf("Diharapkan tax rate nil, tetapi mendapatkan: %+v", foundTaxRate)
	}
}

func TestTaxRateRepository_ListByCompanyID(t *testing.T) {
	defer cleanup()

	company1 := createRandomCompany(t)
	createRandomTaxRate(t, company1.ID)
	createRandomTaxRate(t, company1.ID)

	company2 := createRandomCompany(t)
	createRandomTaxRate(t, company2.ID)

	taxRates, err := taxRateTestRepo.ListByCompanyID(context.Background(), company1.ID)
	if err != nil {
		t.Fatalf("Gagal list tax rates by company ID: %v", err)
	}

	if len(taxRates) != 2 {
		t.Errorf("Diharapkan 2 tax rates untuk company1, tetapi mendapatkan %d", len(taxRates))
	}

	// Verifikasi bahwa semua tax rate yang ditemukan memang milik company1
	for _, tr := range taxRates {
		if tr.CompanyID != company1.ID {
			t.Errorf("Tax rate dengan ID %d seharusnya milik company %s, tetapi milik %s", tr.ID, company1.ID, tr.CompanyID)
		}
	}
}

func TestTaxRateRepository_Update(t *testing.T) {
	defer cleanup()

	company := createRandomCompany(t)
	initialTaxRate := createRandomTaxRate(t, company.ID)

	updatedName := "Updated Tax Name"
	updatedRate := 15.50
	initialTaxRate.Name = updatedName
	initialTaxRate.RatePercentage = updatedRate
	initialTaxRate.IsActive = false
	initialTaxRate.UpdatedAt = time.Now()

	err := taxRateTestRepo.Update(context.Background(), initialTaxRate)
	if err != nil {
		t.Fatalf("Gagal mengupdate tax rate: %v", err)
	}

	foundTaxRate, err := taxRateTestRepo.GetByID(context.Background(), initialTaxRate.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan tax rate setelah update: %v", err)
	}
	if foundTaxRate.Name != updatedName {
		t.Errorf("Name tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedName, foundTaxRate.Name)
	}
	if foundTaxRate.RatePercentage != updatedRate {
		t.Errorf("RatePercentage tidak terupdate. Diharapkan %f, didapatkan %f", updatedRate, foundTaxRate.RatePercentage)
	}
	if foundTaxRate.IsActive != false {
		t.Errorf("IsActive tidak terupdate. Diharapkan false, didapatkan %t", foundTaxRate.IsActive)
	}
}

func TestTaxRateRepository_Delete(t *testing.T) {
	defer cleanup()

	company := createRandomCompany(t)
	taxRateToDelete := createRandomTaxRate(t, company.ID)

	err := taxRateTestRepo.Delete(context.Background(), taxRateToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus tax rate: %v", err)
	}

	_, err = taxRateTestRepo.GetByID(context.Background(), taxRateToDelete.ID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}
