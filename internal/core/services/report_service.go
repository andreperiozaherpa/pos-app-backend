package services

import (
	"context"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

// ReportService menangani pembuatan dan ekspor laporan agregat (opsional, custom).
type ReportService interface {
	GenerateSalesReport(ctx context.Context, companyID uuid.UUID, from, to time.Time) (*models.SalesReport, error)
	GenerateStockReport(ctx context.Context, storeID uuid.UUID, asOf time.Time) (*models.StockReport, error)
	GenerateProfitLossReport(ctx context.Context, companyID uuid.UUID, from, to time.Time) (*models.ProfitLossReport, error)
	GenerateEmployeePerformanceReport(ctx context.Context, companyID uuid.UUID, from, to time.Time) (*models.EmployeePerformanceReport, error)
	GenerateCustomerActivityReport(ctx context.Context, companyID uuid.UUID, from, to time.Time) (*models.CustomerActivityReport, error)
	ExportReportToPDF(ctx context.Context, reportType string, params map[string]interface{}) ([]byte, error)
	ScheduleReportGeneration(ctx context.Context, reportType string, scheduleTime time.Time, params map[string]interface{}) error
	GenerateCustomReport(ctx context.Context, queryParams map[string]interface{}) ([]byte, error)
}
