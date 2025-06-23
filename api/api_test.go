package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"custom-api-server/models"

	"github.com/google/uuid"
)

const baseURL = "http://localhost:8080/users"

func TestCreateUserAPI(t *testing.T) {
	user := models.User{
		ID:    uuid.New(),
		Name:  "API Test",
		Email: "apitest@example.com",
	}
	body, _ := json.Marshal(user)

	res, err := http.Post(baseURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", res.StatusCode)
	}
}

func TestGetUsersAPI(t *testing.T) {
	res, err := http.Get(baseURL)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", res.StatusCode)
	}
}

func TestUpdateUserAPI(t *testing.T) {
	user := models.User{
		ID:    uuid.New(),
		Name:  "Update Me",
		Email: "update@me.com",
	}
	body, _ := json.Marshal(user)
	http.Post(baseURL, "application/json", bytes.NewBuffer(body))

	user.Name = "Updated!"
	user.Email = "updated@example.com"
	updateBody, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPut, baseURL, bytes.NewBuffer(updateBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", res.StatusCode)
	}
}

func TestDeleteUserAPI(t *testing.T) {
	user := models.User{
		ID:    uuid.New(),
		Name:  "Delete Me",
		Email: "delete@me.com",
	}
	body, _ := json.Marshal(user)
	http.Post(baseURL, "application/json", bytes.NewBuffer(body))

	delBody, _ := json.Marshal(map[string]string{"id": user.ID.String()})
	req, _ := http.NewRequest(http.MethodDelete, baseURL, bytes.NewBuffer(delBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Expected 204 No Content, got %d", res.StatusCode)
	}
}
