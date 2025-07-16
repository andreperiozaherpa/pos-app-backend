package repository

import (
	"context"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type EmployeePerformanceReportRepository interface {
	Create(ctx context.Context, report *models.EmployeePerformanceReport) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.EmployeePerformanceReport, error)
	ListByEmployee(ctx context.Context, employeeUserID uuid.UUID, fromDate, toDate time.Time) ([]*models.EmployeePerformanceReport, error)
}
