package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomCustomer adalah fungsi helper untuk membuat dan menyimpan customer baru ke DB.
func createRandomCustomer(t *testing.T) *models.Customer {
	// 1. Buat User dengan tipe Customer
	user := &models.User{
		ID:        uuid.New(),
		UserType:  models.UserTypeCustomer,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := userTestRepo.Create(context.Background(), user)
	if err != nil {
		t.Fatalf("Gagal membuat user dependency untuk test customer: %v", err)
	}

	// 2. Buat Company
	company := createRandomCompany(t)

	// 3. Buat Customer
	customer := &models.Customer{
		UserID:           user.ID,
		CompanyID:        company.ID,
		MembershipNumber: sql.NullString{String: "MEM-" + uuid.NewString()[:8], Valid: true},
		JoinDate:         sql.NullTime{Time: time.Now(), Valid: true},
		Points:           100, // Menggunakan int32 langsung karena kolom points tidak NULL
		Tier:             sql.NullString{String: "Gold", Valid: true},
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err = customerTestRepo.Create(context.Background(), customer)
	if err != nil {
		t.Fatalf("Gagal membuat customer random untuk test: %v", err)
	}
	return customer
}

func TestCustomerRepository_CreateAndGetByUserID(t *testing.T) {
	defer cleanup()

	newCustomer := createRandomCustomer(t)

	foundCustomer, err := customerTestRepo.GetByUserID(context.Background(), newCustomer.UserID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan customer by UserID: %v", err)
	}
	if foundCustomer == nil {
		t.Fatal("Customer yang baru dibuat tidak ditemukan")
	}

	if newCustomer.UserID != foundCustomer.UserID {
		t.Errorf("UserID tidak cocok. Diharapkan '%s', didapatkan '%s'", newCustomer.UserID, foundCustomer.UserID)
	}
	if newCustomer.MembershipNumber.String != foundCustomer.MembershipNumber.String {
		t.Errorf("MembershipNumber tidak cocok. Diharapkan '%s', didapatkan '%s'", newCustomer.MembershipNumber.String, foundCustomer.MembershipNumber.String)
	}
}

func TestCustomerRepository_GetByUserID_NotFound(t *testing.T) {
	defer cleanup()

	nonExistentID := uuid.New()
	foundCustomer, err := customerTestRepo.GetByUserID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundCustomer != nil {
		t.Errorf("Diharapkan customer nil, tetapi mendapatkan: %+v", foundCustomer)
	}
}

func TestCustomerRepository_Update(t *testing.T) {
	defer cleanup()

	initialCustomer := createRandomCustomer(t)

	updatedPoints := int32(500)
	updatedTier := "Platinum"
	initialCustomer.Points = updatedPoints // Menggunakan int32 langsung
	initialCustomer.Tier = sql.NullString{String: updatedTier, Valid: true}
	initialCustomer.UpdatedAt = time.Now()

	err := customerTestRepo.Update(context.Background(), initialCustomer)
	if err != nil {
		t.Fatalf("Gagal mengupdate customer: %v", err)
	}

	foundCustomer, err := customerTestRepo.GetByUserID(context.Background(), initialCustomer.UserID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan customer setelah update: %v", err)
	}

	if foundCustomer.Points != updatedPoints {
		t.Errorf("Points tidak terupdate. Diharapkan '%d', didapatkan '%d'", updatedPoints, foundCustomer.Points)
	}
	if foundCustomer.Tier.String != updatedTier {
		t.Errorf("Tier tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedTier, foundCustomer.Tier.String)
	}
}

func TestCustomerRepository_Delete(t *testing.T) {
	defer cleanup()

	customerToDelete := createRandomCustomer(t)

	err := customerTestRepo.Delete(context.Background(), customerToDelete.UserID)
	if err != nil {
		t.Fatalf("Gagal menghapus customer: %v", err)
	}

	_, err = customerTestRepo.GetByUserID(context.Background(), customerToDelete.UserID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}
