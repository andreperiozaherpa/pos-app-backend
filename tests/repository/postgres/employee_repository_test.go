package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
)

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
