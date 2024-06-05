package tests

import (
	"bytes"
	"car-rental-system/rent-service/database"
	"car-rental-system/rent-service/handlers"
	"car-rental-system/rent-service/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRentCar(t *testing.T) {
	database.Connect()
	defer database.DB.Exec("DELETE FROM rentals")

	newRental := models.Rental{
		UserID: 30,
		CarID:  85,
	}

	rentalData, err := json.Marshal(newRental)
	if err != nil {
		t.Fatalf("Failed to marshal rental data: %v", err)
	}

	req, err := http.NewRequest("POST", "/rent", bytes.NewReader(rentalData))
	if err != nil {
		t.Fatalf("Failed to create POST request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.RentCar)
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("Unexpected status code: %d", recorder.Code)
	}
}

func TestGetRentalHistory(t *testing.T) {
	database.Connect()

	testData := []models.Rental{
		{UserID: 30,
			CarID: 85},
		{UserID: 30,
			CarID: 86},
		{UserID: 30,
			CarID: 87},
	}

	for _, rental := range testData {
		if err := database.DB.Create(&rental).Error; err != nil {
			t.Fatalf("Failed to create test rental data: %v", err)
		}
	}

	req, err := http.NewRequest("GET", "/rent/history?user_id=30", nil)
	if err != nil {
		t.Fatalf("Failed to create GET request: %v", err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetRentalHistory)
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("Unexpected status code: %d", recorder.Code)
	}

	var rentals []models.Rental
	if err := json.NewDecoder(recorder.Body).Decode(&rentals); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if len(rentals) != 3 {
		t.Fatalf("Expected 3 rentals, got %d", len(rentals))
	}
}
