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

func DeleteUserTestable(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var input struct{ ID string }
	json.NewDecoder(r.Body).Decode(&input)

	db.Delete(&User{}, "id = ?", input.ID)
	w.WriteHeader(http.StatusNoContent)
}

func TestDeleteUser_Mock(t *testing.T) {
	sqlDB, mock, _ := sqlmock.New()
	defer sqlDB.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "users" WHERE id = \$1`).
		WithArgs("123").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	body, _ := json.Marshal(map[string]string{"id": "123"})
	req, _ := http.NewRequest("DELETE", "/users", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	DeleteUserTestable(rr, req, gormDB)

	if rr.Code != http.StatusNoContent {
		t.Errorf("Expected 204 No Content, got %d", rr.Code)
	}
}
