package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type employeePerformanceReportRepositoryPG struct {
	db *sql.DB
}

func NewEmployeePerformanceReportRepositoryPG(db *sql.DB) repository.EmployeePerformanceReportRepository {
	return &employeePerformanceReportRepositoryPG{db: db}
}

// Create menambah data laporan kinerja karyawan ke database.
func (r *employeePerformanceReportRepositoryPG) Create(ctx context.Context, report *models.EmployeePerformanceReport) error {
	query := `
        INSERT INTO employee_performance_reports 
        (id, employee_user_id, report_date, performance_score, created_at)
        VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query,
		report.ID, report.EmployeeUserID, report.ReportDate, report.PerformanceScore, report.CreatedAt)
	return err
}

// GetByID mengambil data laporan kinerja berdasarkan ID.
func (r *employeePerformanceReportRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.EmployeePerformanceReport, error) {
	query := `
        SELECT id, employee_user_id, report_date, performance_score, created_at
        FROM employee_performance_reports WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	report := &models.EmployeePerformanceReport{}
	err := row.Scan(&report.ID, &report.EmployeeUserID, &report.ReportDate, &report.PerformanceScore, &report.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return report, nil
}

// ListByEmployee mengambil daftar laporan kinerja berdasarkan employee_user_id dan rentang tanggal.
func (r *employeePerformanceReportRepositoryPG) ListByEmployee(ctx context.Context, employeeUserID uuid.UUID, fromDate, toDate time.Time) ([]*models.EmployeePerformanceReport, error) {
	query := `
        SELECT id, employee_user_id, report_date, performance_score, created_at
        FROM employee_performance_reports
        WHERE employee_user_id = $1 AND report_date BETWEEN $2 AND $3
        ORDER BY report_date DESC`
	rows, err := r.db.QueryContext(ctx, query, employeeUserID, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []*models.EmployeePerformanceReport
	for rows.Next() {
		r := &models.EmployeePerformanceReport{}
		if err := rows.Scan(&r.ID, &r.EmployeeUserID, &r.ReportDate, &r.PerformanceScore, &r.CreatedAt); err != nil {
			return nil, err
		}
		reports = append(reports, r)
	}
	return reports, nil
}
