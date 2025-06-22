package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// DiscountRepository mendefinisikan interface untuk operasi data terkait Discount.
type DiscountRepository interface {
	Create(ctx context.Context, discount *models.Discount) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Discount, error)
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Discount, error)
	Update(ctx context.Context, discount *models.Discount) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// pgDiscountRepository adalah implementasi dari DiscountRepository untuk PostgreSQL.
type pgDiscountRepository struct {
	db DBExecutor
}

// NewPgDiscountRepository adalah constructor untuk membuat instance baru dari pgDiscountRepository.
func NewPgDiscountRepository(db DBExecutor) DiscountRepository {
	return &pgDiscountRepository{db: db}
}

// Create menyisipkan aturan diskon baru.
func (r *pgDiscountRepository) Create(ctx context.Context, d *models.Discount) error {
	query := `
		INSERT INTO discounts (id, company_id, name, description, discount_type, discount_value, applicable_to,
			master_product_id_applicable, store_product_id_applicable, category_applicable, customer_tier_applicable,
			min_purchase_amount, start_date, end_date, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`
	_, err := r.db.ExecContext(ctx, query,
		d.ID, d.CompanyID, d.Name, d.Description, d.DiscountType, d.DiscountValue, d.ApplicableTo,
		d.MasterProductIDApplicable, d.StoreProductIDApplicable, d.CategoryApplicable, d.CustomerTierApplicable,
		d.MinPurchaseAmount, d.StartDate, d.EndDate, d.IsActive, d.CreatedAt, d.UpdatedAt,
	)
	return err
}

// GetByID mengambil diskon berdasarkan ID.
func (r *pgDiscountRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Discount, error) {
	d := &models.Discount{}
	query := ` 
		SELECT id, company_id, name, description, discount_type, discount_value, applicable_to,
			master_product_id_applicable, store_product_id_applicable, category_applicable, customer_tier_applicable,
			min_purchase_amount, start_date, end_date, is_active, created_at, updated_at
		FROM discounts
		WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&d.ID, &d.CompanyID, &d.Name, &d.Description, &d.DiscountType, &d.DiscountValue, &d.ApplicableTo,
		&d.MasterProductIDApplicable, &d.StoreProductIDApplicable, &d.CategoryApplicable, &d.CustomerTierApplicable,
		&d.MinPurchaseAmount, &d.StartDate, &d.EndDate, &d.IsActive, &d.CreatedAt, &d.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return d, nil
}

// ListByCompanyID mengambil daftar diskon untuk sebuah perusahaan.
func (r *pgDiscountRepository) ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Discount, error) {
	query := `
		SELECT id, company_id, name, discount_type, discount_value, start_date, end_date, is_active
		FROM discounts
		WHERE company_id = $1
		ORDER BY start_date DESC`
	rows, err := r.db.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var discounts []*models.Discount
	for rows.Next() {
		d := &models.Discount{}
		if err := rows.Scan(
			&d.ID, &d.CompanyID, &d.Name, &d.DiscountType, &d.DiscountValue,
			&d.StartDate, &d.EndDate, &d.IsActive,
		); err != nil {
			return nil, err
		}
		discounts = append(discounts, d)
	}
	return discounts, rows.Err()
}

// Update memperbarui aturan diskon.
func (r *pgDiscountRepository) Update(ctx context.Context, d *models.Discount) error {
	query := `
		UPDATE discounts 
		SET name = $1, description = $2, discount_type = $3, discount_value = $4, applicable_to = $5,
			master_product_id_applicable = $6, store_product_id_applicable = $7, category_applicable = $8,
			customer_tier_applicable = $9, min_purchase_amount = $10, start_date = $11, end_date = $12,
			is_active = $13, updated_at = $14
		WHERE id = $15 AND company_id = $16`
	_, err := r.db.ExecContext(ctx, query,
		d.Name, d.Description, d.DiscountType, d.DiscountValue, d.ApplicableTo,
		d.MasterProductIDApplicable, d.StoreProductIDApplicable, d.CategoryApplicable,
		d.CustomerTierApplicable, d.MinPurchaseAmount, d.StartDate, d.EndDate,
		d.IsActive, d.UpdatedAt, d.ID, d.CompanyID,
	)
	return err
}

// Delete menghapus aturan diskon.
func (r *pgDiscountRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM discounts WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
