package repository_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// randomString helper function
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

// createRandomCompany adalah fungsi helper untuk membuat dan menyimpan company baru ke DB.
func createRandomCompany(t *testing.T) *models.Company {
	contactInfoJSON, err := json.Marshal(map[string]string{"email": "contact@example.com", "phone": "123-456-7890"})
	if err != nil {
		t.Fatalf("Gagal marshal contact info untuk test: %v", err)
	}

	company := &models.Company{
		ID:                   uuid.New(),
		Name:                 "Test Company " + uuid.NewString(),
		Address:              sql.NullString{String: "123 Test St, Testville", Valid: true},
		ContactInfo:          contactInfoJSON,
		TaxIDNumber:          sql.NullString{String: "TAX" + uuid.NewString(), Valid: true},
		DefaultTaxPercentage: sql.NullFloat64{Float64: 11.50, Valid: true},
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	err = companyTestRepo.Create(context.Background(), company)
	if err != nil {
		t.Fatalf("Gagal membuat company random untuk test: %v", err)
	}
	return company
}

// createRandomBusinessLine adalah helper untuk membuat business line baru.
func createRandomBusinessLine(t *testing.T) (*models.BusinessLine, *models.Company) {
	company := createRandomCompany(t)

	bl := &models.BusinessLine{
		ID:          uuid.New(),
		CompanyID:   company.ID,
		Name:        "Business Line " + uuid.NewString(),
		Description: sql.NullString{String: "Test Description", Valid: true},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := businessLineTestRepo.Create(context.Background(), bl)
	if err != nil {
		t.Fatalf("Gagal membuat business line random untuk test: %v", err)
	}

	return bl, company
}

// createRandomBusinessLineForCompany adalah helper spesifik untuk test list
func createRandomBusinessLineForCompany(t *testing.T, companyID uuid.UUID) {
	bl := &models.BusinessLine{ID: uuid.New(), CompanyID: companyID, Name: "Another BL", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err := businessLineTestRepo.Create(context.Background(), bl)
	if err != nil {
		t.Fatalf("Gagal membuat business line untuk company spesifik: %v", err)
	}
}

// createTestCompanyAndBusinessLine adalah helper untuk membuat dependensi yang dibutuhkan oleh Store.
func createTestCompanyAndBusinessLine(t *testing.T) uuid.UUID {
	businessLine, _ := createRandomBusinessLine(t)
	return businessLine.ID
}

// createRandomStore adalah fungsi helper untuk membuat dan menyimpan store baru ke DB.
func createRandomStore(t *testing.T, businessLineID uuid.UUID) *models.Store {
	store := &models.Store{
		ID:             uuid.New(),
		BusinessLineID: businessLineID,
		ParentStoreID:  uuid.NullUUID{Valid: false}, // Tidak ada parent untuk test sederhana
		Name:           "Test Store " + uuid.NewString(),
		StoreCode:      sql.NullString{String: "TS-" + uuid.NewString()[:6], Valid: true},
		StoreType:      models.StoreTypeCabang,
		Address:        sql.NullString{String: "456 Store Ave", Valid: true},
		PhoneNumber:    sql.NullString{String: "987654321", Valid: true},
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := storeTestRepo.Create(context.Background(), store)
	if err != nil {
		t.Fatalf("Gagal membuat store random untuk test: %v", err)
	}
	return store
}

// createEmployeeDependencies adalah helper untuk membuat semua data yang dibutuhkan oleh seorang Employee.
func createEmployeeDependencies(t *testing.T) (userID, companyID, storeID uuid.UUID) {
	user := &models.User{
		ID:        uuid.New(),
		UserType:  models.UserTypeEmployee,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := userTestRepo.Create(context.Background(), user)
	if err != nil {
		t.Fatalf("Gagal membuat user dependency untuk test employee: %v", err)
	}

	businessLine, company := createRandomBusinessLine(t)
	store := createRandomStore(t, businessLine.ID)

	return user.ID, company.ID, store.ID
}

// createRandomEmployee adalah fungsi helper untuk membuat dan menyimpan employee baru ke DB.
func createRandomEmployee(t *testing.T) *models.Employee {
	userID, companyID, storeID := createEmployeeDependencies(t)

	employee := &models.Employee{
		UserID:           userID,
		CompanyID:        companyID,
		StoreID:          uuid.NullUUID{UUID: storeID, Valid: true},
		EmployeeIDNumber: sql.NullString{String: "EMP-" + uuid.NewString()[:8], Valid: true},
		JoinDate:         sql.NullTime{Time: time.Now(), Valid: true},
		Position:         sql.NullString{String: "Kasir", Valid: true},
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err := employeeTestRepo.Create(context.Background(), employee)
	if err != nil {
		t.Fatalf("Gagal membuat employee random untuk test: %v", err)
	}
	return employee
}

// createRandomCustomer adalah fungsi helper untuk membuat dan menyimpan customer baru ke DB.
func createRandomCustomer(t *testing.T) *models.Customer {
	user := &models.User{
		ID:        uuid.New(),
		UserType:  models.UserTypeCustomer,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := userTestRepo.Create(context.Background(), user)
	if err != nil {
		t.Fatalf("Gagal membuat user dependency untuk test customer: %v", err)
	}

	company := createRandomCompany(t)

	customer := &models.Customer{
		UserID:           user.ID,
		CompanyID:        company.ID,
		MembershipNumber: sql.NullString{String: "MEM-" + uuid.NewString()[:8], Valid: true},
		JoinDate:         sql.NullTime{Time: time.Now(), Valid: true},
		Points:           100,
		Tier:             sql.NullString{String: "Gold", Valid: true},
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err = customerTestRepo.Create(context.Background(), customer)
	if err != nil {
		t.Fatalf("Gagal membuat customer random untuk test: %v", err)
	}
	return customer
}

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

// createRandomSupplier adalah helper untuk membuat supplier baru.
func createRandomSupplier(t *testing.T) (*models.Supplier, *models.Company) {
	company := createRandomCompany(t)

	supplier := &models.Supplier{
		ID:            uuid.New(),
		CompanyID:     company.ID,
		Name:          "Supplier " + randomString(8),
		ContactPerson: sql.NullString{String: "John Doe", Valid: true},
		Email:         sql.NullString{String: "supplier@example.com", Valid: true},
		PhoneNumber:   sql.NullString{String: "111-222-3333", Valid: true},
		Address:       sql.NullString{String: "123 Supplier Lane", Valid: true},
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err := supplierTestRepo.Create(context.Background(), supplier)
	if err != nil {
		t.Fatalf("Gagal membuat supplier random untuk test: %v", err)
	}

	return supplier, company
}

// createRandomMasterProduct adalah helper untuk membuat master product sebagai dependensi.
func createRandomMasterProduct(t *testing.T, companyID uuid.UUID) *models.MasterProduct {
	mp := &models.MasterProduct{
		ID:                uuid.New(),
		CompanyID:         companyID,
		MasterProductCode: "MPC-" + randomString(6),
		Name:              "Master Product " + randomString(8),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	err := masterProductTestRepo.Create(context.Background(), mp)
	if err != nil {
		t.Fatalf("Gagal membuat master product dependency: %v", err)
	}
	return mp
}

// createRandomStoreProduct adalah helper untuk membuat produk baru.
func createRandomStoreProduct(t *testing.T) (*models.StoreProduct, *models.Store, *models.Supplier) {
	businessLine, _ := createRandomBusinessLine(t)
	store := createRandomStore(t, businessLine.ID)

	supplier, _ := createRandomSupplier(t)
	masterProduct := createRandomMasterProduct(t, supplier.CompanyID)

	product := &models.StoreProduct{
		ID:                uuid.New(),
		MasterProductID:   masterProduct.ID,
		StoreID:           store.ID,
		SupplierID:        uuid.NullUUID{UUID: supplier.ID, Valid: true},
		StoreSpecificSKU:  sql.NullString{String: "SKU-" + randomString(6), Valid: true},
		PurchasePrice:     100.50,
		SellingPrice:      150.75,
		WholesalePrice:    sql.NullFloat64{Float64: 120.00, Valid: true},
		Stock:             50,
		MinimumStockLevel: sql.NullInt32{Int32: 10, Valid: true},
		ExpiryDate:        sql.NullTime{Time: time.Now().AddDate(1, 0, 0), Valid: true},
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	err := storeProductTestRepo.Create(context.Background(), product)
	if err != nil {
		t.Fatalf("Gagal membuat produk random untuk test: %v", err)
	}

	return product, store, supplier
}

// createRandomTaxRate adalah helper untuk membuat tax rate baru.
func createRandomTaxRate(t *testing.T, companyID uuid.UUID) *models.TaxRate {
	tr := &models.TaxRate{
		CompanyID:      companyID,
		Name:           "Tax " + randomString(5),
		RatePercentage: float64(r.Intn(2000)) / 100, // Random percentage up to 20.00%
		Description:    sql.NullString{String: "Random tax rate", Valid: true},
		IsActive:       true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := taxRateTestRepo.Create(context.Background(), tr)
	if err != nil {
		t.Fatalf("Gagal membuat tax rate random untuk test: %v", err)
	}
	if tr.ID == 0 {
		t.Fatal("TaxRate ID tidak di-populate setelah create")
	}
	return tr
}

// createRandomDiscount adalah helper untuk membuat diskon baru.
func createRandomDiscount(t *testing.T) (*models.Discount, *models.Company) {
	company := createRandomCompany(t)

	discount := &models.Discount{
		ID:                        uuid.New(),
		CompanyID:                 company.ID,
		Name:                      "Diskon Lebaran " + randomString(5),
		Description:               sql.NullString{String: "Diskon spesial hari raya", Valid: true},
		DiscountType:              models.DiscountTypePercentage,
		DiscountValue:             10.0, // 10%
		ApplicableTo:              models.DiscountApplicableToTotalTransaction,
		MasterProductIDApplicable: uuid.NullUUID{},
		StoreProductIDApplicable:  uuid.NullUUID{},
		CategoryApplicable:        sql.NullString{},
		CustomerTierApplicable:    sql.NullString{},
		MinPurchaseAmount:         sql.NullFloat64{Float64: 100000, Valid: true},
		StartDate:                 time.Now(),
		EndDate:                   time.Now().AddDate(0, 1, 0),
		IsActive:                  true,
		CreatedAt:                 time.Now(),
		UpdatedAt:                 time.Now(),
	}

	err := discountTestRepo.Create(context.Background(), discount)
	if err != nil {
		t.Fatalf("Gagal membuat diskon random untuk test: %v", err)
	}

	return discount, company
}

// createRandomTransaction adalah helper untuk membuat transaksi baru dengan item.
func createRandomTransaction(t *testing.T) *models.Transaction {
	cashier := createRandomEmployee(t)
	customer := createRandomCustomer(t)
	storeProduct, _, _ := createRandomStoreProduct(t)

	item1 := models.TransactionItem{
		ID:                         uuid.New(),
		StoreProductID:             storeProduct.ID,
		Quantity:                   2,
		PricePerUnitAtTransaction:  storeProduct.SellingPrice,
		ItemSubtotalBeforeDiscount: storeProduct.SellingPrice * 2,
		ItemDiscountAmount:         0,
		ItemSubtotalAfterDiscount:  storeProduct.SellingPrice * 2,
		AppliedTaxRateID:           sql.NullInt32{},
		AppliedTaxRatePercentage:   sql.NullFloat64{},
		TaxAmountForItem:           0,
		ItemFinalTotal:             storeProduct.SellingPrice * 2,
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
	}

	finalTotal := item1.ItemFinalTotal
	transaction := &models.Transaction{
		ID:                             uuid.New(),
		TransactionCode:                "TRX-" + randomString(10),
		StoreID:                        cashier.StoreID.UUID,
		CashierEmployeeUserID:          cashier.UserID,
		CustomerUserID:                 uuid.NullUUID{UUID: customer.UserID, Valid: true},
		ActiveShiftID:                  uuid.NullUUID{},
		TransactionDate:                time.Now(),
		SubtotalAmount:                 finalTotal,
		TotalItemDiscountAmount:        0,
		SubtotalAfterItemDiscounts:     finalTotal,
		TransactionLevelDiscountAmount: 0,
		TaxableAmount:                  finalTotal,
		TotalTaxAmount:                 0,
		FinalTotalAmount:               finalTotal,
		ReceivedAmount:                 finalTotal + 10000,
		ChangeAmount:                   10000,
		PaymentMethod:                  sql.NullString{String: "CASH", Valid: true},
		Notes:                          sql.NullString{},
		CreatedAt:                      time.Now(),
		UpdatedAt:                      time.Now(),
		Items:                          []models.TransactionItem{item1},
	}

	err := transactionTestRepo.Create(context.Background(), transaction)
	if err != nil {
		t.Fatalf("Gagal membuat transaksi random untuk test: %v", err)
	}

	return transaction
}

// createRandomInternalStockTransfer adalah helper untuk membuat internal stock transfer baru dengan item.
func createRandomInternalStockTransfer(t *testing.T) *models.InternalStockTransfer {
	// 1. Buat semua dependensi
	businessLine, company := createRandomBusinessLine(t)
	sourceStore := createRandomStore(t, businessLine.ID)
	destinationStore := createRandomStore(t, businessLine.ID)
	storeProduct, _, _ := createRandomStoreProduct(t)
	requestedByUser := createRandomEmployee(t)

	// 2. Buat item transfer
	item1 := models.InternalStockTransferItem{
		ID:                   uuid.New(),
		SourceStoreProductID: storeProduct.ID,
		QuantityRequested:    5,
		QuantityShipped:      0,
		QuantityReceived:     0,
		Notes:                sql.NullString{String: "Item for transfer", Valid: true},
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	// 3. Buat header transfer
	ist := &models.InternalStockTransfer{
		ID:                 uuid.New(),
		TransferCode:       "IST-" + randomString(8),
		CompanyID:          company.ID,
		SourceStoreID:      sourceStore.ID,
		DestinationStoreID: destinationStore.ID,
		TransferDate:       time.Now(),
		Status:             models.StockTransferStatusPending,
		Notes:              sql.NullString{String: "Transfer request", Valid: true},
		RequestedByUserID:  uuid.NullUUID{UUID: requestedByUser.UserID, Valid: true},
		ApprovedByUserID:   uuid.NullUUID{},
		ShippedByUserID:    uuid.NullUUID{},
		ReceivedByUserID:   uuid.NullUUID{},
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		Items:              []models.InternalStockTransferItem{item1},
	}

	// 4. Panggil repository untuk membuat transfer
	err := internalStockTransferTestRepo.Create(context.Background(), ist)
	if err != nil {
		t.Fatalf("Gagal membuat internal stock transfer random untuk test: %v", err)
	}

	return ist
}

// createRandomOperationalExpense is a helper to create a random operational expense.
func createRandomOperationalExpense(t *testing.T) *models.OperationalExpense {
	businessLine, company := createRandomBusinessLine(t)
	store := createRandomStore(t, businessLine.ID)
	user := createRandomEmployee(t)

	expense := &models.OperationalExpense{
		ID:              uuid.New(),
		CompanyID:       company.ID,
		StoreID:         uuid.NullUUID{UUID: store.ID, Valid: true},
		ExpenseDate:     time.Now().AddDate(0, 0, -5),
		Category:        "Gaji Karyawan",
		Description:     sql.NullString{String: "Gaji bulanan karyawan", Valid: true},
		Amount:          1500000.00,
		CreatedByUserID: uuid.NullUUID{UUID: user.UserID, Valid: true},
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err := operationalExpenseTestRepo.Create(context.Background(), expense)
	if err != nil {
		t.Fatalf("Gagal membuat operational expense random untuk test: %v", err)
	}

	return expense
}

// createRandomPurchaseOrder adalah helper untuk membuat purchase order baru dengan item.
func createRandomPurchaseOrder(t *testing.T) *models.PurchaseOrder {
	// 1. Buat semua dependensi
	store := createRandomStore(t, createTestCompanyAndBusinessLine(t))
	supplier, _ := createRandomSupplier(t)
	masterProduct := createRandomMasterProduct(t, supplier.CompanyID)
	user := createRandomEmployee(t) // User yang membuat PO

	// 2. Buat item purchase order
	item1 := models.PurchaseOrderItem{
		ID:                   uuid.New(),
		MasterProductID:      masterProduct.ID,
		QuantityOrdered:      10,
		PurchasePricePerUnit: 50.00,
		QuantityReceived:     0,
		Subtotal:             500.00,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	// 3. Buat header purchase order
	totalAmount := item1.Subtotal
	po := &models.PurchaseOrder{
		ID:                   uuid.New(),
		StoreID:              store.ID,
		SupplierID:           supplier.ID,
		OrderDate:            time.Now(),
		ExpectedDeliveryDate: sql.NullTime{Time: time.Now().AddDate(0, 0, 7), Valid: true},
		Status:               models.POStatusPending,
		TotalAmount:          sql.NullFloat64{Float64: totalAmount, Valid: true},
		Notes:                sql.NullString{String: "Initial order", Valid: true},
		CreatedByUserID:      user.UserID,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
		Items:                []models.PurchaseOrderItem{item1},
	}

	// 4. Panggil repository untuk membuat purchase order
	err := purchaseOrderTestRepo.Create(context.Background(), po)
	if err != nil {
		t.Fatalf("Gagal membuat purchase order random untuk test: %v", err)
	}

	return po
}
