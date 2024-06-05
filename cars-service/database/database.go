package database

import (
	"car-rental-system/cars-service/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := "host=host.docker.internal user=rentaluser password=123456 dbname=car_rental_system sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	DB = db

	if err := DB.AutoMigrate(&models.Car{}); err != nil {
		return fmt.Errorf("failed to perform AutoMigrate: %w", err)
	}

	return nil
}

func ConnectGRPC() error {
	dsn := "host=localhost user=rentaluser password=123456 dbname=car_rental_system sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	DB = db

	if err := DB.AutoMigrate(&models.Car{}); err != nil {
		return fmt.Errorf("failed to perform AutoMigrate: %w", err)
	}

	return nil
}
