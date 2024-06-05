package main

import (
	"car-rental-system/auth-service/handlers"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"

	"car-rental-system/auth-service/database"
	"github.com/gorilla/mux"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	// Connect to database
	err := database.Connect()
	if err != nil {
		logger.Fatalf("failed to initialize database: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")

	// Start server
	logger.Info("Starting Auth Service on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
