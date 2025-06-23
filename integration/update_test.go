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

func TestUpdateUser_Integration(t *testing.T) {
	db := SetupTestDB(t)
	defer CleanupTestDB()
	user := models.User{Name: "Old", Email: "old@example.com"}
	db.Create(&user)

	updated := models.User{
		ID:    user.ID,
		Name:  "Updated",
		Email: "updated@example.com",
	}
	body, _ := json.Marshal(updated)

	req := httptest.NewRequest("PUT", "/users", bytes.NewBuffer(body))
	res := httptest.NewRecorder()

	handlers.UpdateUser(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", res.Code)
	}

	var response models.User
	json.NewDecoder(res.Body).Decode(&response)
	if response.Name != "Updated" {
		t.Errorf("Expected updated name, got %s", response.Name)
	}
}
