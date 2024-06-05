package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"car-rental-system/auth-service/database"
	"car-rental-system/auth-service/models"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAndLogin(t *testing.T) {
	database.Connect()
	database.DB.Exec("DELETE FROM users")

	user := models.User{
		Username: "integrationuser",
		Email:    "integrationuser@example.com",
		Password: "password123",
		FullName: "Integration User",
	}
	body, _ := json.Marshal(user)
	resp, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode, "Expected status code 201")

	var createdUser models.User
	result := database.DB.Where("username = ?", "integrationuser").First(&createdUser)
	assert.NoError(t, result.Error)
	assert.Equal(t, "integrationuser@example.com", createdUser.Email)
	assert.Equal(t, "Integration User", createdUser.FullName)
	assert.NotEmpty(t, createdUser.Password)

	loginUser := models.User{
		Username: "integrationuser",
		Password: "password123",
	}

	loginBody, _ := json.Marshal(loginUser)
	resp, err = http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(loginBody))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
}

func TestGetUsers_DataRetrieval(t *testing.T) {
	database.Connect()
	database.DB.Exec("DELETE FROM users")

	users := []models.User{
		{Username: "user1", Email: "user1@example.com", Password: "password1", FullName: "User One"},
		{Username: "user2", Email: "user2@example.com", Password: "password2", FullName: "User Two"},
	}
	for _, user := range users {
		database.DB.Create(&user)
	}

	resp, err := http.Get("http://localhost:8080/users")
	assert.NoError(t, err)

	var retrievedUsers []models.User
	err = json.NewDecoder(resp.Body).Decode(&retrievedUsers)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(retrievedUsers), "Expected 2 users in the response")
	assert.Equal(t, "user1", retrievedUsers[0].Username)
	assert.Equal(t, "user2", retrievedUsers[1].Username)
}
