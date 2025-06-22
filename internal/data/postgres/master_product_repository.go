package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// MasterProductRepository mendefinisikan interface untuk operasi data terkait MasterProduct.
type MasterProductRepository interface {
	Create(ctx context.Context, mp *models.MasterProduct) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.MasterProduct, error)
	GetByCompanyAndCode(ctx context.Context, companyID uuid.UUID, code string) (*models.MasterProduct, error)
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.MasterProduct, error)
	Update(ctx context.Context, mp *models.MasterProduct) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// pgMasterProductRepository adalah implementasi dari MasterProductRepository untuk PostgreSQL.
type pgMasterProductRepository struct {
	db *sql.DB
}

// NewPgMasterProductRepository adalah constructor untuk membuat instance baru dari pgMasterProductRepository.
func NewPgMasterProductRepository(db *sql.DB) MasterProductRepository {
	return &pgMasterProductRepository{db: db}
}

// Implementasi metode-metode dari interface MasterProductRepository:

func (r *pgMasterProductRepository) Create(ctx context.Context, mp *models.MasterProduct) error {
	query := `
		INSERT INTO master_products (id, company_id, master_product_code, name, description, category,
			unit_of_measure, barcode, default_tax_rate_id, image_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := r.db.ExecContext(ctx, query,
		mp.ID, mp.CompanyID, mp.MasterProductCode, mp.Name, mp.Description, mp.Category,
		mp.UnitOfMeasure, mp.Barcode, mp.DefaultTaxRateID, mp.ImageURL, mp.CreatedAt, mp.UpdatedAt,
	)
	return err
}

func (r *pgMasterProductRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.MasterProduct, error) {
	mp := &models.MasterProduct{}
	query := `
		SELECT id, company_id, master_product_code, name, description, category,
			unit_of_measure, barcode, default_tax_rate_id, image_url, created_at, updated_at
		FROM master_products
		WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&mp.ID, &mp.CompanyID, &mp.MasterProductCode, &mp.Name, &mp.Description, &mp.Category,
		&mp.UnitOfMeasure, &mp.Barcode, &mp.DefaultTaxRateID, &mp.ImageURL, &mp.CreatedAt, &mp.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return mp, nil
}

func (r *pgMasterProductRepository) GetByCompanyAndCode(ctx context.Context, companyID uuid.UUID, code string) (*models.MasterProduct, error) {
	mp := &models.MasterProduct{}
	query := `
		SELECT id, company_id, master_product_code, name, description, category,
			unit_of_measure, barcode, default_tax_rate_id, image_url, created_at, updated_at
		FROM master_products
		WHERE company_id = $1 AND master_product_code = $2`
	err := r.db.QueryRowContext(ctx, query, companyID, code).Scan(
		&mp.ID, &mp.CompanyID, &mp.MasterProductCode, &mp.Name, &mp.Description, &mp.Category,
		&mp.UnitOfMeasure, &mp.Barcode, &mp.DefaultTaxRateID, &mp.ImageURL, &mp.CreatedAt, &mp.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return mp, nil
}

func (r *pgMasterProductRepository) ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.MasterProduct, error) {
	query := `
		SELECT id, company_id, master_product_code, name, description, category,
			unit_of_measure, barcode, default_tax_rate_id, image_url, created_at, updated_at
		FROM master_products
		WHERE company_id = $1
		ORDER BY name ASC`
	rows, err := r.db.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var masterProducts []*models.MasterProduct
	for rows.Next() {
		mp := &models.MasterProduct{}
		if err := rows.Scan(
			&mp.ID, &mp.CompanyID, &mp.MasterProductCode, &mp.Name, &mp.Description, &mp.Category,
			&mp.UnitOfMeasure, &mp.Barcode, &mp.DefaultTaxRateID, &mp.ImageURL, &mp.CreatedAt, &mp.UpdatedAt,
		); err != nil {
			return nil, err
		}
		masterProducts = append(masterProducts, mp)
	}
	return masterProducts, rows.Err()
}

func (r *pgMasterProductRepository) Update(ctx context.Context, mp *models.MasterProduct) error {
	query := `
		UPDATE master_products
		SET company_id = $1, master_product_code = $2, name = $3, description = $4, category = $5,
			unit_of_measure = $6, barcode = $7, default_tax_rate_id = $8, image_url = $9, updated_at = $10
		WHERE id = $11`
	_, err := r.db.ExecContext(ctx, query,
		mp.CompanyID, mp.MasterProductCode, mp.Name, mp.Description, mp.Category,
		mp.UnitOfMeasure, mp.Barcode, mp.DefaultTaxRateID, mp.ImageURL, mp.UpdatedAt, mp.ID,
	)
	return err
}

func (r *pgMasterProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM master_products WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
