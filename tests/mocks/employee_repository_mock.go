package mocks

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type EmployeeRepositoryMock struct {
	mock.Mock
}

func (m *EmployeeRepositoryMock) Create(ctx context.Context, employee *models.Employee) error {
	args := m.Called(ctx, employee)
	return args.Error(0)
}

func (m *EmployeeRepositoryMock) GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Employee, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Employee), args.Error(1)
}

func (m *EmployeeRepositoryMock) ListByCompanyID(ctx context.Context, companyID uuid.UUID) ([]*models.Employee, error) {
	args := m.Called(ctx, companyID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Employee), args.Error(1)
}

func (m *EmployeeRepositoryMock) Update(ctx context.Context, employee *models.Employee) error {
	args := m.Called(ctx, employee)
	return args.Error(0)
}
