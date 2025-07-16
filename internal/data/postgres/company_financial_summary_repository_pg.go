package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type companyFinancialSummaryRepositoryPG struct {
	db *sql.DB
}

func NewCompanyFinancialSummaryRepositoryPG(db *sql.DB) repository.CompanyFinancialSummaryRepository {
	return &companyFinancialSummaryRepositoryPG{db: db}
}

// Create menambahkan data ringkasan keuangan baru ke dalam database.
func (r *companyFinancialSummaryRepositoryPG) Create(ctx context.Context, summary *models.CompanyFinancialSummary) error {
	query := `
        INSERT INTO company_financial_summaries (company_id, total_revenue, total_expense, net_profit, last_updated)
        VALUES ($1, $2, $3, $4, NOW())`
	_, err := r.db.ExecContext(ctx, query,
		summary.CompanyID, summary.TotalRevenue, summary.TotalExpense, summary.NetProfit)
	return err
}

// GetByCompanyID mengambil ringkasan keuangan berdasarkan ID perusahaan.
func (r *companyFinancialSummaryRepositoryPG) GetByCompanyID(ctx context.Context, companyID uuid.UUID) (*models.CompanyFinancialSummary, error) {
	query := `
        SELECT company_id, total_revenue, total_expense, net_profit, last_updated
        FROM company_financial_summaries
        WHERE company_id=$1`
	row := r.db.QueryRowContext(ctx, query, companyID)
	summary := &models.CompanyFinancialSummary{}
	err := row.Scan(
		&summary.CompanyID, &summary.TotalRevenue,
		&summary.TotalExpense, &summary.NetProfit, &summary.LastUpdated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return summary, nil
}

// Update memperbarui data ringkasan keuangan yang sudah ada.
func (r *companyFinancialSummaryRepositoryPG) Update(ctx context.Context, summary *models.CompanyFinancialSummary) error {
	query := `
        UPDATE company_financial_summaries SET total_revenue=$1, total_expense=$2, net_profit=$3, last_updated=NOW()
        WHERE company_id=$4`
	res, err := r.db.ExecContext(ctx, query,
		summary.TotalRevenue, summary.TotalExpense, summary.NetProfit, summary.CompanyID)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// Delete menghapus data ringkasan keuangan berdasarkan ID.
func (r *companyFinancialSummaryRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM company_financial_summaries WHERE company_id = $1`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return repository.ErrNotFound
	}
	return nil
}
