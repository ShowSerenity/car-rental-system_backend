package tests

import (
	"bytes"
	"car-rental-system/cars-service/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"car-rental-system/cars-service/database"
	"car-rental-system/cars-service/models"
	"github.com/stretchr/testify/assert"
)

func TestAddAndGetCars_CorrectData(t *testing.T) {
	database.Connect()
	database.DB.Exec("DELETE FROM cars")

	car := models.Car{
		Make:    "Ford",
		Model:   "Mustang",
		Year:    2021,
		Color:   "Yellow",
		Mileage: 5000,
		Price:   35000.00,
	}
	body, _ := json.Marshal(car)
	resp, err := http.Post("http://localhost:8081/cars", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)

	var createdCar models.Car
	result := database.DB.Where("make = ?", "Ford").First(&createdCar)
	assert.NoError(t, result.Error)
	assert.Equal(t, "Mustang", createdCar.Model)
	assert.Equal(t, 2021, createdCar.Year)
	assert.Equal(t, "Yellow", createdCar.Color)
	assert.Equal(t, 5000, createdCar.Mileage)
	assert.Equal(t, 35000.00, createdCar.Price)

	resp, err = http.Get("http://localhost:8081/cars")
	assert.NoError(t, err)

	var retrievedCars []models.Car
	err = json.NewDecoder(resp.Body).Decode(&retrievedCars)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(retrievedCars), "Expected 1 car in the response")
	assert.Equal(t, "Mustang", retrievedCars[0].Model)
	assert.Equal(t, 2021, retrievedCars[0].Year)
	assert.Equal(t, "Yellow", retrievedCars[0].Color)
	assert.Equal(t, 5000, retrievedCars[0].Mileage)
	assert.Equal(t, 35000.00, retrievedCars[0].Price)
}

func TestAddAndUpdateCar(t *testing.T) {
	database.Connect()
	defer database.DB.Exec("DELETE FROM cars")

	newCar := models.Car{
		Make:    "Toyota",
		Model:   "Corolla",
		Year:    2021,
		Color:   "Red",
		Mileage: 10000,
		Price:   20000.00,
	}

	carData, err := json.Marshal(newCar)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/cars", bytes.NewReader(carData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	http.HandlerFunc(handlers.AddCar).ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	var createdCar models.Car
	err = json.NewDecoder(recorder.Body).Decode(&createdCar)
	assert.NoError(t, err)

	createdCar.Color = "Blue"
	createdCar.Mileage = 12000

	updatedCarData, err := json.Marshal(createdCar)
	assert.NoError(t, err)

	updateReq, err := http.NewRequest("PUT", "/cars/"+strconv.Itoa(createdCar.ID), bytes.NewReader(updatedCarData))
	assert.NoError(t, err)
	updateReq.Header.Set("Content-Type", "application/json")
	updateRecorder := httptest.NewRecorder()
	http.HandlerFunc(handlers.UpdateCar).ServeHTTP(updateRecorder, updateReq)

	assert.Equal(t, http.StatusOK, updateRecorder.Code)

	var updatedCar models.Car
	database.DB.First(&updatedCar, createdCar.ID)

	assert.Equal(t, createdCar.ID, updatedCar.ID)
	assert.Equal(t, createdCar.Make, updatedCar.Make)
	assert.Equal(t, createdCar.Model, updatedCar.Model)
	assert.Equal(t, "Blue", updatedCar.Color)
	assert.Equal(t, 12000, updatedCar.Mileage)
	assert.Equal(t, createdCar.Price, updatedCar.Price)
}
