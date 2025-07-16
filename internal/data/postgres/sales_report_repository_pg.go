package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"
)

type salesReportRepositoryPG struct {
	db *sql.DB
}

func NewSalesReportRepositoryPG(db *sql.DB) repository.SalesReportRepository {
	return &salesReportRepositoryPG{db: db}
}

func (r *salesReportRepositoryPG) Create(ctx context.Context, report *models.SalesReport) error {
	query := `
        INSERT INTO sales_reports (id, store_id, report_date, total_sales, created_at)
        VALUES ($1, $2, $3, $4, NOW())`
	_, err := r.db.ExecContext(ctx, query,
		report.ID, report.StoreID, report.ReportDate, report.TotalSales)
	return err
}

func (r *salesReportRepositoryPG) GetByID(ctx context.Context, id string) (*models.SalesReport, error) {
	query := `
        SELECT id, store_id, report_date, total_sales, created_at
        FROM sales_reports WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	report := &models.SalesReport{}
	err := row.Scan(&report.ID, &report.StoreID, &report.ReportDate, &report.TotalSales, &report.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return report, nil
}

func (r *salesReportRepositoryPG) ListByStore(ctx context.Context, storeID string, fromDate, toDate time.Time) ([]*models.SalesReport, error) {
	query := `
        SELECT id, store_id, report_date, total_sales, created_at
        FROM sales_reports
        WHERE store_id = $1 AND report_date BETWEEN $2 AND $3
        ORDER BY report_date DESC`
	rows, err := r.db.QueryContext(ctx, query, storeID, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []*models.SalesReport
	for rows.Next() {
		r := &models.SalesReport{}
		if err := rows.Scan(&r.ID, &r.StoreID, &r.ReportDate, &r.TotalSales, &r.CreatedAt); err != nil {
			return nil, err
		}
		reports = append(reports, r)
	}
	return reports, nil
}
