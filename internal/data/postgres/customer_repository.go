package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// CustomerRepository mendefinisikan interface untuk operasi data terkait Customer.
type CustomerRepository interface {
	Create(ctx context.Context, customer *models.Customer) error
	GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Customer, error)
	Update(ctx context.Context, customer *models.Customer) error
	Delete(ctx context.Context, userID uuid.UUID) error
}

// pgCustomerRepository adalah implementasi dari CustomerRepository untuk PostgreSQL.
type pgCustomerRepository struct {
	db *sql.DB
}

// NewPgCustomerRepository adalah constructor untuk membuat instance baru dari pgCustomerRepository.
func NewPgCustomerRepository(db *sql.DB) CustomerRepository {
	return &pgCustomerRepository{db: db}
}

// Implementasi metode-metode dari interface CustomerRepository:

func (r *pgCustomerRepository) Create(ctx context.Context, customer *models.Customer) error {
	query := `
		INSERT INTO customers (user_id, company_id, membership_number, join_date, points, tier, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.ExecContext(ctx, query,
		customer.UserID,
		customer.CompanyID,
		customer.MembershipNumber,
		customer.JoinDate,
		customer.Points,
		customer.Tier,
		customer.CreatedAt,
		customer.UpdatedAt,
	)
	return err
}

func (r *pgCustomerRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Customer, error) {
	customer := &models.Customer{}
	query := `
		SELECT user_id, company_id, membership_number, join_date, points, tier, created_at, updated_at
		FROM customers
		WHERE user_id = $1`
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&customer.UserID,
		&customer.CompanyID,
		&customer.MembershipNumber,
		&customer.JoinDate,
		&customer.Points,
		&customer.Tier,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return customer, nil
}

func (r *pgCustomerRepository) Update(ctx context.Context, customer *models.Customer) error {
	query := `UPDATE customers SET membership_number = $1, join_date = $2, points = $3, tier = $4, updated_at = $5 WHERE user_id = $6`
	_, err := r.db.ExecContext(ctx, query, customer.MembershipNumber, customer.JoinDate, customer.Points, customer.Tier, customer.UpdatedAt, customer.UserID)
	return err
}

func (r *pgCustomerRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM customers WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}
