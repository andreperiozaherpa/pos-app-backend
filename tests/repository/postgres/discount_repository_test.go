package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

func TestDiscountRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	newDiscount, _ := createRandomDiscount(t)

	foundDiscount, err := discountTestRepo.GetByID(context.Background(), newDiscount.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan diskon by ID: %v", err)
	}
	if foundDiscount == nil {
		t.Fatal("Diskon yang baru dibuat tidak ditemukan")
	}

	if newDiscount.ID != foundDiscount.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newDiscount.ID, foundDiscount.ID)
	}
	if newDiscount.Name != foundDiscount.Name {
		t.Errorf("Name tidak cocok. Diharapkan '%s', didapatkan '%s'", newDiscount.Name, foundDiscount.Name)
	}
	if newDiscount.DiscountValue != foundDiscount.DiscountValue {
		t.Errorf("DiscountValue tidak cocok. Diharapkan %f, didapatkan %f", newDiscount.DiscountValue, foundDiscount.DiscountValue)
	}
}

func TestDiscountRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()
	nonExistentID := uuid.New()
	foundDiscount, err := discountTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundDiscount != nil {
		t.Errorf("Diharapkan diskon nil, tetapi mendapatkan: %+v", foundDiscount)
	}
}

func TestDiscountRepository_ListByCompanyID(t *testing.T) {
	defer cleanup()

	// Buat 2 diskon untuk company 1
	_, company1 := createRandomDiscount(t)
	discount2 := &models.Discount{ID: uuid.New(), CompanyID: company1.ID, Name: "Discount 2", DiscountType: models.DiscountTypeFixedAmount, DiscountValue: 5000, ApplicableTo: models.DiscountApplicableToTotalTransaction, StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 7), IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := discountTestRepo.Create(context.Background(), discount2)
	if err != nil {
		t.Fatalf("Gagal membuat diskon kedua: %v", err)
	}

	// Buat 1 diskon untuk company 2
	createRandomDiscount(t)

	discounts, err := discountTestRepo.ListByCompanyID(context.Background(), company1.ID)
	if err != nil {
		t.Fatalf("Gagal list diskon by company ID: %v", err)
	}

	if len(discounts) != 2 {
		t.Errorf("Diharapkan 2 diskon, tetapi mendapatkan %d", len(discounts))
	}
}

func TestDiscountRepository_Update(t *testing.T) {
	defer cleanup()

	initialDiscount, _ := createRandomDiscount(t)

	updatedName := "Diskon Kemerdekaan"
	initialDiscount.Name = updatedName
	initialDiscount.IsActive = false
	initialDiscount.UpdatedAt = time.Now()

	err := discountTestRepo.Update(context.Background(), initialDiscount)
	if err != nil {
		t.Fatalf("Gagal mengupdate diskon: %v", err)
	}

	foundDiscount, err := discountTestRepo.GetByID(context.Background(), initialDiscount.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan diskon setelah update: %v", err)
	}
	if foundDiscount.Name != updatedName {
		t.Errorf("Name tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedName, foundDiscount.Name)
	}
	if foundDiscount.IsActive != false {
		t.Errorf("IsActive tidak terupdate. Diharapkan false, didapatkan %t", foundDiscount.IsActive)
	}
}
