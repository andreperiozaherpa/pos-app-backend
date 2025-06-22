package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// createRandomActivityLog is a helper to create a random activity log.
func createRandomActivityLog(t *testing.T) *models.ActivityLog {
	// Create dependencies
	employee := createRandomEmployee(t) // This creates user, company, store
	companyID := employee.CompanyID     // employee.CompanyID sudah bertipe uuid.UUID
	storeID := employee.StoreID.UUID

	log := &models.ActivityLog{
		UserID:         uuid.NullUUID{UUID: employee.UserID, Valid: true},
		CompanyID:      uuid.NullUUID{UUID: companyID, Valid: true},
		StoreID:        uuid.NullUUID{UUID: storeID, Valid: true},
		ActionType:     "USER_LOGIN",
		Description:    fmt.Sprintf("User %s logged in", employee.UserID.String()),
		TargetEntity:   sql.NullString{String: "User", Valid: true},
		TargetEntityID: uuid.NullUUID{UUID: employee.UserID, Valid: true},
		IPAddress:      sql.NullString{String: "192.168.1.1", Valid: true},
		UserAgent:      sql.NullString{String: "Mozilla/5.0", Valid: true},
		LogTime:        time.Now(),
	}

	err := activityLogTestRepo.Create(context.Background(), log)
	if err != nil {
		t.Fatalf("Failed to create random activity log for test: %v", err)
	}
	// The ID field of the log struct will be populated by the Create method
	// because of the RETURNING id clause in the SQL query.
	if log.ID == 0 {
		t.Fatalf("ActivityLog ID was not populated after creation")
	}

	return log
}

func TestActivityLogRepository_CreateAndListByUserID(t *testing.T) {
	defer cleanup()

	newLog := createRandomActivityLog(t)

	// Create another log for the same user
	log2 := &models.ActivityLog{
		UserID:         newLog.UserID, // Same user
		CompanyID:      newLog.CompanyID,
		StoreID:        newLog.StoreID,
		ActionType:     "PRODUCT_UPDATE",
		Description:    "Updated product XYZ",
		TargetEntity:   sql.NullString{String: "Product", Valid: true},
		TargetEntityID: uuid.NullUUID{UUID: uuid.New(), Valid: true},
		IPAddress:      sql.NullString{String: "192.168.1.1", Valid: true},
		UserAgent:      sql.NullString{String: "Mozilla/5.0", Valid: true},
		LogTime:        time.Now().Add(time.Minute), // Ensure different time for ordering
	}
	err := activityLogTestRepo.Create(context.Background(), log2)
	if err != nil {
		t.Fatalf("Failed to create second activity log: %v", err)
	}

	// Create a log for a different user
	createRandomActivityLog(t)

	foundLogs, err := activityLogTestRepo.ListByUserID(context.Background(), newLog.UserID.UUID)
	if err != nil {
		t.Fatalf("Failed to list activity logs by user ID: %v", err)
	}

	if len(foundLogs) != 2 {
		t.Errorf("Expected 2 activity logs for user, got %d", len(foundLogs))
	}

	// Verify order (most recent first)
	if foundLogs[0].ID != log2.ID {
		t.Errorf("Expected most recent log to be first. Expected %d, got %d", log2.ID, foundLogs[0].ID)
	}
	if foundLogs[1].ID != newLog.ID {
		t.Errorf("Expected second log to be second. Expected %d, got %d", newLog.ID, foundLogs[1].ID)
	}

	// Verify content of the first log
	if foundLogs[0].ActionType != log2.ActionType {
		t.Errorf("ActionType mismatch. Expected '%s', got '%s'", log2.ActionType, foundLogs[0].ActionType)
	}
}

func TestActivityLogRepository_ListByUserID_Empty(t *testing.T) {
	defer cleanup()
	nonExistentUserID := uuid.New()
	foundLogs, err := activityLogTestRepo.ListByUserID(context.Background(), nonExistentUserID)
	if err != nil {
		t.Fatalf("Failed to list activity logs by non-existent user ID: %v", err)
	}
	if len(foundLogs) != 0 {
		t.Errorf("Expected 0 activity logs, got %d", len(foundLogs))
	}
}

func TestActivityLogRepository_ListByCompanyID(t *testing.T) {
	defer cleanup()

	log1 := createRandomActivityLog(t)
	companyID := log1.CompanyID.UUID

	// Create another log for the same company, different user
	employee2 := createRandomEmployee(t)
	employee2.CompanyID = companyID // employee2.CompanyID bertipe uuid.UUID, companyID juga uuid.UUID
	log3 := &models.ActivityLog{
		UserID:      uuid.NullUUID{UUID: employee2.UserID, Valid: true},
		CompanyID:   uuid.NullUUID{UUID: companyID, Valid: true},
		StoreID:     employee2.StoreID,
		ActionType:  "REPORT_GENERATED",
		Description: "Sales report generated",
		LogTime:     time.Now().Add(2 * time.Minute),
	}
	err := activityLogTestRepo.Create(context.Background(), log3)
	if err != nil {
		t.Fatalf("Failed to create third activity log: %v", err)
	}

	// Create a log for a different company
	createRandomActivityLog(t)

	foundLogs, err := activityLogTestRepo.ListByCompanyID(context.Background(), companyID)
	if err != nil {
		t.Fatalf("Failed to list activity logs by company ID: %v", err)
	}

	if len(foundLogs) != 2 {
		t.Errorf("Expected 2 activity logs for company, got %d", len(foundLogs))
	}
	// Verify order (most recent first)
	if foundLogs[0].ID != log3.ID {
		t.Errorf("Expected most recent log for company to be first. Expected %d, got %d", log3.ID, foundLogs[0].ID)
	}
}

func TestActivityLogRepository_ListByStoreID(t *testing.T) {
	defer cleanup()

	log1 := createRandomActivityLog(t)
	storeID := log1.StoreID.UUID

	// Create another log for the same store, different user
	employee2 := createRandomEmployee(t)
	employee2.StoreID = uuid.NullUUID{UUID: storeID, Valid: true} // Ensure same store
	log3 := &models.ActivityLog{
		UserID:      uuid.NullUUID{UUID: employee2.UserID, Valid: true},    // ActivityLog.UserID adalah uuid.NullUUID
		CompanyID:   uuid.NullUUID{UUID: employee2.CompanyID, Valid: true}, // ActivityLog.CompanyID adalah uuid.NullUUID, employee2.CompanyID adalah uuid.UUID
		StoreID:     uuid.NullUUID{UUID: storeID, Valid: true},
		ActionType:  "STOCK_ADJUSTMENT",
		Description: "Adjusted stock for product ABC",
		LogTime:     time.Now().Add(3 * time.Minute),
	}
	err := activityLogTestRepo.Create(context.Background(), log3)
	if err != nil {
		t.Fatalf("Failed to create fourth activity log: %v", err)
	}

	// Create a log for a different store
	createRandomActivityLog(t)

	foundLogs, err := activityLogTestRepo.ListByStoreID(context.Background(), storeID)
	if err != nil {
		t.Fatalf("Failed to list activity logs by store ID: %v", err)
	}

	if len(foundLogs) != 2 {
		t.Errorf("Expected 2 activity logs for store, got %d", len(foundLogs))
	}
	// Verify order (most recent first)
	if foundLogs[0].ID != log3.ID {
		t.Errorf("Expected most recent log for store to be first. Expected %d, got %d", log3.ID, foundLogs[0].ID)
	}
}
