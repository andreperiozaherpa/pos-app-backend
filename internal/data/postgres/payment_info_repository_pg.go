package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

type paymentInfoRepositoryPG struct {
	db *sql.DB
}

func NewPaymentInfoRepositoryPG(db *sql.DB) repository.PaymentInfoRepository {
	return &paymentInfoRepositoryPG{db: db}
}

func (r *paymentInfoRepositoryPG) Create(ctx context.Context, payment *models.PaymentInfo) error {
	query := `
        INSERT INTO payment_info (id, transaction_id, payment_method, amount, payment_date)
        VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query,
		payment.ID,
		payment.TransactionID,
		payment.PaymentMethod,
		payment.Amount,
		payment.PaymentDate,
	)
	return err
}

func (r *paymentInfoRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.PaymentInfo, error) {
	query := `
        SELECT id, transaction_id, payment_method, amount, payment_date
        FROM payment_info WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, id)
	payment := &models.PaymentInfo{}
	err := row.Scan(
		&payment.ID,
		&payment.TransactionID,
		&payment.PaymentMethod,
		&payment.Amount,
		&payment.PaymentDate,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return payment, nil
}

func (r *paymentInfoRepositoryPG) ListByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]*models.PaymentInfo, error) {
	query := `
        SELECT id, transaction_id, payment_method, amount, payment_date
        FROM payment_info WHERE transaction_id=$1 ORDER BY payment_date DESC`
	rows, err := r.db.QueryContext(ctx, query, transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*models.PaymentInfo
	for rows.Next() {
		p := &models.PaymentInfo{}
		if err := rows.Scan(
			&p.ID,
			&p.TransactionID,
			&p.PaymentMethod,
			&p.Amount,
			&p.PaymentDate,
		); err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}
	return payments, nil
}
