package repository

import (
	"context"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

// ShiftRepository mendefinisikan interface untuk operasi data terkait Shift.
type ShiftRepository interface {
	// Create membuat data shift baru.
	Create(ctx context.Context, shift *models.Shift) error

	// GetByID mengambil data shift berdasarkan ID.
	GetByID(ctx context.Context, id uuid.UUID) (*models.Shift, error)

	// ListByEmployeeID mengambil daftar shift berdasarkan user ID karyawan.
	ListByEmployeeID(ctx context.Context, employeeUserID uuid.UUID) ([]*models.Shift, error)

	// ListByStoreIDAndDateRange mengambil daftar shift berdasarkan store ID dan rentang tanggal.
	ListByStoreIDAndDateRange(ctx context.Context, storeID uuid.UUID, startDate, endDate time.Time) ([]*models.Shift, error)

	// Update memperbarui data shift.
	Update(ctx context.Context, shift *models.Shift) error

	// Delete menghapus data shift berdasarkan ID.
	Delete(ctx context.Context, id uuid.UUID) error
}
