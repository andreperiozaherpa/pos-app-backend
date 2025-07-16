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

type stockReportRepositoryPG struct {
	db *sql.DB
}

func NewStockReportRepositoryPG(db *sql.DB) repository.StockReportRepository {
	return &stockReportRepositoryPG{db: db}
}

func (r *stockReportRepositoryPG) Create(ctx context.Context, report *models.StockReport) error {
	query := `
        INSERT INTO stock_reports (id, store_id, report_date, total_stock_value, created_at)
        VALUES ($1, $2, $3, $4, NOW())`
	_, err := r.db.ExecContext(ctx, query,
		report.ID, report.StoreID, report.ReportDate, report.TotalStockValue)
	return err
}

func (r *stockReportRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.StockReport, error) {
	query := `
        SELECT id, store_id, report_date, total_stock_value, created_at
        FROM stock_reports WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	report := &models.StockReport{}
	err := row.Scan(&report.ID, &report.StoreID, &report.ReportDate, &report.TotalStockValue, &report.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return report, nil
}

func (r *stockReportRepositoryPG) ListByStore(ctx context.Context, storeID uuid.UUID, fromDate, toDate time.Time) ([]*models.StockReport, error) {
	query := `
        SELECT id, store_id, report_date, total_stock_value, created_at
        FROM stock_reports
        WHERE store_id = $1 AND report_date BETWEEN $2 AND $3
        ORDER BY report_date DESC`
	rows, err := r.db.QueryContext(ctx, query, storeID, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []*models.StockReport
	for rows.Next() {
		r := &models.StockReport{}
		if err := rows.Scan(&r.ID, &r.StoreID, &r.ReportDate, &r.TotalStockValue, &r.CreatedAt); err != nil {
			return nil, err
		}
		reports = append(reports, r)
	}
	return reports, nil
}
