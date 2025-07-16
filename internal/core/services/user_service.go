package services

import (
	"context"
	"pos-app/backend/internal/models"
	"time"

	"github.com/google/uuid"
)

// UserService mendefinisikan kontrak use case terkait user (karyawan & pelanggan/member).
type UserService interface {
	// RegisterUser untuk membuat user baru (employee/customer).
	RegisterUser(ctx context.Context, user *models.User) (uuid.UUID, error)

	// GetUserByID mengambil user berdasarkan ID.
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)

	// GetUserByUsername mengambil user berdasarkan username.
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)

	// GetUserByEmail mengambil user berdasarkan email.
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	// GetUserByPhoneNumber mengambil user berdasarkan nomor telepon.
	GetUserByPhoneNumber(ctx context.Context, phone string) (*models.User, error)

	// UpdateUser memperbarui data user.
	UpdateUser(ctx context.Context, user *models.User) error

	// DeleteUser menghapus user berdasarkan ID.
	DeleteUser(ctx context.Context, id uuid.UUID) error

	// ListUsersByType mengambil semua user berdasarkan tipe (EMPLOYEE/CUSTOMER).
	ListUsersByType(ctx context.Context, userType string) ([]*models.User, error)

	// AuthenticateUser memvalidasi login (username/email/phone + password).
	AuthenticateUser(ctx context.Context, identity string, password string) (*models.User, error)

	// ChangePassword mengganti password user.
	ChangePassword(ctx context.Context, userID uuid.UUID, newPassword string) error

	// ActivateUser & DeactivateUser untuk blokir user tanpa delete.
	ActivateUser(ctx context.Context, userID uuid.UUID) error
	DeactivateUser(ctx context.Context, userID uuid.UUID) error

	// SearchUsers pencarian user fleksibel.
	SearchUsers(ctx context.Context, query string, userType string, limit int) ([]*models.User, error)

	// ResetPasswordViaEmail mengirimkan email untuk reset password ke user yang lupa password.
	ResetPasswordViaEmail(ctx context.Context, email string) error

	// ResetPasswordViaOTP mengirimkan OTP untuk reset password via nomor HP user.
	ResetPasswordViaOTP(ctx context.Context, phone string) error

	// LockUserAccount mengunci akun user sementara, misal karena alasan keamanan.
	LockUserAccount(ctx context.Context, userID uuid.UUID, reason string, durasi time.Duration) error

	// UnlockUserAccount membuka kunci akun user.
	UnlockUserAccount(ctx context.Context, userID uuid.UUID) error

	// UpdateUserProfilePicture memperbarui foto profil user.
	UpdateUserProfilePicture(ctx context.Context, userID uuid.UUID, pictureData []byte) error

	// ListUserRoles menampilkan daftar role yang dimiliki user.
	ListUserRoles(ctx context.Context, userID uuid.UUID) ([]*models.Role, error)

	// GetUserLoginHistory mengambil histori login user dalam rentang waktu tertentu.
	GetUserLoginHistory(ctx context.Context, userID uuid.UUID, from, to time.Time) ([]*models.UserLoginHistory, error)

	// SendUserVerificationEmail mengirimkan email/wa verifikasi ke user.
	SendUserVerificationEmail(ctx context.Context, userID uuid.UUID, channel string) error

	// VerifyUserEmailToken memverifikasi token email verifikasi user.
	VerifyUserEmailToken(ctx context.Context, userID uuid.UUID, token string) (bool, error)

	// SendUserOTP mengirimkan OTP ke user melalui sms/wa/email.
	SendUserOTP(ctx context.Context, userID uuid.UUID, metode string) error

	// ValidateUserOTP memvalidasi OTP yang diinput user.
	ValidateUserOTP(ctx context.Context, userID uuid.UUID, otp string) (bool, error)

	// ListUserActivityLogs menampilkan daftar aktivitas user dalam aplikasi.
	ListUserActivityLogs(ctx context.Context, userID uuid.UUID, from, to time.Time) ([]*models.ActivityLog, error)

	// EnableTwoFactorAuth mengaktifkan autentikasi dua langkah untuk user.
	EnableTwoFactorAuth(ctx context.Context, userID uuid.UUID, metode string) error

	// DisableTwoFactorAuth menonaktifkan autentikasi dua langkah user.
	DisableTwoFactorAuth(ctx context.Context, userID uuid.UUID) error
}
