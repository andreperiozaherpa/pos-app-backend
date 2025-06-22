package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"pos-app/backend/internal/models"
)

// createRandomPermission adalah helper untuk membuat izin baru.
func createRandomPermission(t *testing.T) *models.Permission {
	permission := &models.Permission{
		Name:        "test:permission:" + randomString(6),
		Description: sql.NullString{String: "A test permission", Valid: true},
		GroupName:   sql.NullString{String: "Testing", Valid: true},
	}

	err := permissionTestRepo.Create(context.Background(), permission)
	if err != nil {
		t.Fatalf("Gagal membuat permission random untuk test: %v", err)
	}
	if permission.ID == 0 {
		t.Fatal("Permission ID tidak di-populate setelah create")
	}

	return permission
}

func TestPermissionRepository_CreateAndGetByID(t *testing.T) {
	defer cleanup()

	newPermission := createRandomPermission(t)

	foundPermission, err := permissionTestRepo.GetByID(context.Background(), newPermission.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan permission by ID: %v", err)
	}
	if foundPermission == nil {
		t.Fatal("Permission yang baru dibuat tidak ditemukan")
	}

	if newPermission.ID != foundPermission.ID {
		t.Errorf("ID tidak cocok. Diharapkan %d, didapatkan %d", newPermission.ID, foundPermission.ID)
	}
	if newPermission.Name != foundPermission.Name {
		t.Errorf("Name tidak cocok. Diharapkan '%s', didapatkan '%s'", newPermission.Name, foundPermission.Name)
	}
}

func TestPermissionRepository_GetByID_NotFound(t *testing.T) {
	defer cleanup()
	nonExistentID := int32(999999)
	foundPermission, err := permissionTestRepo.GetByID(context.Background(), nonExistentID)

	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows, tetapi mendapatkan: %v", err)
	}
	if foundPermission != nil {
		t.Errorf("Diharapkan permission nil, tetapi mendapatkan: %+v", foundPermission)
	}
}

func TestPermissionRepository_ListAll(t *testing.T) {
	defer cleanup()

	// Buat 3 izin
	createRandomPermission(t)
	createRandomPermission(t)
	createRandomPermission(t)

	permissions, err := permissionTestRepo.ListAll(context.Background())
	if err != nil {
		t.Fatalf("Gagal list permissions: %v", err)
	}

	if len(permissions) != 3 {
		t.Errorf("Diharapkan 3 permissions, tetapi mendapatkan %d", len(permissions))
	}
}

func TestPermissionRepository_Update(t *testing.T) {
	defer cleanup()

	initialPermission := createRandomPermission(t)

	updatedName := "updated:permission:" + randomString(6)
	updatedDescription := "Updated description"
	initialPermission.Name = updatedName
	initialPermission.Description.String = updatedDescription

	err := permissionTestRepo.Update(context.Background(), initialPermission)
	if err != nil {
		t.Fatalf("Gagal mengupdate permission: %v", err)
	}

	foundPermission, err := permissionTestRepo.GetByID(context.Background(), initialPermission.ID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan permission setelah update: %v", err)
	}

	if foundPermission.Name != updatedName {
		t.Errorf("Name tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedName, foundPermission.Name)
	}
	if foundPermission.Description.String != updatedDescription {
		t.Errorf("Description tidak terupdate. Diharapkan '%s', didapatkan '%s'", updatedDescription, foundPermission.Description.String)
	}
}

func TestPermissionRepository_Delete(t *testing.T) {
	defer cleanup()

	permissionToDelete := createRandomPermission(t)

	err := permissionTestRepo.Delete(context.Background(), permissionToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus permission: %v", err)
	}

	_, err = permissionTestRepo.GetByID(context.Background(), permissionToDelete.ID)
	if err != sql.ErrNoRows {
		t.Errorf("Diharapkan error sql.ErrNoRows setelah delete, tetapi mendapatkan: %v", err)
	}
}
