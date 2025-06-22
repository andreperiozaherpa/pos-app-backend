package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomOperationalExpense is a helper to create a random operational expense.
func createRandomOperationalExpense(t *testing.T) *models.OperationalExpense {
	// Create dependencies: businessLine and its associated company.
	// createRandomBusinessLine already creates a company internally.
	businessLine, company := createRandomBusinessLine(t)

	// Create a store under this business line.
	store := createRandomStore(t, businessLine.ID)

	// Create an employee (user).
	user := createRandomEmployee(t)

	expense := &models.OperationalExpense{
		ID:              uuid.New(),
		CompanyID:       company.ID,
		StoreID:         uuid.NullUUID{UUID: store.ID, Valid: true}, // Associate with a store
		ExpenseDate:     time.Now().AddDate(0, 0, -5),
		Category:        "Gaji Karyawan",
		Description:     sql.NullString{String: "Gaji bulanan karyawan", Valid: true},
		Amount:          1500000.00,
		CreatedByUserID: uuid.NullUUID{UUID: user.UserID, Valid: true},
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err := operationalExpenseTestRepo.Create(context.Background(), expense)
	if err != nil {
		t.Fatalf("Gagal membuat operational expense random untuk test: %v", err)
	}

	return expense
}

func TestOperationalExpenseRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	newExpense := createRandomOperationalExpense(t)

	foundExpense, err := operationalExpenseTestRepo.GetByID(context.Background(), newExpense.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan operational expense by ID: %v", err)
	}
	if foundExpense == nil {
		t.Fatal("Operational expense yang baru dibuat tidak ditemukan")
	}

	if newExpense.ID != foundExpense.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newExpense.ID, foundExpense.ID)
	}
	if newExpense.CompanyID != foundExpense.CompanyID {
		t.Errorf("CompanyID tidak cocok. Diharapkan '%s', didapatkan '%s'", newExpense.CompanyID, foundExpense.CompanyID)
	}
	if newExpense.Category != foundExpense.Category {
		t.Errorf("Category tidak cocok. Diharapkan '%s', didapatkan '%s'", newExpense.Category, foundExpense.Category)
	}
	if newExpense.Amount != foundExpense.Amount {
		t.Errorf("Amount tidak cocok. Diharapkan %f, didapatkan %f", newExpense.Amount, foundExpense.Amount)
	}
	// Normalisasi kedua tanggal ke awal hari di UTC untuk perbandingan yang akurat.
	// Ini penting karena kolom database kemungkinan adalah tipe DATE.
	normalizedNewExpenseDate := time.Date(newExpense.ExpenseDate.Year(), newExpense.ExpenseDate.Month(), newExpense.ExpenseDate.Day(), 0, 0, 0, 0, time.UTC)
	normalizedFoundExpenseDate := time.Date(foundExpense.ExpenseDate.Year(), foundExpense.ExpenseDate.Month(), foundExpense.ExpenseDate.Day(), 0, 0, 0, 0, time.UTC)
	if !normalizedNewExpenseDate.Equal(normalizedFoundExpenseDate) {
		t.Errorf("ExpenseDate tidak cocok. Diharapkan '%s', didapatkan '%s'", newExpense.ExpenseDate.Format("2006-01-02"), foundExpense.ExpenseDate.Format("2006-01-02"))
	}
}

func TestOperationalExpenseRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()
	nonExistentID := uuid.New()
	foundExpense, err := operationalExpenseTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundExpense != nil {
		t.Errorf("Diharapkan operational expense nil, tetapi mendapatkan: %+v", foundExpense)
	}
}

