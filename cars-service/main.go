package main

import (
	"car-rental-system/cars-service/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"

	"car-rental-system/cars-service/database"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	err := database.Connect()
	if err != nil {
		logger.Fatalf("failed to initialize database: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/cars", handlers.GetCars).Methods("GET")
	r.HandleFunc("/cars/{id}", handlers.GetCar).Methods("GET")
	r.HandleFunc("/cars", handlers.AddCar).Methods("POST")
	r.HandleFunc("/cars/{id}", handlers.UpdateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", handlers.DeleteCar).Methods("DELETE")

	logger.Info("Starting Cars Service on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
