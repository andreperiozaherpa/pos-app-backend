package mocks

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// EmployeeRoleRepositoryMock adalah mock untuk postgres.EmployeeRoleRepository.
type EmployeeRoleRepositoryMock struct {
	mock.Mock
}

func (m *EmployeeRoleRepositoryMock) AssignRoleToEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error {
	args := m.Called(ctx, employeeUserID, roleID)
	return args.Error(0)
}

func (m *EmployeeRoleRepositoryMock) RemoveRoleFromEmployee(ctx context.Context, employeeUserID uuid.UUID, roleID int32) error {
	args := m.Called(ctx, employeeUserID, roleID)
	return args.Error(0)
}

func (m *EmployeeRoleRepositoryMock) GetRolesForEmployee(ctx context.Context, employeeUserID uuid.UUID) ([]*models.Role, error) {
	args := m.Called(ctx, employeeUserID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	// Pastikan tipe kembalian sesuai dengan interface: []*models.Role
	if roles, ok := args.Get(0).([]*models.Role); ok {
		return roles, args.Error(1)
	}
	// Fallback jika tipe tidak cocok (misalnya, mock mengembalikan tipe yang salah)
	return nil, args.Error(1)
}
