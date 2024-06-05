package main

import (
	"car-rental-system/cars-service/models"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"net"

	"car-rental-system/cars-service/database"
	pb "car-rental-system/cars-service/proto"
	"google.golang.org/grpc"
)

// CarServiceServerImpl implements the CarServiceServer interface.
type CarServiceServerImpl struct {
	pb.UnimplementedCarServiceServer
}

func (s *CarServiceServerImpl) GetCars(ctx context.Context, req *pb.GetCarsRequest) (*pb.GetCarsResponse, error) {
	var cars []*pb.Car

	filter := req.GetFilter()
	sortBy := req.GetSortBy()
	page := int(req.GetPage())
	limit := int(req.GetLimit())

	query := database.DB.Model(&models.Car{})

	if filter != "" {
		query = query.Where(filter)
	}

	if sortBy != "" {
		query = query.Order(sortBy)
	}

	if page != 0 && limit != 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	var dbCars []models.Car
	if err := query.Find(&dbCars).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch cars from database: %v", err)
	}

	for _, car := range dbCars {
		pbCar := &pb.Car{
			Id:      int32(car.ID),
			Make:    car.Make,
			Model:   car.Model,
			Year:    int32(car.Year),
			Color:   car.Color,
			Mileage: int32(car.Mileage),
			Price:   float32(car.Price),
		}
		cars = append(cars, pbCar)
	}

	return &pb.GetCarsResponse{
		Cars: cars,
	}, nil
}

func (s *CarServiceServerImpl) GetCar(ctx context.Context, req *pb.GetCarRequest) (*pb.Car, error) {
	// Extract car ID from the request
	id := req.GetId()

	// Query the database to find the car with the specified ID
	var dbCar models.Car
	if err := database.DB.First(&dbCar, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "Car not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to fetch car from database: %v", err)
	}

	// Convert the database car to protobuf car
	pbCar := &pb.Car{
		Id:      int32(dbCar.ID),
		Make:    dbCar.Make,
		Model:   dbCar.Model,
		Year:    int32(dbCar.Year),
		Color:   dbCar.Color,
		Mileage: int32(dbCar.Mileage),
		Price:   float32(dbCar.Price),
	}

	// Return the protobuf car
	return pbCar, nil
}

func (s *CarServiceServerImpl) AddCar(ctx context.Context, req *pb.Car) (*pb.Car, error) {
	// Extract car data from the request
	make := req.GetMake()
	model := req.GetModel()
	year := req.GetYear()
	color := req.GetColor()
	mileage := req.GetMileage()
	price := req.GetPrice()

	// Create a new car model object with the extracted data
	newCar := models.Car{
		Make:    make,
		Model:   model,
		Year:    int(year),
		Color:   color,
		Mileage: int(mileage),
		Price:   float64(price),
	}

	// Save the new car to the database
	if err := database.DB.Create(&newCar).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create car: %v", err)
	}

	// Convert the created car to pb.Car message
	// Return the pb.Car message
	return &pb.Car{
		Id:      int32(newCar.ID),
		Make:    newCar.Make,
		Model:   newCar.Model,
		Year:    int32(newCar.Year),
		Color:   newCar.Color,
		Mileage: int32(newCar.Mileage),
		Price:   float32(newCar.Price),
	}, nil
}

// UpdateCar implements the UpdateCar RPC method.
func (s *CarServiceServerImpl) UpdateCar(ctx context.Context, req *pb.UpdateCarRequest) (*pb.Car, error) {
	// Extract car ID and updated car data from the request
	id := req.GetId()
	updatedCar := req.GetCar()

	// Find the car in the database based on the ID
	var existingCar models.Car
	if err := database.DB.First(&existingCar, id).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "Car not found")
	}

	// Update the existing car's fields with the updated data
	existingCar.Make = updatedCar.Make
	existingCar.Model = updatedCar.Model
	existingCar.Year = int(updatedCar.Year)
	existingCar.Color = updatedCar.Color
	existingCar.Mileage = int(updatedCar.Mileage)
	existingCar.Price = float64(updatedCar.Price)

	// Save the updated car to the database
	database.DB.Save(&existingCar)

	// Convert the updated car to pb.Car message
	// Return the pb.Car message
	return &pb.Car{
		Id:      int32(existingCar.ID),
		Make:    existingCar.Make,
		Model:   existingCar.Model,
		Year:    int32(existingCar.Year),
		Color:   existingCar.Color,
		Mileage: int32(existingCar.Mileage),
		Price:   float32(existingCar.Price),
	}, nil
}

// DeleteCar implements the DeleteCar RPC method.
func (s *CarServiceServerImpl) DeleteCar(ctx context.Context, req *pb.DeleteCarRequest) (*empty.Empty, error) {
	// Extract car ID from the request
	id := req.GetId()

	// Delete the car from the database based on the ID
	if err := database.DB.Delete(&models.Car{}, id).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "Car not found")
	}

	// Return an empty response
	return &empty.Empty{}, nil
}

func main() {
	if err := database.ConnectGRPC(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterCarServiceServer(server, &CarServiceServerImpl{})

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	log.Println("Starting gRPC server on port 50052...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
