package repository_test

import (
	"context"
	"testing"

	"pos-app/backend/internal/models"
)

// createRandomRolePermission adalah helper untuk membuat hubungan role-permission baru.
func createRandomRolePermission(t *testing.T) (*models.RolePermission, *models.Role, *models.Permission) {
	role := createRandomRole(t)
	permission := createRandomPermission(t) // Helper ini sudah ada di user_repository_test.go

	rp := &models.RolePermission{
		RoleID:       role.ID,
		PermissionID: permission.ID,
	}

	err := rolePermissionTestRepo.Create(context.Background(), rp)
	if err != nil {
		t.Fatalf("Gagal membuat role permission random untuk test: %v", err)
	}
	return rp, role, permission
}

func TestRolePermissionRepository_Create(t *testing.T) {
	defer cleanup()
	_, _, _ = createRandomRolePermission(t) // Cukup buat satu untuk memastikan tidak ada error
}

func TestRolePermissionRepository_ListByRoleID(t *testing.T) {
	defer cleanup()

	rp1, role, _ := createRandomRolePermission(t)
	// Buat izin lain untuk peran yang sama
	permission2 := createRandomPermission(t)
	rp2 := &models.RolePermission{
		RoleID:       role.ID,
		PermissionID: permission2.ID,
	}
	err := rolePermissionTestRepo.Create(context.Background(), rp2)
	if err != nil {
		t.Fatalf("Gagal membuat role permission kedua: %v", err)
	}

	// Buat hubungan role-permission untuk peran yang berbeda
	createRandomRolePermission(t)

	foundRps, err := rolePermissionTestRepo.ListByRoleID(context.Background(), role.ID)
	if err != nil {
		t.Fatalf("Gagal list role permissions by role ID: %v", err)
	}

	if len(foundRps) != 2 {
		t.Errorf("Diharapkan 2 role permissions, tetapi mendapatkan %d", len(foundRps))
	}

	// Verifikasi bahwa kedua izin yang dibuat ditemukan
	foundMap := make(map[int32]bool)
	for _, rp := range foundRps {
		foundMap[rp.PermissionID] = true
	}
	if !foundMap[rp1.PermissionID] || !foundMap[rp2.PermissionID] {
		t.Errorf("Tidak semua permission yang diharapkan ditemukan untuk role ID %d", role.ID)
	}
}

func TestRolePermissionRepository_ListByPermissionID(t *testing.T) {
	defer cleanup()

	rp1, _, permission := createRandomRolePermission(t)
	// Buat peran lain untuk izin yang sama
	role2 := createRandomRole(t)
	rp2 := &models.RolePermission{
		RoleID:       role2.ID,
		PermissionID: permission.ID,
	}
	err := rolePermissionTestRepo.Create(context.Background(), rp2)
	if err != nil {
		t.Fatalf("Gagal membuat role permission kedua: %v", err)
	}

	// Buat hubungan role-permission untuk izin yang berbeda
	createRandomRolePermission(t)

	foundRps, err := rolePermissionTestRepo.ListByPermissionID(context.Background(), permission.ID)
	if err != nil {
		t.Fatalf("Gagal list role permissions by permission ID: %v", err)
	}

	if len(foundRps) != 2 {
		t.Errorf("Diharapkan 2 role permissions, tetapi mendapatkan %d", len(foundRps))
	}

	// Verifikasi bahwa kedua peran yang dibuat ditemukan
	foundMap := make(map[int32]bool)
	for _, rp := range foundRps {
		foundMap[rp.RoleID] = true
	}
	if !foundMap[rp1.RoleID] || !foundMap[rp2.RoleID] {
		t.Errorf("Tidak semua role yang diharapkan ditemukan untuk permission ID %d", permission.ID)
	}
}

func TestRolePermissionRepository_Delete(t *testing.T) {
	defer cleanup()

	rpToDelete, _, _ := createRandomRolePermission(t)

	err := rolePermissionTestRepo.Delete(context.Background(), rpToDelete.RoleID, rpToDelete.PermissionID)
	if err != nil {
		t.Fatalf("Gagal menghapus role permission: %v", err)
	}

	// Verifikasi bahwa sudah dihapus
	foundRps, err := rolePermissionTestRepo.ListByRoleID(context.Background(), rpToDelete.RoleID)
	if err != nil {
		t.Fatalf("Gagal list role permissions setelah delete: %v", err)
	}
	for _, rp := range foundRps {
		if rp.PermissionID == rpToDelete.PermissionID {
			t.Errorf("Role permission %d-%d seharusnya sudah dihapus", rpToDelete.RoleID, rpToDelete.PermissionID)
		}
	}
}

func TestRolePermissionRepository_DeleteByRoleID(t *testing.T) {
	defer cleanup()

	_, roleToDelete, _ := createRandomRolePermission(t) // Mengabaikan rp1 karena tidak digunakan
	// Buat izin lain untuk peran yang sama
	permission2 := createRandomPermission(t)
	rp2 := &models.RolePermission{
		RoleID:       roleToDelete.ID,
		PermissionID: permission2.ID,
	}
	err := rolePermissionTestRepo.Create(context.Background(), rp2)
	if err != nil {
		t.Fatalf("Gagal membuat role permission kedua untuk delete by role: %v", err)
	}

	// Buat hubungan role-permission untuk peran yang berbeda (seharusnya tidak dihapus)
	createRandomRolePermission(t)

	err = rolePermissionTestRepo.DeleteByRoleID(context.Background(), roleToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus role permissions by role ID: %v", err)
	}

	// Verifikasi bahwa sudah dihapus
	foundRps, err := rolePermissionTestRepo.ListByRoleID(context.Background(), roleToDelete.ID)
	if err != nil {
		t.Fatalf("Gagal list role permissions setelah delete by role ID: %v", err)
	}
	if len(foundRps) != 0 {
		t.Errorf("Diharapkan 0 role permissions setelah delete by role ID, tetapi mendapatkan %d", len(foundRps))
	}
}
