package services

import (
	"context"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

// ActivityLogService menangani pencatatan dan pengelolaan log aktivitas user/sistem.
type ActivityLogService interface {
	// CRUD utama
	CreateActivityLog(ctx context.Context, log *models.ActivityLog) (int64, error)
	ListActivityLogsByUserID(ctx context.Context, userID uuid.UUID) ([]*models.ActivityLog, error)
	ListActivityLogsByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.ActivityLog, error)
	ListActivityLogsByStoreID(ctx context.Context, storeID uuid.UUID) ([]*models.ActivityLog, error)

	// Opsional & Custom
	SearchActivityLogs(ctx context.Context, query string, from, to time.Time) ([]*models.ActivityLog, error)
	ExportActivityLogs(ctx context.Context, companyID uuid.UUID, from, to time.Time) ([]byte, error)
	DeleteOldActivityLogs(ctx context.Context, olderThan time.Time) error
	GetActivityLogDetail(ctx context.Context, logID int64) (*models.ActivityLog, error)
	ListActivityLogsByDateRange(ctx context.Context, companyID uuid.UUID, from, to time.Time) ([]*models.ActivityLog, error)
}
