package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// BusinessLineRepository mendefinisikan interface untuk operasi data terkait BusinessLine.
type BusinessLineRepository interface {
	Create(ctx context.Context, bl *models.BusinessLine) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.BusinessLine, error)
	ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.BusinessLine, error)
	Update(ctx context.Context, bl *models.BusinessLine) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// pgBusinessLineRepository adalah implementasi dari BusinessLineRepository untuk PostgreSQL.
type pgBusinessLineRepository struct {
	db *sql.DB
}

// NewPgBusinessLineRepository adalah constructor untuk membuat instance baru dari pgBusinessLineRepository.
func NewPgBusinessLineRepository(db *sql.DB) BusinessLineRepository {
	return &pgBusinessLineRepository{db: db}
}

// Implementasi metode-metode dari interface BusinessLineRepository:

func (r *pgBusinessLineRepository) Create(ctx context.Context, bl *models.BusinessLine) error {
	query := `
		INSERT INTO business_lines (id, company_id, name, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, bl.ID, bl.CompanyID, bl.Name, bl.Description, bl.CreatedAt, bl.UpdatedAt)
	return err
}

func (r *pgBusinessLineRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.BusinessLine, error) {
	bl := &models.BusinessLine{}
	query := `
		SELECT id, company_id, name, description, created_at, updated_at
		FROM business_lines
		WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&bl.ID, &bl.CompanyID, &bl.Name, &bl.Description, &bl.CreatedAt, &bl.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return bl, nil
}

func (r *pgBusinessLineRepository) ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.BusinessLine, error) {
	query := `
		SELECT id, company_id, name, description, created_at, updated_at
		FROM business_lines
		WHERE company_id = $1
		ORDER BY name ASC`
	rows, err := r.db.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var businessLines []*models.BusinessLine
	for rows.Next() {
		bl := &models.BusinessLine{}
		if err := rows.Scan(&bl.ID, &bl.CompanyID, &bl.Name, &bl.Description, &bl.CreatedAt, &bl.UpdatedAt); err != nil {
			return nil, err
		}
		businessLines = append(businessLines, bl)
	}
	return businessLines, rows.Err()
}

func (r *pgBusinessLineRepository) Update(ctx context.Context, bl *models.BusinessLine) error {
	query := `UPDATE business_lines SET name = $1, description = $2, updated_at = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, bl.Name, bl.Description, bl.UpdatedAt, bl.ID)
	return err
}

func (r *pgBusinessLineRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM business_lines WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
