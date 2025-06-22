package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// CompanyRepository mendefinisikan interface untuk operasi data terkait Company.
type CompanyRepository interface {
	Create(ctx context.Context, company *models.Company) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Company, error)
	Update(ctx context.Context, company *models.Company) error
	Delete(ctx context.Context, id uuid.UUID) error
	// Kita mungkin perlu metode lain seperti ListAll, dll.
}

// pgCompanyRepository adalah implementasi dari CompanyRepository untuk PostgreSQL.
type pgCompanyRepository struct {
	db DBExecutor
}

// NewPgCompanyRepository adalah constructor untuk membuat instance baru dari pgCompanyRepository.
func NewPgCompanyRepository(db DBExecutor) CompanyRepository {
	return &pgCompanyRepository{db: db}
}

// Implementasi metode-metode dari interface CompanyRepository:

func (r *pgCompanyRepository) Create(ctx context.Context, company *models.Company) error {
	query := `
		INSERT INTO companies (id, name, address, contact_info, tax_id_number, default_tax_percentage, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.ExecContext(ctx, query,
		company.ID,
		company.Name,
		company.Address,
		company.ContactInfo,
		company.TaxIDNumber,
		company.DefaultTaxPercentage,
		company.CreatedAt,
		company.UpdatedAt,
	)
	return err
}

func (r *pgCompanyRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Company, error) {
	company := &models.Company{}
	query := `
		SELECT id, name, address, contact_info, tax_id_number, default_tax_percentage, created_at, updated_at
		FROM companies
		WHERE id = $1`

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&company.ID,
		&company.Name,
		&company.Address,
		&company.ContactInfo,
		&company.TaxIDNumber,
		&company.DefaultTaxPercentage,
		&company.CreatedAt,
		&company.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return company, nil
}

func (r *pgCompanyRepository) Update(ctx context.Context, company *models.Company) error {
	query := `
		UPDATE companies
		SET name = $1, address = $2, contact_info = $3, tax_id_number = $4, default_tax_percentage = $5, updated_at = $6
		WHERE id = $7`

	_, err := r.db.ExecContext(ctx, query, company.Name, company.Address, company.ContactInfo, company.TaxIDNumber, company.DefaultTaxPercentage, company.UpdatedAt, company.ID)
	return err
}

func (r *pgCompanyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	// Implementasi saat ini adalah hard delete.
	// Untuk soft delete, Anda perlu menambahkan kolom 'is_active' atau 'deleted_at' ke tabel 'companies'.
	query := `DELETE FROM companies WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
