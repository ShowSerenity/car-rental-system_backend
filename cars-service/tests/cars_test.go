package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"car-rental-system/cars-service/database"
	"car-rental-system/cars-service/handlers"
	"car-rental-system/cars-service/models"
	"github.com/stretchr/testify/assert"
)

func TestAddCar_CarCreated(t *testing.T) {
	database.Connect()
	database.DB.Exec("DELETE FROM cars") // Clean up table before test

	car := models.Car{
		Make:    "Toyota",
		Model:   "Corolla",
		Year:    2020,
		Color:   "Blue",
		Mileage: 10000,
		Price:   20000.00,
	}
	body, _ := json.Marshal(car)
	req, err := http.NewRequest("POST", "/cars", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.AddCar)
	handler.ServeHTTP(rr, req)

	// Verify that the car is created in the database
	var createdCar models.Car
	result := database.DB.Where("make = ?", "Toyota").First(&createdCar)
	assert.NoError(t, result.Error)
	assert.Equal(t, "Corolla", createdCar.Model)
	assert.Equal(t, 2020, createdCar.Year)
	assert.Equal(t, "Blue", createdCar.Color)
	assert.Equal(t, 10000, createdCar.Mileage)
	assert.Equal(t, 20000.00, createdCar.Price)
}

func TestUpdateCar_CarUpdated(t *testing.T) {
	// Set up a test database
	database.Connect()
	defer database.DB.Exec("DELETE FROM cars")

	// Create a new car
	newCar := models.Car{
		Make:    "Toyota",
		Model:   "Corolla",
		Year:    2021,
		Color:   "Red",
		Mileage: 10000,
		Price:   20000.00,
	}

	// Create the car and retrieve its ID
	database.DB.Create(&newCar)
	createdCarID := newCar.ID

	// Define updated car details
	updatedCar := models.Car{
		ID:      createdCarID,
		Make:    "Toyota",
		Model:   "Corolla",
		Year:    2021,
		Color:   "Blue",
		Mileage: 12000,
		Price:   22000.00,
	}

	// Marshal the updated car data to JSON
	updatedCarData, err := json.Marshal(updatedCar)
	assert.NoError(t, err)

	// Send a PUT request to update the car
	req, err := http.NewRequest("PUT", "/cars/"+strconv.Itoa(createdCarID), bytes.NewReader(updatedCarData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	http.HandlerFunc(handlers.UpdateCar).ServeHTTP(recorder, req)

	// Verify that the car was updated successfully
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Retrieve the updated car from the database
	var retrievedCar models.Car
	database.DB.First(&retrievedCar, createdCarID)

	// Verify that the car data is updated correctly
	assert.Equal(t, updatedCar.Color, retrievedCar.Color)
	assert.Equal(t, updatedCar.Mileage, retrievedCar.Mileage)
	assert.Equal(t, updatedCar.Price, retrievedCar.Price)
}
