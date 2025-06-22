package mocks

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/stretchr/testify/mock"
)

// RoleRepositoryMock adalah mock untuk postgres.RoleRepository.
type RoleRepositoryMock struct {
	mock.Mock
}

func (m *RoleRepositoryMock) Create(ctx context.Context, role *models.Role) error {
	args := m.Called(ctx, role)
	return args.Error(0)
}

func (m *RoleRepositoryMock) GetByID(ctx context.Context, id int32) (*models.Role, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Role), args.Error(1)
}

func (m *RoleRepositoryMock) GetByName(ctx context.Context, name string) (*models.Role, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Role), args.Error(1)
}

func (m *RoleRepositoryMock) ListAll(ctx context.Context) ([]*models.Role, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	// Pastikan tipe kembalian sesuai dengan interface: []*models.Role
	if roles, ok := args.Get(0).([]*models.Role); ok {
		return roles, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *RoleRepositoryMock) Update(ctx context.Context, role *models.Role) error {
	args := m.Called(ctx, role)
	return args.Error(0)
}

func (m *RoleRepositoryMock) Delete(ctx context.Context, id int32) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
