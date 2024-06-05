package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"car-rental-system/auth-service/database"
	"car-rental-system/auth-service/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Server error, unable to create your account.", 500)
		return
	}
	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, "Server error, unable to create your account.", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	var dbUser models.User
	if err := database.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		http.Error(w, "Invalid username or password.", 401)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid username or password.", 401)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := query.Get("filter")
	sortBy := query.Get("sort_by")
	page := query.Get("page")
	limit := query.Get("limit")

	var users []models.User
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

	if err := db.Find(&users).Error; err != nil {
		http.Error(w, "Server error, unable to retrieve users.", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
