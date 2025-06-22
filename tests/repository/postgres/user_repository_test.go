package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // Import driver
)

// cleanup adalah fungsi pembantu untuk membersihkan tabel setelah setiap test.
func cleanup() {
	// Tambahkan penjaga untuk mencegah panic jika testDB belum diinisialisasi.
	if testDB == nil {
		return
	}
	testDB.Exec("DELETE FROM operational_expenses")          // Hapus operational_expenses
	testDB.Exec("DELETE FROM applied_item_discounts")        // Hapus applied_item_discounts
	testDB.Exec("DELETE FROM activity_logs")                 // Hapus activity_logs
	testDB.Exec("DELETE FROM applied_transaction_discounts") // Hapus applied_transaction_discounts
	testDB.Exec("DELETE FROM discounts")                     // Hapus discounts
	testDB.Exec("DELETE FROM stock_movements")               // Hapus movements
	testDB.Exec("DELETE FROM internal_stock_transfer_items") // Hapus item sebelum transfer
	testDB.Exec("DELETE FROM internal_stock_transfers")      // Hapus transfer
	testDB.Exec("DELETE FROM purchase_order_items")          // Hapus item sebelum order
	testDB.Exec("DELETE FROM purchase_orders")               // Hapus order
	testDB.Exec("DELETE FROM shifts")                        // Hapus shifts sebelum employees dan stores
	testDB.Exec("DELETE FROM transaction_items")
	testDB.Exec("DELETE FROM transactions")
	testDB.Exec("DELETE FROM store_products")  // Hapus dulu store_products karena bergantung pada store dan supplier
	testDB.Exec("DELETE FROM master_products") // Hapus master_products setelah store_products
	// Hapus dengan urutan terbalik dari dependensi Foreign Key
	testDB.Exec("DELETE FROM suppliers")      // Tabel store_products bergantung pada ini
	testDB.Exec("DELETE FROM employee_roles") // Hapus dulu tabel pivot
	testDB.Exec("DELETE FROM employees")
	testDB.Exec("DELETE FROM role_permissions") // Hapus sebelum roles dan permissions
	testDB.Exec("DELETE FROM permissions")      // Hapus sebelum roles jika ada foreign key
	testDB.Exec("DELETE FROM customers")        // Hapus customers
	testDB.Exec("DELETE FROM stores")
	testDB.Exec("DELETE FROM tax_rates") // Tambahkan ini untuk membersihkan tax_rates
	testDB.Exec("DELETE FROM business_lines")
	testDB.Exec("DELETE FROM users")
	testDB.Exec("DELETE FROM companies")
	testDB.Exec("DELETE FROM roles") // Pastikan ini tidak dikomentari
}

// TestUserRepository_CreateAndGetByID menguji fungsionalitas Create dan GetByID.
func TestUserRepository_CreateAndGetByID(t *testing.T) {
	// Panggil cleanup di akhir test untuk memastikan database bersih untuk test selanjutnya.
	defer cleanup()

	// 1. Siapkan data user baru
	newUser := &models.User{
		ID:           uuid.New(),
		UserType:     models.UserTypeEmployee,
		Username:     sql.NullString{String: "testuser", Valid: true},
		PasswordHash: sql.NullString{String: "hashedpassword", Valid: true},
		FullName:     sql.NullString{String: "Test User", Valid: true},
		Email:        sql.NullString{String: "test@example.com", Valid: true},
		PhoneNumber:  sql.NullString{String: "123456789", Valid: true},
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// 2. Uji metode Create
	err := userTestRepo.Create(context.Background(), newUser)
	if err != nil {
		t.Fatalf("Gagal membuat user: %v", err)
	}

	// 3. Uji metode GetByID
	foundUser, err := userTestRepo.GetByID(context.Background(), newUser.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan user by ID: %v", err)
	}
	if foundUser == nil {
		t.Fatal("User yang baru dibuat tidak ditemukan")
	}

	// 4. Verifikasi data yang ditemukan
	if foundUser.Username.String != newUser.Username.String {
		t.Errorf("Username tidak cocok. Diharapkan '%s', didapatkan '%s'", newUser.Username.String, foundUser.Username.String)
	}
	if foundUser.Email.String != newUser.Email.String {
		t.Errorf("Email tidak cocok. Diharapkan '%s', didapatkan '%s'", newUser.Email.String, foundUser.Email.String)
	}
}

// TestUserRepository_GetByID_NotFound menguji kasus di mana user tidak ditemukan.
func TestUserRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()

	// 1. Buat ID acak yang pasti tidak ada di database.
	nonExistentID := uuid.New()

	// 2. Panggil GetByID dengan ID tersebut.
	foundUser, err := userTestRepo.GetByID(context.Background(), nonExistentID)

	// 3. Verifikasi hasilnya.
	// Kita mengharapkan error, dan error tersebut harus sql.ErrNoRows.
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	// Kita juga mengharapkan user yang dikembalikan adalah nil.
	if foundUser != nil {
		t.Errorf("Diharapkan user nil, tetapi mendapatkan user: %+v", foundUser)
	}
}

