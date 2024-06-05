package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"car-rental-system/cars-service/database"
	"car-rental-system/cars-service/models"
	"github.com/gorilla/mux"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := query.Get("filter")
	sortBy := query.Get("sort_by")
	page := query.Get("page")
	limit := query.Get("limit")

	fmt.Println("Filter:", filter)

	var cars []models.Car
	db := database.DB

	if filter != "" {
		db = db.Where(filter)
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

	if err := db.Find(&cars).Error; err != nil {
		http.Error(w, "Server error, unable to retrieve cars.", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cars)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var car models.Car
	database.DB.First(&car, id)
	json.NewEncoder(w).Encode(car)
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = database.DB.Create(&car).Error
	if err != nil {
		http.Error(w, "Failed to create car", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var car models.Car
	database.DB.First(&car, id)
	json.NewDecoder(r.Body).Decode(&car)
	database.DB.Save(&car)
	w.WriteHeader(http.StatusOK)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	database.DB.Delete(&models.Car{}, id)
	w.WriteHeader(http.StatusOK)
}
