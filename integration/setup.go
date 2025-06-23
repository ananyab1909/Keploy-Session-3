package integration

import (
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func SetupTestDB(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to DB: %v", err)
	}
	TestDB = db
	return db
}

func CleanupTestDB() {
	sqlDB, _ := TestDB.DB()
	sqlDB.Close()
}
