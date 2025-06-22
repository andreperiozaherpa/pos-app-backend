// Package postgres berisi implementasi repository untuk interaksi dengan database PostgreSQL.
package postgres

import (
	"context"
	"database/sql"
	"pos-app/backend/internal/models" // Import package models yang sudah kita buat

	"github.com/google/uuid"
)

// UserRepository mendefinisikan interface untuk operasi data terkait User.
// Ini memungkinkan kita untuk membuat implementasi yang berbeda (misalnya, untuk testing)
// dan memisahkan logika bisnis dari detail akses data.
type UserRepository interface {
	// Create membuat pengguna baru di database. Menggunakan context untuk timeout/cancellation.
	Create(ctx context.Context, user *models.User) error
	// GetByID mengambil pengguna berdasarkan ID. Menggunakan context.
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	// GetByUsername mengambil pengguna berdasarkan username. Menggunakan context.
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	// GetByEmail mengambil pengguna berdasarkan email. Menggunakan context.
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	// Update memperbarui data pengguna yang ada. Menggunakan context.
	Update(ctx context.Context, user *models.User) error
	// Delete menghapus pengguna berdasarkan ID (biasanya soft delete dengan flag is_active).
	// Untuk contoh ini, kita bisa implementasikan hard delete atau soft delete.
	Delete(id uuid.UUID) error
}

// pgUserRepository adalah implementasi dari UserRepository untuk PostgreSQL.
type pgUserRepository struct {
	db DBExecutor // Dependensi ke koneksi database
}

// NewPgUserRepository membuat instance baru dari pgUserRepository.
// Ini adalah constructor function.
func NewPgUserRepository(db DBExecutor) UserRepository {
	return &pgUserRepository{db: db}
}

// Implementasi metode-metode dari interface UserRepository:

// Create mengimplementasikan UserRepository.Create
func (r *pgUserRepository) Create(ctx context.Context, user *models.User) error {
	// Query SQL untuk memasukkan data pengguna baru.
	// Diasumsikan user.ID, user.CreatedAt, dan user.UpdatedAt sudah diisi sebelum memanggil metode ini (biasanya oleh service).
	// PasswordHash juga diasumsikan sudah di-hash oleh service.
	query := `
		INSERT INTO users (id, user_type, username, password_hash, full_name, email, phone_number, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	// Menggunakan ExecContext untuk menjalankan query dengan context.
	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.UserType,
		user.Username,     // sql.NullString akan ditangani dengan benar oleh driver
		user.PasswordHash, // sql.NullString
		user.FullName,     // sql.NullString
		user.Email,        // sql.NullString
		user.PhoneNumber,  // sql.NullString
		user.IsActive,
		user.CreatedAt,
		user.UpdatedAt,
	)
	// Mengembalikan error jika ada.
	return err
}

// GetByID mengimplementasikan UserRepository.GetByID
func (r *pgUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	// Membuat instance User untuk menampung hasil query.
	user := &models.User{}
	// Query SQL untuk mengambil pengguna berdasarkan ID.
	query := `
		SELECT id, user_type, username, password_hash, full_name, email, phone_number, is_active, created_at, updated_at
		FROM users
		WHERE id = $1`

	// Menggunakan QueryRowContext untuk mengambil satu baris.
	// Metode Scan akan memetakan kolom hasil ke field-field struct User.
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.UserType,
		&user.Username,
		&user.PasswordHash,
		&user.FullName,
		&user.Email,
		&user.PhoneNumber,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		// Jika tidak ada baris yang ditemukan, sql.ErrNoRows akan dikembalikan.
		// Mengembalikan sql.ErrNoRows agar lapisan service bisa menanganinya.
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		// Mengembalikan error lain jika terjadi.
		return nil, err
	}
	// Mengembalikan pengguna yang ditemukan.
	return user, nil
}

// GetByUsername mengimplementasikan UserRepository.GetByUsername
func (r *pgUserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, user_type, username, password_hash, full_name, email, phone_number, is_active, created_at, updated_at
		FROM users
		WHERE username = $1`

	// Perlu sql.NullString untuk username karena bisa saja username yang dicari tidak ada
	// dan kita perlu membedakan antara username NULL di DB dengan username yang tidak ditemukan.
	// Namun, karena username di tabel users adalah UNIQUE dan kemungkinan NOT NULL (jika ada),
	// kita bisa asumsikan username yang dicari adalah string non-NULL.
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.UserType,
		&user.Username, // Jika username di DB bisa NULL, pastikan model.User.Username adalah sql.NullString
		&user.PasswordHash,
		&user.FullName,
		&user.Email,
		&user.PhoneNumber,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return user, nil
}

// GetByEmail mengimplementasikan UserRepository.GetByEmail
func (r *pgUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, user_type, username, password_hash, full_name, email, phone_number, is_active, created_at, updated_at
		FROM users
		WHERE email = $1`

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.UserType,
		&user.Username,
		&user.PasswordHash,
		&user.FullName,
		&user.Email, // Jika email di DB bisa NULL, pastikan model.User.Email adalah sql.NullString
		&user.PhoneNumber,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return user, nil
}

// Update mengimplementasikan UserRepository.Update
func (r *pgUserRepository) Update(ctx context.Context, user *models.User) error {
	// Query SQL untuk memperbarui data pengguna.
	// Diasumsikan user.UpdatedAt akan diisi oleh service atau menggunakan NOW() di query.
	query := `
		UPDATE users
		SET user_type = $1, username = $2, password_hash = $3, full_name = $4, email = $5, 
		    phone_number = $6, is_active = $7, updated_at = $8
		WHERE id = $9`

	_, err := r.db.ExecContext(ctx, query,
		user.UserType,
		user.Username,
		user.PasswordHash,
		user.FullName,
		user.Email,
		user.PhoneNumber,
		user.IsActive,
		user.UpdatedAt, // Sebaiknya diisi time.Now() oleh service atau gunakan NOW() di query
		user.ID,
	)
	return err
}

// Delete mengimplementasikan UserRepository.Delete
func (r *pgUserRepository) Delete(id uuid.UUID) error {
	// Contoh implementasi soft delete:
	// Mengubah is_active menjadi false dan memperbarui updated_at.
	// Menggunakan context.Background() jika tidak ada context spesifik yang perlu diteruskan untuk operasi ini.
	// Namun, lebih baik jika context diteruskan dari pemanggil.
	query := `UPDATE users SET is_active = false, updated_at = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(context.Background(), query, id)
	return err

	// Contoh implementasi hard delete (hapus baris dari database):
	// query := `DELETE FROM users WHERE id = $1`
	// _, err := r.db.ExecContext(context.Background(), query, id)
	// return err
}
