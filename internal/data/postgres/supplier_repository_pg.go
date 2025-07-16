package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// supplierRepositoryPG implementasi repository Supplier menggunakan PostgreSQL
type supplierRepositoryPG struct {
	db *sql.DB
}

// NewSupplierRepositoryPG membuat instance baru supplierRepositoryPG
func NewSupplierRepositoryPG(db *sql.DB) repository.SupplierRepository {
	return &supplierRepositoryPG{db: db}
}

func (r *supplierRepositoryPG) Create(ctx context.Context, s *models.Supplier) error {
	query := `INSERT INTO suppliers (id, company_id, name, contact_person, email, phone_number, address, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())`
	_, err := r.db.ExecContext(ctx, query, s.ID, s.CompanyID, s.Name, s.ContactPerson, s.Email, s.PhoneNumber, s.Address)
	return err
}

func (r *supplierRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.Supplier, error) {
	query := `SELECT id, company_id, name, contact_person, email, phone_number, address, created_at, updated_at FROM suppliers WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	s := &models.Supplier{}
	err := row.Scan(&s.ID, &s.CompanyID, &s.Name, &s.ContactPerson, &s.Email, &s.PhoneNumber, &s.Address, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return s, nil
}

func (r *supplierRepositoryPG) Update(ctx context.Context, s *models.Supplier) error {
	query := `UPDATE suppliers SET company_id = $1, name = $2, contact_person = $3, email = $4, phone_number = $5, address = $6, updated_at = NOW() WHERE id = $7`
	result, err := r.db.ExecContext(ctx, query, s.CompanyID, s.Name, s.ContactPerson, s.Email, s.PhoneNumber, s.Address, s.ID)
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

func (r *supplierRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM suppliers WHERE id = $1`
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

// SetActiveStatus tidak digunakan karena field is_active dihapus sesuai ERD

// ArchiveSupplier tidak digunakan karena field is_archived dihapus sesuai ERD

// ListSuppliers mengambil daftar supplier, tanpa filter is_active
func (r *supplierRepositoryPG) ListSuppliers(ctx context.Context, activeOnly bool, limit, offset int) ([]*models.Supplier, error) {
	query := `SELECT id, company_id, name, contact_person, email, phone_number, address, created_at, updated_at FROM suppliers ORDER BY name ASC LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
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
	return suppliers, nil
}

// ListByCompanyID mengambil daftar supplier berdasarkan company_id dengan pagination.
func (r *supplierRepositoryPG) ListByCompanyID(ctx context.Context, companyID uuid.UUID, limit, offset int) ([]*models.Supplier, error) {
	query := `
		SELECT id, company_id, name, contact_person, email, phone_number, address, created_at, updated_at
		FROM suppliers
		WHERE company_id = $1
		ORDER BY name ASC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, query, companyID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []*models.Supplier
	for rows.Next() {
		s := &models.Supplier{}
		err := rows.Scan(
			&s.ID,
			&s.CompanyID,
			&s.Name,
			&s.ContactPerson,
			&s.Email,
			&s.PhoneNumber,
			&s.Address,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, s)
	}

	return suppliers, nil
}
