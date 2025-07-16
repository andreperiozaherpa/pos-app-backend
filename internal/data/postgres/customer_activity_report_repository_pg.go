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

type customerActivityReportRepositoryPG struct {
	db *sql.DB
}

func NewCustomerActivityReportRepositoryPG(db *sql.DB) repository.CustomerActivityReportRepository {
	return &customerActivityReportRepositoryPG{db: db}
}

// Create menambahkan laporan aktivitas customer.
func (r *customerActivityReportRepositoryPG) Create(ctx context.Context, report *models.CustomerActivityReport) error {
	query := `
        INSERT INTO customer_activity_reports (
			id, customer_user_id, company_id, activity_date, total_transactions, 
			total_amount, points_earned, last_transaction_at, created_at
		)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())`
	_, err := r.db.ExecContext(ctx, query,
		report.ID, report.CustomerUserID, report.CompanyID, report.ActivityDate,
		report.TotalTransactions, report.TotalAmount, report.PointsEarned, report.LastTransactionAt,
	)
	return err
}

// GetByID mengambil laporan berdasarkan ID.
func (r *customerActivityReportRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.CustomerActivityReport, error) {
	query := `
        SELECT id, customer_user_id, company_id, activity_date, total_transactions, 
			   total_amount, points_earned, last_transaction_at, created_at
        FROM customer_activity_reports WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	report := &models.CustomerActivityReport{}
	err := row.Scan(
		&report.ID, &report.CustomerUserID, &report.CompanyID, &report.ActivityDate,
		&report.TotalTransactions, &report.TotalAmount, &report.PointsEarned,
		&report.LastTransactionAt, &report.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return report, nil
}

// ListByCustomer mengambil laporan aktivitas customer per periode.
func (r *customerActivityReportRepositoryPG) ListByCustomer(ctx context.Context, customerUserID uuid.UUID, fromDate, toDate time.Time) ([]*models.CustomerActivityReport, error) {
	query := `
        SELECT id, customer_user_id, company_id, activity_date, total_transactions, 
			   total_amount, points_earned, last_transaction_at, created_at
        FROM customer_activity_reports
        WHERE customer_user_id = $1 AND activity_date BETWEEN $2 AND $3
        ORDER BY activity_date DESC`
	rows, err := r.db.QueryContext(ctx, query, customerUserID, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []*models.CustomerActivityReport
	for rows.Next() {
		r := &models.CustomerActivityReport{}
		if err := rows.Scan(
			&r.ID, &r.CustomerUserID, &r.CompanyID, &r.ActivityDate,
			&r.TotalTransactions, &r.TotalAmount, &r.PointsEarned,
			&r.LastTransactionAt, &r.CreatedAt,
		); err != nil {
			return nil, err
		}
		reports = append(reports, r)
	}
	return reports, nil
}

// ListByCompany mengambil seluruh laporan aktivitas customer untuk suatu company.
func (r *customerActivityReportRepositoryPG) ListByCompany(ctx context.Context, companyID uuid.UUID, fromDate, toDate time.Time) ([]*models.CustomerActivityReport, error) {
	query := `
        SELECT id, customer_user_id, company_id, activity_date, total_transactions, 
			   total_amount, points_earned, last_transaction_at, created_at
        FROM customer_activity_reports
        WHERE company_id = $1
        ORDER BY activity_date DESC`
	rows, err := r.db.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []*models.CustomerActivityReport
	for rows.Next() {
		r := &models.CustomerActivityReport{}
		if err := rows.Scan(
			&r.ID, &r.CustomerUserID, &r.CompanyID, &r.ActivityDate,
			&r.TotalTransactions, &r.TotalAmount, &r.PointsEarned,
			&r.LastTransactionAt, &r.CreatedAt,
		); err != nil {
			return nil, err
		}
		reports = append(reports, r)
	}
	return reports, nil
}
