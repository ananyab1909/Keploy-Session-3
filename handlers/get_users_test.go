package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetUsersTestable(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func TestGetUsers_Mock(t *testing.T) {
	sqlDB, mock, _ := sqlmock.New()
	defer sqlDB.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow("1", "Alice", "alice@example.com").
		AddRow("2", "Bob", "bob@example.com")

	mock.ExpectQuery(`SELECT \* FROM "users"`).WillReturnRows(rows)

	req, _ := http.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()

	GetUsersTestable(rr, req, gormDB)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
}
