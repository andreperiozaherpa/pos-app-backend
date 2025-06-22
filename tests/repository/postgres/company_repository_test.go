package repository_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomCompany adalah fungsi helper untuk membuat dan menyimpan company baru ke DB.
func createRandomCompany(t *testing.T) *models.Company {
	contactInfoJSON, err := json.Marshal(map[string]string{"email": "contact@example.com", "phone": "123-456-7890"})
	if err != nil {
		t.Fatalf("Gagal marshal contact info untuk test: %v", err)
	}

	company := &models.Company{
		ID:                   uuid.New(),
		Name:                 "Test Company " + uuid.NewString(),
		Address:              sql.NullString{String: "123 Test St, Testville", Valid: true},
		ContactInfo:          contactInfoJSON,
		TaxIDNumber:          sql.NullString{String: "TAX" + uuid.NewString(), Valid: true},
		DefaultTaxPercentage: sql.NullFloat64{Float64: 11.50, Valid: true},
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	err = companyTestRepo.Create(context.Background(), company)
	if err != nil {
		t.Fatalf("Gagal membuat company random untuk test: %v", err)
	}
	return company
}

func TestCompanyRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	// 1. Buat company baru menggunakan helper
	newCompany := createRandomCompany(t)

	// 2. Ambil company berdasarkan ID
	foundCompany, err := companyTestRepo.GetByID(context.Background(), newCompany.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan company by ID: %v", err)
	}
	if foundCompany == nil {
		t.Fatal("Company yang baru dibuat tidak ditemukan")
	}

	// 3. Verifikasi data
	if newCompany.ID != foundCompany.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newCompany.ID, foundCompany.ID)
	}
	if newCompany.Name != foundCompany.Name {
		t.Errorf("Name tidak cocok. Diharapkan '%s', didapatkan '%s'", newCompany.Name, foundCompany.Name)
	}

	// Cara yang benar untuk membandingkan JSON: unmarshal ke map dan bandingkan map-nya.
	var expectedContact, actualContact map[string]interface{}
	if err := json.Unmarshal(newCompany.ContactInfo, &expectedContact); err != nil {
		t.Fatalf("Gagal unmarshal expected contact info: %v", err)
	}
	if err := json.Unmarshal(foundCompany.ContactInfo, &actualContact); err != nil {
		t.Fatalf("Gagal unmarshal actual contact info: %v", err)
	}

	// Gunakan reflect.DeepEqual untuk perbandingan struktur yang andal.
	if !reflect.DeepEqual(expectedContact, actualContact) {
		t.Errorf("ContactInfo tidak cocok.\nDiharapkan: %v\nDidapatkan: %v", expectedContact, actualContact)
	}
}

func TestCompanyRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()

	nonExistentID := uuid.New()
	foundCompany, err := companyTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundCompany != nil {
		t.Errorf("Diharapkan company nil, tetapi mendapatkan: %+v", foundCompany)
	}
}

func TestCompanyRepository_Update(t *testing.T) {
	defer cleanup()

	// 1. Buat company awal
	initialCompany := createRandomCompany(t)

	// 2. Modifikasi data
	updatedName := "Updated Company Name"
	initialCompany.Name = updatedName
	initialCompany.UpdatedAt = time.Now()

	// 3. Panggil Update
	err := companyTestRepo.Update(context.Background(), initialCompany)
	if err != nil {
		t.Fatalf("Gagal mengupdate company: %v", err)
	}

	// 4. Ambil kembali dan verifikasi
	foundCompany, err := companyTestRepo.GetByID(context.Background(), initialCompany.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan company setelah update: %v", err)
	}
	if foundCompany.Name != updatedName {
		t.Errorf("Name tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedName, foundCompany.Name)
	}
}

func TestCompanyRepository_Delete(t *testing.T) {
	defer cleanup()

	// 1. Buat company untuk dihapus
	companyToDelete := createRandomCompany(t)

	// 2. Panggil Delete
	err := companyTestRepo.Delete(context.Background(), companyToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus company: %v", err)
	}

	// 3. Verifikasi bahwa company sudah tidak ada
	_, err = companyTestRepo.GetByID(context.Background(), companyToDelete.ID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}
