package repository_test

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"pos-app/backend/internal/models"
)

// createRandomRole adalah fungsi helper untuk membuat dan menyimpan role baru ke DB.
func createRandomRole(t *testing.T) *models.Role {
	role := &models.Role{
		Name:        "Role " + randomString(8),
		Description: sql.NullString{String: "Description for " + randomString(10), Valid: true},
	}

	err := roleTestRepo.Create(context.Background(), role)
	if err != nil {
		t.Fatalf("Gagal membuat role random untuk test: %v", err)
	}
	return role
}

func TestRoleRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	newRole := createRandomRole(t)

	foundRole, err := roleTestRepo.GetByID(context.Background(), newRole.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan role by ID: %v", err)
	}
	if foundRole == nil {
		t.Fatal("Role yang baru dibuat tidak ditemukan")
	}

	if newRole.ID == 0 { // ID harus sudah di-generate oleh DB
		t.Errorf("ID role tidak di-generate. Diharapkan > 0, didapatkan %d", newRole.ID)
	}
	if newRole.Name != foundRole.Name {
		t.Errorf("Name tidak cocok. Diharapkan '%s', didapatkan '%s'", newRole.Name, foundRole.Name)
	}
}

func TestRoleRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()

	nonExistentID := int32(99999) // ID yang sangat besar, kemungkinan tidak ada
	foundRole, err := roleTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundRole != nil {
		t.Errorf("Diharapkan role nil, tetapi mendapatkan: %+v", foundRole)
	}
}

func TestRoleRepository_GetByName(t *testing.T) {
	defer cleanup()

	newRole := createRandomRole(t)

	foundRole, err := roleTestRepo.GetByName(context.Background(), newRole.Name)
	if err != nil {
		t.Fatalf("Gagal mendapatkan role by Name: %v", err)
	}
	if foundRole == nil {
		t.Fatal("Role yang baru dibuat tidak ditemukan berdasarkan nama")
	}

	if newRole.ID != foundRole.ID {
		t.Errorf("ID tidak cocok. Diharapkan '%d', didapatkan '%d'", newRole.ID, foundRole.ID)
	}
	if newRole.Name != foundRole.Name {
		t.Errorf("Name tidak cocok. Diharapkan '%s', didapatkan '%s'", newRole.Name, foundRole.Name)
	}

	// Test GetByName untuk nama yang tidak ada
	_, err = roleTestRepo.GetByName(context.Background(), "NonExistentRole")
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows untuk nama yang tidak ada, tetapi mendapatkan: %v", err)
	}
}

func TestRoleRepository_ListAll(t *testing.T) {
	defer cleanup()

	createRandomRole(t)
	createRandomRole(t)
	createRandomRole(t)

	roles, err := roleTestRepo.ListAll(context.Background())
	if err != nil {
		t.Fatalf("Gagal melakukan list all roles: %v", err)
	}

	if len(roles) != 3 {
		t.Errorf("Diharapkan 3 roles, tetapi mendapatkan %d", len(roles))
	}
}

func TestRoleRepository_Update(t *testing.T) {
	defer cleanup()

	initialRole := createRandomRole(t)

	updatedName := "Updated Role " + randomString(8) // Gunakan nama acak untuk update
	initialRole.Name = updatedName

	err := roleTestRepo.Update(context.Background(), initialRole)
	if err != nil {
		t.Fatalf("Gagal mengupdate role: %v", err)
	}

	foundRole, err := roleTestRepo.GetByID(context.Background(), initialRole.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan role setelah update: %v", err)
	}
	if foundRole.Name != updatedName {
		t.Errorf("Name tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedName, foundRole.Name)
	}
}

func TestRoleRepository_Delete(t *testing.T) {
	defer cleanup()

	roleToDelete := createRandomRole(t)

	err := roleTestRepo.Delete(context.Background(), roleToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus role: %v", err)
	}

	_, err = roleTestRepo.GetByID(context.Background(), roleToDelete.ID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}

// randomString helper function (bisa dipindahkan ke file utilitas jika sering dipakai)
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
