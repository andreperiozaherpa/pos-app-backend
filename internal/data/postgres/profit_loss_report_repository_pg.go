package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"
)

type profitLossReportRepositoryPG struct {
	db *sql.DB
}

func NewProfitLossReportRepositoryPG(db *sql.DB) repository.ProfitLossReportRepository {
	return &profitLossReportRepositoryPG{db: db}
}

func (r *profitLossReportRepositoryPG) Create(ctx context.Context, report *models.ProfitLossReport) error {
	query := `
        INSERT INTO profit_loss_reports (id, company_id, report_date, total_revenue, total_expense, net_profit, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, NOW())`
	_, err := r.db.ExecContext(ctx, query,
		report.ID, report.CompanyID, report.ReportDate, report.TotalRevenue, report.TotalExpense, report.NetProfit)
	return err
}

func (r *profitLossReportRepositoryPG) GetByID(ctx context.Context, id string) (*models.ProfitLossReport, error) {
	query := `
        SELECT id, company_id, report_date, total_revenue, total_expense, net_profit, created_at
        FROM profit_loss_reports WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	report := &models.ProfitLossReport{}
	err := row.Scan(&report.ID, &report.CompanyID, &report.ReportDate, &report.TotalRevenue, &report.TotalExpense, &report.NetProfit, &report.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return report, nil
}

func (r *profitLossReportRepositoryPG) ListByCompany(ctx context.Context, companyID string, fromDate, toDate time.Time) ([]*models.ProfitLossReport, error) {
	query := `
        SELECT id, company_id, report_date, total_revenue, total_expense, net_profit, created_at
        FROM profit_loss_reports
        WHERE company_id = $1 AND report_date BETWEEN $2 AND $3
        ORDER BY report_date DESC`
	rows, err := r.db.QueryContext(ctx, query, companyID, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []*models.ProfitLossReport
	for rows.Next() {
		r := &models.ProfitLossReport{}
		if err := rows.Scan(&r.ID, &r.CompanyID, &r.ReportDate, &r.TotalRevenue, &r.TotalExpense, &r.NetProfit, &r.CreatedAt); err != nil {
			return nil, err
		}
		reports = append(reports, r)
	}
	return reports, nil
}
