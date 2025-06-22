package repository_test

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"pos-app/backend/internal/data/postgres"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import driver
)

var testDB *sql.DB
var userTestRepo postgres.UserRepository
var companyTestRepo postgres.CompanyRepository
var storeTestRepo postgres.StoreRepository
var businessLineTestRepo postgres.BusinessLineRepository
var employeeTestRepo postgres.EmployeeRepository
var customerTestRepo postgres.CustomerRepository
var roleTestRepo postgres.RoleRepository
var employeeRoleTestRepo postgres.EmployeeRoleRepository
var supplierTestRepo postgres.SupplierRepository
var storeProductTestRepo postgres.StoreProductRepository
var masterProductTestRepo postgres.MasterProductRepository
var transactionTestRepo postgres.TransactionRepository
var purchaseOrderTestRepo postgres.PurchaseOrderRepository
var internalStockTransferTestRepo postgres.InternalStockTransferRepository
var stockMovementTestRepo postgres.StockMovementRepository
var activityLogTestRepo postgres.ActivityLogRepository
var operationalExpenseTestRepo postgres.OperationalExpenseRepository
var appliedItemDiscountTestRepo postgres.AppliedItemDiscountRepository
var rolePermissionTestRepo postgres.RolePermissionRepository
var taxRateTestRepo postgres.TaxRateRepository
var permissionTestRepo postgres.PermissionRepository
var appliedTransactionDiscountTestRepo postgres.AppliedTransactionDiscountRepository
var discountTestRepo postgres.DiscountRepository
var shiftTestRepo postgres.ShiftRepository
var r *rand.Rand

// TestMain adalah fungsi khusus yang dijalankan sebelum dan sesudah semua test dalam package ini.
func TestMain(m *testing.M) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		if err := godotenv.Load("../../../.env"); err != nil {
			log.Fatalf("Error loading .env file for testing: %v", err)
		}
	}

	testDBURL := os.Getenv("DATABASE_URL_TEST")
	if testDBURL == "" {
		log.Fatal("DATABASE_URL_TEST tidak diatur untuk testing")
	}

	var err error
	testDB, err = sql.Open("postgres", testDBURL)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database tes: %v", err)
	}

	cleanup()

	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	userTestRepo = postgres.NewPgUserRepository(testDB)
	companyTestRepo = postgres.NewPgCompanyRepository(testDB)
	storeTestRepo = postgres.NewPgStoreRepository(testDB)
	employeeTestRepo = postgres.NewPgEmployeeRepository(testDB)
	businessLineTestRepo = postgres.NewPgBusinessLineRepository(testDB)
	customerTestRepo = postgres.NewPgCustomerRepository(testDB)
	roleTestRepo = postgres.NewPgRoleRepository(testDB)
	employeeRoleTestRepo = postgres.NewPgEmployeeRoleRepository(testDB)
	storeProductTestRepo = postgres.NewPgStoreProductRepository(testDB)
	supplierTestRepo = postgres.NewPgSupplierRepository(testDB)
	masterProductTestRepo = postgres.NewPgMasterProductRepository(testDB)
	transactionTestRepo = postgres.NewPgTransactionRepository(testDB)
	shiftTestRepo = postgres.NewPgShiftRepository(testDB) // Already DBExecutor
	internalStockTransferTestRepo = postgres.NewPgInternalStockTransferRepository(testDB)
	purchaseOrderTestRepo = postgres.NewPgPurchaseOrderRepository(testDB)
	stockMovementTestRepo = postgres.NewPgStockMovementRepository(testDB)
	activityLogTestRepo = postgres.NewPgActivityLogRepository(testDB)
	operationalExpenseTestRepo = postgres.NewPgOperationalExpenseRepository(testDB)
	appliedItemDiscountTestRepo = postgres.NewPgAppliedItemDiscountRepository(testDB)
	taxRateTestRepo = postgres.NewPgTaxRateRepository(testDB)
	rolePermissionTestRepo = postgres.NewPgRolePermissionRepository(testDB)
	permissionTestRepo = postgres.NewPgPermissionRepository(testDB)
	appliedTransactionDiscountTestRepo = postgres.NewPgAppliedTransactionDiscountRepository(testDB)
	discountTestRepo = postgres.NewPgDiscountRepository(testDB)

	code := m.Run()

	testDB.Close()

	os.Exit(code)
}
