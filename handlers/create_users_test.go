package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	t.Cleanup(func() { sqlDB.Close() })

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm db: %v", err)
	}

	return db, mock
}

func TestCreateUserHandler(t *testing.T) {
	db, mock := setupDB(t)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO \"users\"").
		WithArgs("generated-uuid", "John Doe", "john@example.com").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	CreateUser(w, req, db)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response User
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.ID != "generated-uuid" {
		t.Errorf("expected ID 'generated-uuid', got %q", response.ID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %v", err)
	}
}

func TestCreateUserHandler_BadRequest(t *testing.T) {
	db, _ := setupDB(t)

	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer([]byte("{")))
	w := httptest.NewRecorder()

	CreateUser(w, req, db)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status code %d for bad request, got %d",
			http.StatusBadRequest, w.Code)
	}
}
