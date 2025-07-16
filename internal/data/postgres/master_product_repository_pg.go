package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type masterProductRepositoryPG struct {
	db *sql.DB
}

func NewMasterProductRepositoryPG(db *sql.DB) repository.MasterProductRepository {
	return &masterProductRepositoryPG{db: db}
}

// Create menambahkan produk pusat ke database.
func (r *masterProductRepositoryPG) Create(ctx context.Context, p *models.MasterProduct) error {
	query := `
        INSERT INTO master_products 
        (id, company_id, master_product_code, name, description, category, unit_of_measure, barcode, default_tax_rate_id, image_url, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW())`
	_, err := r.db.ExecContext(ctx, query,
		p.ID,
		p.CompanyID,
		p.MasterProductCode,
		p.Name,
		p.Description,
		p.Category,
		p.UnitOfMeasure,
		p.Barcode,
		p.DefaultTaxRateID,
		p.ImageURL,
	)
	return err
}

func (r *masterProductRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.MasterProduct, error) {
	query := `
        SELECT id, company_id, master_product_code, name, description, category, unit_of_measure, barcode, default_tax_rate_id, image_url, created_at, updated_at
        FROM master_products WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	p := &models.MasterProduct{}
	err := row.Scan(
		&p.ID,
		&p.CompanyID,
		&p.MasterProductCode,
		&p.Name,
		&p.Description,
		&p.Category,
		&p.UnitOfMeasure,
		&p.Barcode,
		&p.DefaultTaxRateID,
		&p.ImageURL,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return p, nil
}

// GetByCompanyAndCode mengambil produk pusat berdasarkan company_id dan master_product_code.
func (r *masterProductRepositoryPG) GetByCompanyAndCode(ctx context.Context, companyID uuid.UUID, code string) (*models.MasterProduct, error) {
	query := `
        SELECT id, company_id, master_product_code, name, description, category, unit_of_measure, barcode, default_tax_rate_id, image_url, created_at, updated_at
        FROM master_products
        WHERE company_id = $1 AND master_product_code = $2`
	row := r.db.QueryRowContext(ctx, query, companyID, code)

	p := &models.MasterProduct{}
	err := row.Scan(
		&p.ID,
		&p.CompanyID,
		&p.MasterProductCode,
		&p.Name,
		&p.Description,
		&p.Category,
		&p.UnitOfMeasure,
		&p.Barcode,
		&p.DefaultTaxRateID,
		&p.ImageURL,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return p, nil
}

func (r *masterProductRepositoryPG) Update(ctx context.Context, p *models.MasterProduct) error {
	query := `
        UPDATE master_products SET 
            company_id = $1,
            master_product_code = $2,
            name = $3,
            description = $4,
            category = $5,
            unit_of_measure = $6,
            barcode = $7,
            default_tax_rate_id = $8,
            image_url = $9,
            updated_at = NOW()
        WHERE id = $10`
	result, err := r.db.ExecContext(ctx, query,
		p.CompanyID,
		p.MasterProductCode,
		p.Name,
		p.Description,
		p.Category,
		p.UnitOfMeasure,
		p.Barcode,
		p.DefaultTaxRateID,
		p.ImageURL,
		p.ID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

func (r *masterProductRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM master_products WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// UpdateStock menyesuaikan stok produk pusat
func (r *masterProductRepositoryPG) UpdateStock(ctx context.Context, productID uuid.UUID, quantity int) error {
	query := `UPDATE master_products SET stock = $1, updated_at = NOW() WHERE id = $2`
	result, err := r.db.ExecContext(ctx, query, quantity, productID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// ListVariants mengambil daftar varian produk dari master product
func (r *masterProductRepositoryPG) ListVariants(ctx context.Context, masterProductID uuid.UUID) ([]*models.MasterProductVariant, error) {
	query := `
        SELECT id, master_product_id, variant_name, created_at, updated_at 
        FROM master_product_variants WHERE master_product_id = $1 ORDER BY variant_name`
	rows, err := r.db.QueryContext(ctx, query, masterProductID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var variants []*models.MasterProductVariant
	for rows.Next() {
		v := &models.MasterProductVariant{}
		if err := rows.Scan(&v.ID, &v.MasterProductID, &v.Name, &v.CreatedAt, &v.UpdatedAt); err != nil {
			return nil, err
		}
		variants = append(variants, v)
	}
	return variants, nil
}

// SearchProducts mencari produk pusat berdasarkan filter nama/kategori dsb.
func (r *masterProductRepositoryPG) SearchProducts(ctx context.Context, companyID uuid.UUID, filter string) ([]*models.MasterProduct, error) {
	query := `
        SELECT id, company_id, master_product_code, name, description, category, unit_of_measure, barcode, default_tax_rate_id, image_url, created_at, updated_at
        FROM master_products
        WHERE company_id = $1 AND (name ILIKE '%' || $2 || '%' OR category ILIKE '%' || $2 || '%')
        ORDER BY name`
	rows, err := r.db.QueryContext(ctx, query, companyID, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.MasterProduct
	for rows.Next() {
		p := &models.MasterProduct{}
		if err := rows.Scan(
			&p.ID,
			&p.CompanyID,
			&p.MasterProductCode,
			&p.Name,
			&p.Description,
			&p.Category,
			&p.UnitOfMeasure,
			&p.Barcode,
			&p.DefaultTaxRateID,
			&p.ImageURL,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// ListByCompanyID mengambil semua master product untuk sebuah company.
func (r *masterProductRepositoryPG) ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.MasterProduct, error) {
	query := `
		SELECT id, company_id, master_product_code, name, description, category, unit_of_measure, barcode, default_tax_rate_id, image_url, created_at, updated_at
		FROM master_products
		WHERE company_id = $1
		ORDER BY name`
	rows, err := r.db.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.MasterProduct
	for rows.Next() {
		p := &models.MasterProduct{}
		if err := rows.Scan(
			&p.ID,
			&p.CompanyID,
			&p.MasterProductCode,
			&p.Name,
			&p.Description,
			&p.Category,
			&p.UnitOfMeasure,
			&p.Barcode,
			&p.DefaultTaxRateID,
			&p.ImageURL,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
