package postgres

import (
	"context"
	"database/sql"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// customerRepositoryPG adalah implementasi repository Customer menggunakan PostgreSQL.
type customerRepositoryPG struct {
	db *sql.DB
}

// NewCustomerRepositoryPG membuat instance baru CustomerRepository untuk PostgreSQL.
func NewCustomerRepositoryPG(db *sql.DB) repository.CustomerRepository {
	return &customerRepositoryPG{db: db}
}

// Create menambahkan customer baru ke database.
func (r *customerRepositoryPG) Create(ctx context.Context, c *models.Customer) error {
	query := `
		INSERT INTO customers (
			user_id, company_id, membership_number, join_date, points, tier,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())`

	_, err := r.db.ExecContext(ctx, query,
		c.UserID, c.CompanyID, c.MembershipNumber, c.JoinDate, c.Points, c.Tier,
	)
	return err
}

// GetByUserID mengambil data customer berdasarkan user_id.
func (r *customerRepositoryPG) GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Customer, error) {
	query := `
		SELECT user_id, company_id, membership_number, join_date, points, tier,
			created_at, updated_at
		FROM customers WHERE user_id = $1`

	row := r.db.QueryRowContext(ctx, query, userID)
	var c models.Customer
	err := row.Scan(
		&c.UserID, &c.CompanyID, &c.MembershipNumber, &c.JoinDate, &c.Points, &c.Tier,
		&c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &c, nil
}

// Update memperbarui data customer yang sudah ada.
func (r *customerRepositoryPG) Update(ctx context.Context, c *models.Customer) error {
	query := `
		UPDATE customers
		SET membership_number = $1, join_date = $2, points = $3, tier = $4, updated_at = NOW()
		WHERE user_id = $5`

	_, err := r.db.ExecContext(ctx, query,
		c.MembershipNumber, c.JoinDate, c.Points, c.Tier, c.UserID,
	)
	return err
}

// Delete menghapus customer berdasarkan user_id.
func (r *customerRepositoryPG) Delete(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM customers WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

// GetLoyaltyPoints mendapatkan poin loyalti customer berdasarkan user_id.
func (r *customerRepositoryPG) GetLoyaltyPoints(ctx context.Context, customerID uuid.UUID) (int, error) {
	query := `SELECT points FROM customers WHERE user_id = $1`
	row := r.db.QueryRowContext(ctx, query, customerID)
	var points int
	err := row.Scan(&points)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, repository.ErrNotFound
		}
		return 0, err
	}
	return points, nil
}

// AddLoyaltyPoints menambahkan poin loyalti ke customer.
func (r *customerRepositoryPG) AddLoyaltyPoints(ctx context.Context, customerID uuid.UUID, points int) error {
	query := `
		UPDATE customers
		SET points = points + $1, updated_at = NOW()
		WHERE user_id = $2`
	_, err := r.db.ExecContext(ctx, query, points, customerID)
	return err
}

// ListTransactionHistory mengambil histori transaksi customer dengan pagination.
func (r *customerRepositoryPG) ListTransactionHistory(ctx context.Context, customerID uuid.UUID, limit, offset int) ([]*models.Transaction, error) {
	query := `
		SELECT id, store_id, transaction_code, transaction_date, final_total_amount,
			payment_method, created_at, updated_at
		FROM transactions
		WHERE customer_user_id = $1
		ORDER BY transaction_date DESC
		LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, query, customerID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(
			&t.ID, &t.StoreID, &t.TransactionCode, &t.TransactionDate, &t.FinalTotalAmount,
			&t.PaymentMethod, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

// SearchCustomers mencari customer berdasarkan membership number atau lain-lain dengan pagination.
func (r *customerRepositoryPG) SearchCustomers(ctx context.Context, keyword string, limit, offset int) ([]*models.Customer, error) {
	keyword = "%" + keyword + "%"
	query := `
		SELECT user_id, company_id, membership_number, join_date, points, tier,
			created_at, updated_at
		FROM customers
		WHERE membership_number ILIKE $1
		ORDER BY membership_number
		LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, query, keyword, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*models.Customer
	for rows.Next() {
		var c models.Customer
		err := rows.Scan(
			&c.UserID, &c.CompanyID, &c.MembershipNumber, &c.JoinDate, &c.Points, &c.Tier,
			&c.CreatedAt, &c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}

// GetByMembershipNumber mengambil customer berdasarkan nomor keanggotaan/member.
func (r *customerRepositoryPG) GetByMembershipNumber(ctx context.Context, companyID uuid.UUID, membershipNumber string) (*models.Customer, error) {
	query := `
		SELECT user_id, company_id, membership_number, join_date, points, tier,
			created_at, updated_at
		FROM customers
		WHERE company_id = $1 AND membership_number = $2`

	row := r.db.QueryRowContext(ctx, query, companyID, membershipNumber)
	var c models.Customer
	err := row.Scan(
		&c.UserID, &c.CompanyID, &c.MembershipNumber, &c.JoinDate, &c.Points, &c.Tier,
		&c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &c, nil
}

// ListByCompanyID mengambil semua customer dalam satu perusahaan.
func (r *customerRepositoryPG) ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Customer, error) {
	query := `
		SELECT user_id, company_id, membership_number, join_date, points, tier,
			created_at, updated_at
		FROM customers
		WHERE company_id = $1
		ORDER BY membership_number`

	rows, err := r.db.QueryContext(ctx, query, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*models.Customer
	for rows.Next() {
		var c models.Customer
		err := rows.Scan(
			&c.UserID, &c.CompanyID, &c.MembershipNumber, &c.JoinDate, &c.Points, &c.Tier,
			&c.CreatedAt, &c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}
