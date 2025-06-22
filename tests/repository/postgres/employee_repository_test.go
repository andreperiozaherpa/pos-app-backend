package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createEmployeeDependencies adalah helper untuk membuat semua data yang dibutuhkan oleh seorang Employee.
func createEmployeeDependencies(t *testing.T) (userID, companyID, storeID uuid.UUID) {
	// 1. Buat User
	user := &models.User{
		ID:        uuid.New(),
		UserType:  models.UserTypeEmployee,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := userTestRepo.Create(context.Background(), user)
	if err != nil {
		t.Fatalf("Gagal membuat user dependency untuk test employee: %v", err)
	}

	// 2. Buat Company dan BusinessLine
	businessLineID := createTestCompanyAndBusinessLine(t)

	// 3. Ambil Company ID dari BusinessLine (ini asumsi, lebih baik jika helper mengembalikan companyID juga)
	var compID uuid.UUID
	err = testDB.QueryRow(`SELECT company_id FROM business_lines WHERE id = $1`, businessLineID).Scan(&compID)
	if err != nil {
		t.Fatalf("Gagal mengambil company_id dari business_line: %v", err)
	}

	// 4. Buat Store
	store := createRandomStore(t, businessLineID)

	return user.ID, compID, store.ID
}

// createRandomEmployee adalah fungsi helper untuk membuat dan menyimpan employee baru ke DB.
func createRandomEmployee(t *testing.T) *models.Employee {
	userID, companyID, storeID := createEmployeeDependencies(t)

	employee := &models.Employee{
		UserID:           userID,
		CompanyID:        companyID,
		StoreID:          uuid.NullUUID{UUID: storeID, Valid: true},
		EmployeeIDNumber: sql.NullString{String: "EMP-" + uuid.NewString()[:8], Valid: true},
		JoinDate:         sql.NullTime{Time: time.Now(), Valid: true},
		Position:         sql.NullString{String: "Kasir", Valid: true},
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err := employeeTestRepo.Create(context.Background(), employee)
	if err != nil {
		t.Fatalf("Gagal membuat employee random untuk test: %v", err)
	}
	return employee
}

func TestEmployeeRepository_CreateAndGetByUserID(t *testing.T) {
	defer cleanup()

	newEmployee := createRandomEmployee(t)

	foundEmployee, err := employeeTestRepo.GetByUserID(context.Background(), newEmployee.UserID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan employee by UserID: %v", err)
	}
	if foundEmployee == nil {
		t.Fatal("Employee yang baru dibuat tidak ditemukan")
	}

	if newEmployee.UserID != foundEmployee.UserID {
		t.Errorf("UserID tidak cocok. Diharapkan '%s', didapatkan '%s'", newEmployee.UserID, foundEmployee.UserID)
	}
	if newEmployee.CompanyID != foundEmployee.CompanyID {
		t.Errorf("CompanyID tidak cocok. Diharapkan '%s', didapatkan '%s'", newEmployee.CompanyID, foundEmployee.CompanyID)
	}
	if newEmployee.Position.String != foundEmployee.Position.String {
		t.Errorf("Position tidak cocok. Diharapkan '%s', didapatkan '%s'", newEmployee.Position.String, foundEmployee.Position.String)
	}
}

func TestEmployeeRepository_GetByUserID_NotFound(t *testing.T) {
	defer cleanup()

	nonExistentID := uuid.New()
	foundEmployee, err := employeeTestRepo.GetByUserID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundEmployee != nil {
		t.Errorf("Diharapkan employee nil, tetapi mendapatkan: %+v", foundEmployee)
	}
}

func TestEmployeeRepository_Update(t *testing.T) {
	defer cleanup()

	initialEmployee := createRandomEmployee(t)

	updatedPosition := "Manajer Toko"
	initialEmployee.Position = sql.NullString{String: updatedPosition, Valid: true}
	initialEmployee.UpdatedAt = time.Now()

	err := employeeTestRepo.Update(context.Background(), initialEmployee)
	if err != nil {
		t.Fatalf("Gagal mengupdate employee: %v", err)
	}

	foundEmployee, err := employeeTestRepo.GetByUserID(context.Background(), initialEmployee.UserID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan employee setelah update: %v", err)
	}

	if foundEmployee.Position.String != updatedPosition {
		t.Errorf("Position tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedPosition, foundEmployee.Position.String)
	}
}

func TestEmployeeRepository_Delete(t *testing.T) {
	defer cleanup()

	employeeToDelete := createRandomEmployee(t)

	err := employeeTestRepo.Delete(context.Background(), employeeToDelete.UserID)
	if err != nil {
		t.Fatalf("Gagal menghapus employee: %v", err)
	}

	_, err = employeeTestRepo.GetByUserID(context.Background(), employeeToDelete.UserID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}
