package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pos-app/backend/internal/core/repository"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// userRepositoryPG implements repository.UserRepository
type userRepositoryPG struct {
	db *sql.DB
}

// NewUserRepositoryPG membuat instance baru userRepositoryPG
func NewUserRepositoryPG(db *sql.DB) repository.UserRepository {
	return &userRepositoryPG{db: db}
}

// Create menambahkan user baru ke database
func (r *userRepositoryPG) Create(ctx context.Context, user *models.User) error {
	query := `
        INSERT INTO users (id, username, email, phone_number, password_hash, user_type, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
    `
	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Username,
		user.Email,
		user.PhoneNumber,
		user.PasswordHash,
		user.UserType,
	)
	return err
}

// GetByID mengambil user berdasarkan ID
func (r *userRepositoryPG) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `
        SELECT id, username, email, phone_number, password_hash, user_type, created_at, updated_at
        FROM users
        WHERE id = $1
    `
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.PasswordHash,
		&user.UserType,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

// Update memperbarui data user
func (r *userRepositoryPG) Update(ctx context.Context, user *models.User) error {
	query := `
        UPDATE users
        SET username = $1, email = $2, phone_number = $3, password_hash = $4, user_type = $5, updated_at = NOW()
        WHERE id = $6
    `
	res, err := r.db.ExecContext(ctx, query,
		user.Username,
		user.Email,
		user.PhoneNumber,
		user.PasswordHash,
		user.UserType,
		user.ID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// Delete menghapus user berdasarkan ID
func (r *userRepositoryPG) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// GetByUsername mengambil user berdasarkan username
func (r *userRepositoryPG) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `
        SELECT id, username, email, phone_number, password_hash, user_type, created_at, updated_at
        FROM users
        WHERE username = $1
    `
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.PasswordHash,
		&user.UserType,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

// GetByEmail mengambil user berdasarkan email
func (r *userRepositoryPG) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
        SELECT id, username, email, phone_number, password_hash, user_type, created_at, updated_at
        FROM users
        WHERE email = $1
    `
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.PasswordHash,
		&user.UserType,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

// GetByPhoneNumber mengambil user berdasarkan nomor telepon
func (r *userRepositoryPG) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error) {
	query := `
        SELECT id, username, email, phone_number, password_hash, user_type, created_at, updated_at
        FROM users
        WHERE phone_number = $1
    `
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, phoneNumber).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.PasswordHash,
		&user.UserType,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return user, nil
}

// Authenticate memverifikasi user berdasarkan identifier (username/email) dan password hash
func (r *userRepositoryPG) Authenticate(ctx context.Context, identifier string, passwordHash string) (*models.User, error) {
	query := `
        SELECT id, username, email, phone_number, password_hash, user_type, created_at, updated_at
        FROM users
        WHERE (username = $1 OR email = $1 OR phone_number = $1) AND password_hash = $2
    `
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, identifier, passwordHash).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.PasswordHash,
		&user.UserType,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrUnauthorized
		}
		return nil, err
	}
	return user, nil
}

// ChangePassword mengganti password hash user berdasarkan ID
func (r *userRepositoryPG) ChangePassword(ctx context.Context, id uuid.UUID, newPasswordHash string) error {
	query := `
        UPDATE users
        SET password_hash = $1, updated_at = NOW()
        WHERE id = $2
    `
	res, err := r.db.ExecContext(ctx, query, newPasswordHash, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

// ListUsersByType mengambil daftar user berdasarkan tipe user (EMPLOYEE, CUSTOMER)
func (r *userRepositoryPG) ListUsersByType(ctx context.Context, userType string) ([]*models.User, error) {
	query := `
        SELECT id, username, email, phone_number, password_hash, user_type, created_at, updated_at
        FROM users
        WHERE user_type = $1
        ORDER BY created_at DESC
    `
	rows, err := r.db.QueryContext(ctx, query, userType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.PhoneNumber,
			&user.PasswordHash,
			&user.UserType,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
