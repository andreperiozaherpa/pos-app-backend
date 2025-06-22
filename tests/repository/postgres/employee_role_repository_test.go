package repository_test

import (
	"context"
	"testing"
)

func TestEmployeeRoleRepository_AssignAndGetRoles(t *testing.T) {
	defer cleanup()

	// 1. Buat dependensi
	employee := createRandomEmployee(t)
	role1 := createRandomRole(t)
	role2 := createRandomRole(t)

	// 2. Assign role pertama
	err := employeeRoleTestRepo.AssignRoleToEmployee(context.Background(), employee.UserID, role1.ID)
	if err != nil {
		t.Fatalf("Gagal assign role pertama: %v", err)
	}

	// 3. Assign role kedua
	err = employeeRoleTestRepo.AssignRoleToEmployee(context.Background(), employee.UserID, role2.ID)
	if err != nil {
		t.Fatalf("Gagal assign role kedua: %v", err)
	}

	// 4. Ambil kembali semua role untuk employee tersebut
	assignedRoles, err := employeeRoleTestRepo.GetRolesForEmployee(context.Background(), employee.UserID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan roles untuk employee: %v", err)
	}

	// 5. Verifikasi
	if len(assignedRoles) != 2 {
		t.Fatalf("Diharapkan 2 role, tetapi mendapatkan %d", len(assignedRoles))
	}

	// Cek apakah ID role yang benar ada di dalam hasil
	roleIDs := make(map[int32]bool)
	for _, r := range assignedRoles {
		roleIDs[r.ID] = true
	}

	if !roleIDs[role1.ID] {
		t.Errorf("Role dengan ID %d tidak ditemukan", role1.ID)
	}
	if !roleIDs[role2.ID] {
		t.Errorf("Role dengan ID %d tidak ditemukan", role2.ID)
	}
}

func TestEmployeeRoleRepository_RemoveRole(t *testing.T) {
	defer cleanup()

	// 1. Buat dependensi dan assign role
	employee := createRandomEmployee(t)
	roleToKeep := createRandomRole(t)
	roleToRemove := createRandomRole(t)

	_ = employeeRoleTestRepo.AssignRoleToEmployee(context.Background(), employee.UserID, roleToKeep.ID)
	_ = employeeRoleTestRepo.AssignRoleToEmployee(context.Background(), employee.UserID, roleToRemove.ID)

	// 2. Hapus salah satu role
	err := employeeRoleTestRepo.RemoveRoleFromEmployee(context.Background(), employee.UserID, roleToRemove.ID)
	if err != nil {
		t.Fatalf("Gagal menghapus role dari employee: %v", err)
	}

	// 3. Ambil kembali roles
	assignedRoles, err := employeeRoleTestRepo.GetRolesForEmployee(context.Background(), employee.UserID)
	if err != nil {
		t.Fatalf("Gagal mendapatkan roles setelah menghapus: %v", err)
	}

	// 4. Verifikasi
	if len(assignedRoles) != 1 {
		t.Fatalf("Diharapkan 1 role tersisa, tetapi mendapatkan %d", len(assignedRoles))
	}

	if assignedRoles[0].ID != roleToKeep.ID {
		t.Errorf("Role yang tersisa salah. Diharapkan ID %d, didapatkan %d", roleToKeep.ID, assignedRoles[0].ID)
	}
}
