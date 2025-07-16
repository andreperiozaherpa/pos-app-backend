package services

import (
	"context"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

// ImportExportService menangani proses import/export data aplikasi (bulk data, migrasi, dsb).
type ImportExportService interface {
	ImportData(ctx context.Context, data []byte, dataType string, uploaderID uuid.UUID) (*models.ImportResult, error)
	ExportData(ctx context.Context, dataType string, filter map[string]interface{}) ([]byte, error)
	ValidateImportData(ctx context.Context, data []byte, dataType string) (*models.ImportValidationResult, error)
	GenerateImportTemplate(ctx context.Context, dataType string) ([]byte, error)
	ListImportHistory(ctx context.Context, companyID uuid.UUID) ([]*models.ImportHistory, error)
	ListExportHistory(ctx context.Context, companyID uuid.UUID) ([]*models.ExportHistory, error)
	CancelImportExportTask(ctx context.Context, taskID uuid.UUID) error
	ScheduleImport(ctx context.Context, data []byte, dataType string, scheduleTime time.Time, uploaderID uuid.UUID) error
	ScheduleExport(ctx context.Context, dataType string, filter map[string]interface{}, scheduleTime time.Time, requesterID uuid.UUID) error
}
