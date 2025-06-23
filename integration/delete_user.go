package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"custom-api-server/handlers"
	"custom-api-server/models"
)

func TestDeleteUser_Integration(t *testing.T) {
	SetupTestDB(t)
	defer CleanupTestDB()

	user := models.User{Name: "DeleteMe", Email: "delete@example.com"}
	if err := TestDB.Create(&user).Error; err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	body, _ := json.Marshal(map[string]interface{}{"id": user.ID})
	req := httptest.NewRequest("DELETE", "/users", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	handlers.DeleteUser(res, req)

	if res.Code != http.StatusNoContent {
		t.Errorf("Expected 204 No Content, got %d", res.Code)
	}

	var deleted models.User
	err := TestDB.First(&deleted, "id = ?", user.ID).Error
	if err == nil {
		t.Errorf("Expected user to be deleted, but still found in DB")
	}
}
