package repository

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// ActivityLogRepository mendefinisikan interface untuk operasi data terkait ActivityLog.
type ActivityLogRepository interface {
	// Create membuat data activity log baru.
	Create(ctx context.Context, log *models.ActivityLog) error

	// ListByUserID mengambil daftar activity log berdasarkan user ID.
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]*models.ActivityLog, error)

	// ListByCompanyID mengambil daftar activity log berdasarkan company ID.
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.ActivityLog, error)

	// ListByStoreID mengambil daftar activity log berdasarkan store ID.
	ListByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.ActivityLog, error)
}
