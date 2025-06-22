package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// TaxRateRepository mendefinisikan interface untuk operasi data terkait TaxRate.
type TaxRateRepository interface {
	Create(ctx context.Context, taxRate *models.TaxRate) error
	GetByID(ctx context.Context, id int32) (*models.TaxRate, error)
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.TaxRate, error)
	Update(ctx context.Context, taxRate *models.TaxRate) error
	Delete(ctx context.Context, id int32) error
}

// pgTaxRateRepository adalah implementasi dari TaxRateRepository untuk PostgreSQL.
type pgTaxRateRepository struct {
	db DBExecutor
}

// NewPgTaxRateRepository adalah constructor untuk membuat instance baru dari pgTaxRateRepository.
func NewPgTaxRateRepository(db DBExecutor) TaxRateRepository {
	return &pgTaxRateRepository{db: db}
}

// Create menyisipkan tarif pajak baru dan mengembalikan ID yang di-generate.
func (r *pgTaxRateRepository) Create(ctx context.Context, tr *models.TaxRate) error {
	query := `
		INSERT INTO tax_rates (company_id, name, rate_percentage, description, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
	return r.db.QueryRowContext(ctx, query,
		tr.CompanyID, tr.Name, tr.RatePercentage, tr.Description, tr.IsActive, tr.CreatedAt, tr.UpdatedAt,
	).Scan(&tr.ID)
}

// GetByID mengambil tarif pajak berdasarkan ID.
func (r *pgTaxRateRepository) GetByID(ctx context.Context, id int32) (*models.TaxRate, error) {
	tr := &models.TaxRate{}
	query := `
		SELECT id, company_id, name, rate_percentage, description, is_active, created_at, updated_at
		FROM tax_rates
		WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&tr.ID, &tr.CompanyID, &tr.Name, &tr.RatePercentage, &tr.Description, &tr.IsActive, &tr.CreatedAt, &tr.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return tr, nil
}

// ListByCompanyID mengambil daftar tarif pajak untuk perusahaan tertentu.
func (r *pgTaxRateRepository) ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.TaxRate, error) {
	query := `
		SELECT id, company_id, name, rate_percentage, description, is_active, created_at, updated_at
		FROM tax_rates
		WHERE company_id = $1
		ORDER BY name ASC`
	rows, err := r.db.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var taxRates []*models.TaxRate
	for rows.Next() {
		tr := &models.TaxRate{}
		if err := rows.Scan(
			&tr.ID, &tr.CompanyID, &tr.Name, &tr.RatePercentage, &tr.Description, &tr.IsActive, &tr.CreatedAt, &tr.UpdatedAt,
		); err != nil {
			return nil, err
		}
		taxRates = append(taxRates, tr)
	}
	return taxRates, rows.Err()
}

// Update memperbarui data tarif pajak.
func (r *pgTaxRateRepository) Update(ctx context.Context, tr *models.TaxRate) error {
	query := `
		UPDATE tax_rates
		SET name = $1, rate_percentage = $2, description = $3, is_active = $4, updated_at = $5
		WHERE id = $6 AND company_id = $7`
	_, err := r.db.ExecContext(ctx, query,
		tr.Name, tr.RatePercentage, tr.Description, tr.IsActive, tr.UpdatedAt, tr.ID, tr.CompanyID,
	)
	return err
}

// Delete menghapus tarif pajak dari database.
func (r *pgTaxRateRepository) Delete(ctx context.Context, id int32) error {
	query := `DELETE FROM tax_rates WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
