package main

import (
	"car-rental-system/rent-service/database"
	"car-rental-system/rent-service/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
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
	r.HandleFunc("/rent", handlers.RentCar).Methods("POST")
	r.HandleFunc("/rent/history", handlers.GetRentalHistory).Methods("GET")

	logger.Info("Starting Rent Service on :8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
