package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"custom-api-server/handlers"
	"custom-api-server/models"

	"github.com/google/uuid"
)

func TestCreateUser_Integration(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB()

	user := models.User{
		ID:    uuid.New(),
		Name:  "Integration User",
		Email: "integration@example.com",
	}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	handlers.CreateUser(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", res.Code)
	}

	var created models.User
	json.NewDecoder(res.Body).Decode(&created)
	if created.Email != user.Email {
		t.Errorf("Expected email %s, got %s", user.Email, created.Email)
	}
}
