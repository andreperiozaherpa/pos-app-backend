package mock

import (
	"context"
	"database/sql/driver"

	"github.com/stretchr/testify/mock"
)

// DBMock adalah mock untuk objek database/sql DB atau Conn
type DBMock struct {
	mock.Mock
}

func (m *DBMock) ExecContext(ctx context.Context, query string, args ...interface{}) (driver.Result, error) {
	ret := m.Called(ctx, query, args)
	return ret.Get(0).(driver.Result), ret.Error(1)
}

func (m *DBMock) QueryRowContext(ctx context.Context, query string, args ...interface{}) RowMockInterface {
	ret := m.Called(ctx, query, args)
	return ret.Get(0).(RowMockInterface)
}

// RowMockInterface adalah interface minimal untuk baris hasil query
type RowMockInterface interface {
	Scan(dest ...interface{}) error
}

// RowMock implementasi mock untuk hasil query satu baris
type RowMock struct {
	ScanFunc func(dest ...interface{}) error
}

func (r *RowMock) Scan(dest ...interface{}) error {
	return r.ScanFunc(dest...)
}
