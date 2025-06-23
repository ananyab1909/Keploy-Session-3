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

func UpdateUserTestable(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var input User
	json.NewDecoder(r.Body).Decode(&input)

	var user User
	db.First(&user, "id = ?", input.ID)

	user.Name = input.Name
	user.Email = input.Email
	db.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

func TestUpdateUser_Mock(t *testing.T) {
	sqlDB, mock, _ := sqlmock.New()
	defer sqlDB.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE id = \$1 ORDER BY "users"."id" LIMIT \$2`).
		WithArgs("123", 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow("123", "Updated Name", "updated@example.com"))

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users"`).
		WithArgs("Updated", "updated@example.com", "123").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	input := User{ID: "123", Name: "Updated", Email: "updated@example.com"}
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("PUT", "/users", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	UpdateUserTestable(rr, req, gormDB)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
}