func TestOperationalExpenseRepository_ListByCompanyID(t *testing.T) {
	defer cleanup()

	expense1 := createRandomOperationalExpense(t)
	companyID := expense1.CompanyID

	// Create another expense for the same company
	expense2 := &models.OperationalExpense{
		ID:              uuid.New(),
		CompanyID:       companyID,
		StoreID:         expense1.StoreID, // Same store
		ExpenseDate:     time.Now().AddDate(0, 0, -10),
		Category:        "Sewa Toko",
		Amount:          5000000.00,
		CreatedByUserID: expense1.CreatedByUserID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	err := operationalExpenseTestRepo.Create(context.Background(), expense2)
	if err != nil {
		t.Fatalf("Gagal membuat expense kedua: %v", err)
	}

	// Create an expense for a different company
	createRandomOperationalExpense(t)

	foundExpenses, err := operationalExpenseTestRepo.ListByCompanyID(context.Background(), companyID)
	if err != nil {
		t.Fatalf("Gagal list operational expenses by company ID: %v", err)
	}

	if len(foundExpenses) != 2 {
		t.Errorf("Diharapkan 2 operational expenses untuk company, tetapi mendapatkan %d", len(foundExpenses))
	}
	// Verifikasi urutan (terbaru pertama berdasarkan expense_date)
	normalizedExpense1Date := time.Date(expense1.ExpenseDate.Year(), expense1.ExpenseDate.Month(), expense1.ExpenseDate.Day(), 0, 0, 0, 0, time.UTC)
	normalizedFound0Date := time.Date(foundExpenses[0].ExpenseDate.Year(), foundExpenses[0].ExpenseDate.Month(), foundExpenses[0].ExpenseDate.Day(), 0, 0, 0, 0, time.UTC)
	if !normalizedFound0Date.Equal(normalizedExpense1Date) {
		t.Errorf("Expected most recent expense to be first. Expected '%s', got '%s'", expense1.ExpenseDate.Format("2006-01-02"), foundExpenses[0].ExpenseDate.Format("2006-01-02"))
	}
}

func TestOperationalExpenseRepository_ListByStoreID(t *testing.T) {
	defer cleanup()

	expense1 := createRandomOperationalExpense(t)
	storeID := expense1.StoreID.UUID

	// Create another expense for the same store
	expense2 := &models.OperationalExpense{
		ID:              uuid.New(),
		CompanyID:       expense1.CompanyID,
		StoreID:         expense1.StoreID, // Same store
		ExpenseDate:     time.Now().AddDate(0, 0, -15),
		Category:        "Listrik & Air",
		Amount:          750000.00,
		CreatedByUserID: expense1.CreatedByUserID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	err := operationalExpenseTestRepo.Create(context.Background(), expense2)
	if err != nil {
		t.Fatalf("Gagal membuat expense kedua untuk store: %v", err)
	}

	// Create an expense for a different store
	createRandomOperationalExpense(t)

	foundExpenses, err := operationalExpenseTestRepo.ListByStoreID(context.Background(), storeID)
	if err != nil {
		t.Fatalf("Gagal list operational expenses by store ID: %v", err)
	}

	if len(foundExpenses) != 2 {
		t.Errorf("Diharapkan 2 operational expenses untuk store, tetapi mendapatkan %d", len(foundExpenses))
	}
	// Verifikasi urutan (terbaru pertama berdasarkan expense_date)
	normalizedExpense1Date := time.Date(expense1.ExpenseDate.Year(), expense1.ExpenseDate.Month(), expense1.ExpenseDate.Day(), 0, 0, 0, 0, time.UTC)
	normalizedFound0Date := time.Date(foundExpenses[0].ExpenseDate.Year(), foundExpenses[0].ExpenseDate.Month(), foundExpenses[0].ExpenseDate.Day(), 0, 0, 0, 0, time.UTC)
	if !normalizedFound0Date.Equal(normalizedExpense1Date) {
		t.Errorf("Expected most recent expense to be first. Expected '%s', got '%s'", expense1.ExpenseDate.Format("2006-01-02"), foundExpenses[0].ExpenseDate.Format("2006-01-02"))
	}
}

func TestOperationalExpenseRepository_Update(t *testing.T) {
	defer cleanup()

	initialExpense := createRandomOperationalExpense(t)

	updatedAmount := 2000000.00
	updatedDescription := sql.NullString{String: "Gaji bulanan karyawan yang direvisi", Valid: true}
	initialExpense.Amount = updatedAmount
	initialExpense.Description = updatedDescription
	initialExpense.UpdatedAt = time.Now()

	err := operationalExpenseTestRepo.Update(context.Background(), initialExpense)
	if err != nil {
		t.Fatalf("Gagal mengupdate operational expense: %v", err)
	}

	foundExpense, err := operationalExpenseTestRepo.GetByID(context.Background(), initialExpense.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan operational expense setelah update: %v", err)
	}
	if foundExpense.Amount != updatedAmount {
		t.Errorf("Amount tidak terupdate. Diharapkan %f, didapatkan %f", updatedAmount, foundExpense.Amount)
	}
	if foundExpense.Description.String != updatedDescription.String {
		t.Errorf("Description tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedDescription.String, foundExpense.Description.String)
	}
}

func TestOperationalExpenseRepository_Delete(t *testing.T) {
	defer cleanup()

	expenseToDelete := createRandomOperationalExpense(t)

	err := operationalExpenseTestRepo.Delete(context.Background(), expenseToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus operational expense: %v", err)
	}

	_, err = operationalExpenseTestRepo.GetByID(context.Background(), expenseToDelete.ID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}
