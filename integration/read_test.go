package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"custom-api-server/handlers"
	"custom-api-server/models"
)

func TestGetUsers_Integration(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB()

	req := httptest.NewRequest("GET", "/users", nil)
	res := httptest.NewRecorder()

	handlers.GetUsers(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", res.Code)
	}

	var users []models.User
	err := json.NewDecoder(res.Body).Decode(&users)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
}