// TestUserRepository_Update menguji fungsionalitas Update.
func TestUserRepository_Update(t *testing.T) {
	defer cleanup()

	// 1. Buat user awal terlebih dahulu.
	initialUser := &models.User{
		ID:        uuid.New(),
		UserType:  models.UserTypeEmployee,
		Username:  sql.NullString{String: "updateuser", Valid: true},
		Email:     sql.NullString{String: "update@example.com", Valid: true},
		FullName:  sql.NullString{String: "Initial Name", Valid: true},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := userTestRepo.Create(context.Background(), initialUser)
	if err != nil {
		t.Fatalf("Gagal membuat user awal untuk diupdate: %v", err)
	}

	// 2. Modifikasi data user.
	updatedFullName := "Updated Full Name"
	initialUser.FullName = sql.NullString{String: updatedFullName, Valid: true}
	initialUser.IsActive = false
	initialUser.UpdatedAt = time.Now() // Di aplikasi nyata, service layer akan mengatur ini.

	// 3. Panggil metode Update.
	err = userTestRepo.Update(context.Background(), initialUser)
	if err != nil {
		t.Fatalf("Gagal mengupdate user: %v", err)
	}

	// 4. Ambil kembali user yang sudah diupdate untuk verifikasi.
	foundUser, err := userTestRepo.GetByID(context.Background(), initialUser.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan user setelah update: %v", err)
	}

	// 5. Verifikasi bahwa data telah berubah.
	if foundUser.FullName.String != updatedFullName {
		t.Errorf("FullName tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedFullName, foundUser.FullName.String)
	}
	if foundUser.IsActive != false {
		t.Errorf("IsActive tidak terupdate. Diharapkan 'false', didapatkan '%t'", foundUser.IsActive)
	}
}

// TestUserRepository_Delete menguji fungsionalitas Delete (soft delete).
func TestUserRepository_Delete(t *testing.T) {
	defer cleanup()

	// 1. Buat user untuk dihapus.
	userToDelete := &models.User{ID: uuid.New(), UserType: models.UserTypeEmployee, IsActive: true}
	err := userTestRepo.Create(context.Background(), userToDelete)
	if err != nil {
		t.Fatalf("Gagal membuat user untuk dihapus: %v", err)
	}

	// 2. Panggil metode Delete.
	err = userTestRepo.Delete(userToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus user: %v", err)
	}

	// 3. Ambil kembali user untuk memeriksa statusnya.
	foundUser, err := userTestRepo.GetByID(context.Background(), userToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan user setelah soft delete: %v", err)
	}

	// 4. Verifikasi bahwa user sekarang tidak aktif.
	if foundUser.IsActive != false {
		t.Errorf("User tidak di-soft-delete dengan benar. Diharapkan IsActive 'false', didapatkan '%t'", foundUser.IsActive)
	}
}
