package tests

import (
	"bytes"
	"car-rental-system/auth-service/database"
	"car-rental-system/auth-service/handlers"
	"car-rental-system/auth-service/models"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConnect(t *testing.T) {
	err := database.Connect()
	assert.NoError(t, err, "Failed to connect to the database")
	assert.NotNil(t, database.DB, "Database connection is nil")
}

func TestRegister(t *testing.T) {
	database.Connect()
	database.DB.Exec("DELETE FROM users")

	user := models.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "password123",
		FullName: "Test User",
	}
	body, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.Register)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status code 201")

	var createdUser models.User
	result := database.DB.Where("username = ?", "testuser").First(&createdUser)
	assert.NoError(t, result.Error)
	assert.Equal(t, "testuser@example.com", createdUser.Email)
	assert.Equal(t, "Test User", createdUser.FullName)
	assert.NotEmpty(t, createdUser.Password)
}

func TestLogin(t *testing.T) {
	database.Connect()
	database.DB.Exec("DELETE FROM users")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := models.User{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: string(hashedPassword),
		FullName: "Test User",
	}
	database.DB.Create(&user)

	loginUser := models.User{
		Username: "testuser",
		Password: "password123",
	}
	body, _ := json.Marshal(loginUser)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.Login)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")
}

func TestGetUsers(t *testing.T) {
	database.Connect()
	database.DB.Exec("DELETE FROM users")

	users := []models.User{
		{Username: "user1", Email: "user1@example.com", Password: "password1"},
		{Username: "user2", Email: "user2@example.com", Password: "password2"},
	}
	for _, user := range users {
		database.DB.Create(&user)
	}

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetUsers)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")

	var retrievedUsers []models.User
	err = json.NewDecoder(rr.Body).Decode(&retrievedUsers)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(retrievedUsers), "Expected 2 users in the response")
	assert.Equal(t, "user1", retrievedUsers[0].Username)
	assert.Equal(t, "user2", retrievedUsers[1].Username)
}
