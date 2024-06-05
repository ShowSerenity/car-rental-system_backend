package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"car-rental-system/rent-service/database"
	"car-rental-system/rent-service/models"
)

func RentCar(w http.ResponseWriter, r *http.Request) {
	var rental models.Rental
	json.NewDecoder(r.Body).Decode(&rental)
	database.DB.Create(&rental)
	w.WriteHeader(http.StatusCreated)
}

func GetRentalHistory(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userID := query.Get("user_id")
	sortBy := query.Get("sort_by")
	page := query.Get("page")
	limit := query.Get("limit")

	var rentals []models.Rental
	db := database.DB

	if userID != "" {
		db = db.Where("user_id = ?", userID)
	}

	if sortBy != "" {
		db = db.Order(sortBy)
	}

	if page != "" && limit != "" {
		pageNum, _ := strconv.Atoi(page)
		limitNum, _ := strconv.Atoi(limit)
		offset := (pageNum - 1) * limitNum
		db = db.Offset(offset).Limit(limitNum)
	}

	if err := db.Find(&rentals).Error; err != nil {
		http.Error(w, "Server error, unable to retrieve rental history.", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(rentals)
}
