package tests

import (
	"bytes"
	"car-rental-system/rent-service/database"
	"car-rental-system/rent-service/handlers"
	"car-rental-system/rent-service/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRentAndGetRentalHistory(t *testing.T) {
	router := mux.NewRouter()
	database.Connect()
	handlers.SetupRoutes(router)

	rental := models.Rental{UserID: 30, CarID: 85}
	body, _ := json.Marshal(rental)

	req, _ := http.NewRequest("POST", "/rent", bytes.NewBuffer(body))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusCreated, response.Code, "Expected response code 201")

	req, _ = http.NewRequest("GET", "/rent/history?user_id=30", nil)
	response = httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Expected response code 200")

	var rentals []models.Rental
	json.NewDecoder(response.Body).Decode(&rentals)
	assert.NotEmpty(t, rentals, "Expected non-empty list of rentals")
}
