package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomShift adalah helper untuk membuat shift baru.
func createRandomShift(t *testing.T) (*models.Shift, *models.Employee) {
	employee := createRandomEmployee(t) // Helper ini sudah membuat user, company, store

	// Pastikan employee memiliki storeID yang valid
	if !employee.StoreID.Valid {
		t.Fatalf("Employee yang dibuat tidak memiliki StoreID yang valid")
	}

	shiftDate := time.Now().AddDate(0, 0, 1) // Besok
	startTime := time.Date(shiftDate.Year(), shiftDate.Month(), shiftDate.Day(), 9, 0, 0, 0, time.Local)
	endTime := time.Date(shiftDate.Year(), shiftDate.Month(), shiftDate.Day(), 17, 0, 0, 0, time.Local)

	shift := &models.Shift{
		ID:              uuid.New(),
		EmployeeUserID:  employee.UserID,
		StoreID:         employee.StoreID.UUID,
		ShiftDate:       startTime, // DATE field, gunakan bagian tanggal dari time.Time
		StartTime:       startTime,
		EndTime:         endTime,
		ActualCheckIn:   sql.NullTime{},
		ActualCheckOut:  sql.NullTime{},
		Notes:           sql.NullString{String: "Morning shift", Valid: true},
		CreatedByUserID: uuid.NullUUID{UUID: employee.UserID, Valid: true}, // Asumsi karyawan yang membuat shift
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err := shiftTestRepo.Create(context.Background(), shift)
	if err != nil {
		t.Fatalf("Gagal membuat shift random untuk test: %v", err)
	}

	return shift, employee
}

func TestShiftRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	newShift, _ := createRandomShift(t)

	foundShift, err := shiftTestRepo.GetByID(context.Background(), newShift.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan shift by ID: %v", err)
	}
	if foundShift == nil {
		t.Fatal("Shift yang baru dibuat tidak ditemukan")
	}

	if newShift.ID != foundShift.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%s', didapatkan '%s'", newShift.ID, foundShift.ID)
	}
	if newShift.EmployeeUserID != foundShift.EmployeeUserID {
		t.Errorf("EmployeeUserID tidak cocok. Diharapkan '%s', didapatkan '%s'", newShift.EmployeeUserID, foundShift.EmployeeUserID)
	}
	if newShift.StoreID != foundShift.StoreID {
		t.Errorf("StoreID tidak cocok. Diharapkan '%s', didapatkan '%s'", newShift.StoreID, foundShift.StoreID)
	}
	// Perbandingan waktu perlu hati-hati karena tipe DATE di PostgreSQL hanya menyimpan tanggal.
	// Kita hanya perlu membandingkan bagian tanggalnya saja, mengabaikan waktu dan zona waktu.
	expectedShiftDate := newShift.ShiftDate.Truncate(24 * time.Hour)
	actualShiftDate := foundShift.ShiftDate.Truncate(24 * time.Hour)
	if !expectedShiftDate.Equal(actualShiftDate) {
		t.Errorf("ShiftDate tidak cocok. Diharapkan '%s', didapatkan '%s'", expectedShiftDate, actualShiftDate)
	}
}

func TestShiftRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()
	nonExistentID := uuid.New()
	foundShift, err := shiftTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundShift != nil {
		t.Errorf("Diharapkan shift nil, tetapi mendapatkan: %+v", foundShift)
	}
}

func TestShiftRepository_ListByEmployeeID(t *testing.T) {
	defer cleanup()

	shift1, employee := createRandomShift(t)
	// Buat shift lain untuk karyawan yang sama
	shift2Date := shift1.ShiftDate.AddDate(0, 0, 2)
	shift2 := &models.Shift{
		ID:             uuid.New(),
		EmployeeUserID: employee.UserID,
		StoreID:        employee.StoreID.UUID,
		ShiftDate:      shift2Date,
		StartTime:      time.Date(shift2Date.Year(), shift2Date.Month(), shift2Date.Day(), 8, 0, 0, 0, time.Local),
		EndTime:        time.Date(shift2Date.Year(), shift2Date.Month(), shift2Date.Day(), 16, 0, 0, 0, time.Local),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	err := shiftTestRepo.Create(context.Background(), shift2)
	if err != nil {
		t.Fatalf("Gagal membuat shift kedua untuk employee: %v", err)
	}

	// Buat shift untuk karyawan lain
	createRandomShift(t)

	shifts, err := shiftTestRepo.ListByEmployeeID(context.Background(), employee.UserID)
	if err != nil {
		t.Fatalf("Gagal list shifts by employee ID: %v", err)
	}

	if len(shifts) != 2 {
		t.Errorf("Diharapkan 2 shifts, tetapi mendapatkan %d", len(shifts))
	}
}

func TestShiftRepository_Update(t *testing.T) {
	defer cleanup()

	initialShift, _ := createRandomShift(t)

	updatedNotes := "Updated notes for shift"
	actualCheckIn := time.Now().Add(1 * time.Hour)
	initialShift.Notes = sql.NullString{String: updatedNotes, Valid: true}
	initialShift.ActualCheckIn = sql.NullTime{Time: actualCheckIn, Valid: true}
	initialShift.UpdatedAt = time.Now()

	err := shiftTestRepo.Update(context.Background(), initialShift)
	if err != nil {
		t.Fatalf("Gagal mengupdate shift: %v", err)
	}

	foundShift, err := shiftTestRepo.GetByID(context.Background(), initialShift.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan shift setelah update: %v", err)
	}
	if foundShift.Notes.String != updatedNotes {
		t.Errorf("Notes tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedNotes, foundShift.Notes.String)
	}
	if !foundShift.ActualCheckIn.Time.Truncate(time.Second).Equal(actualCheckIn.Truncate(time.Second)) {
		t.Errorf("ActualCheckIn tidak terupdate. Diharapkan '%s', didapatkan '%s'", actualCheckIn, foundShift.ActualCheckIn.Time)
	}
}

func TestShiftRepository_Delete(t *testing.T) {
	defer cleanup()

	shiftToDelete, _ := createRandomShift(t)

	err := shiftTestRepo.Delete(context.Background(), shiftToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus shift: %v", err)
	}

	_, err = shiftTestRepo.GetByID(context.Background(), shiftToDelete.ID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}
