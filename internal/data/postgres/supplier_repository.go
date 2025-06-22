package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// SupplierRepository mendefinisikan interface untuk operasi data terkait Supplier.
type SupplierRepository interface {
	Create(ctx context.Context, supplier *models.Supplier) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Supplier, error)
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Supplier, error)
	Update(ctx context.Context, supplier *models.Supplier) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// pgSupplierRepository adalah implementasi dari SupplierRepository untuk PostgreSQL.
type pgSupplierRepository struct {
	db *sql.DB
}

// NewPgSupplierRepository adalah constructor untuk membuat instance baru dari pgSupplierRepository.
func NewPgSupplierRepository(db *sql.DB) SupplierRepository {
	return &pgSupplierRepository{db: db}
}

// Implementasi metode-metode dari interface SupplierRepository:

func (r *pgSupplierRepository) Create(ctx context.Context, s *models.Supplier) error {
	query := `
		INSERT INTO suppliers (id, company_id, name, contact_person, email, phone_number, address, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.db.ExecContext(ctx, query, s.ID, s.CompanyID, s.Name, s.ContactPerson, s.Email, s.PhoneNumber, s.Address, s.CreatedAt, s.UpdatedAt)
	return err
}

func (r *pgSupplierRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Supplier, error) {
	s := &models.Supplier{}
	query := `
		SELECT id, company_id, name, contact_person, email, phone_number, address, created_at, updated_at
		FROM suppliers
		WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&s.ID, &s.CompanyID, &s.Name, &s.ContactPerson, &s.Email, &s.PhoneNumber, &s.Address, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return s, nil
}

func (r *pgSupplierRepository) ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Supplier, error) {
	query := `
		SELECT id, company_id, name, contact_person, email, phone_number, address, created_at, updated_at
		FROM suppliers
		WHERE company_id = $1
		ORDER BY name ASC`
	rows, err := r.db.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []*models.Supplier
	for rows.Next() {
		s := &models.Supplier{}
		if err := rows.Scan(&s.ID, &s.CompanyID, &s.Name, &s.ContactPerson, &s.Email, &s.PhoneNumber, &s.Address, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		suppliers = append(suppliers, s)
	}
	return suppliers, rows.Err()
}

func (r *pgSupplierRepository) Update(ctx context.Context, s *models.Supplier) error {
	query := `
		UPDATE suppliers 
		SET name = $1, contact_person = $2, email = $3, phone_number = $4, address = $5, updated_at = $6 
		WHERE id = $7`
	_, err := r.db.ExecContext(ctx, query, s.Name, s.ContactPerson, s.Email, s.PhoneNumber, s.Address, s.UpdatedAt, s.ID)
	return err
}

func (r *pgSupplierRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM suppliers WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
