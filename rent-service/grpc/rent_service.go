package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	"car-rental-system/rent-service/database"
	"car-rental-system/rent-service/models"
	rentpb "car-rental-system/rent-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RentServiceServerImpl struct {
	rentpb.UnimplementedRentServiceServer
}

func (s *RentServiceServerImpl) RentCar(ctx context.Context, req *rentpb.RentRequest) (*rentpb.RentResponse, error) {
	rental := models.Rental{
		UserID:    int(req.GetUserId()),
		CarID:     int(req.GetCarId()),
		StartDate: req.GetStartDate().AsTime(),
		EndDate:   req.GetEndDate().AsTime(),
	}

	if err := database.DB.Create(&rental).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to rent car: %v", err)
	}

	return &rentpb.RentResponse{
		Rental: &rentpb.Rental{
			Id:        int32(rental.ID),
			UserId:    int32(rental.UserID),
			CarId:     int32(rental.CarID),
			StartDate: timestamppb.New(rental.StartDate),
			EndDate:   timestamppb.New(rental.EndDate),
		},
	}, nil
}

func (s *RentServiceServerImpl) GetRentalHistory(ctx context.Context, req *rentpb.RentalHistoryRequest) (*rentpb.RentalHistoryResponse, error) {
	userID := req.GetUserId()
	sortBy := req.GetSortBy()
	page := int(req.GetPage())
	limit := int(req.GetLimit())

	var rentals []models.Rental
	db := database.DB

	if userID != 0 {
		db = db.Where("user_id = ?", userID)
	}

	if sortBy != "" {
		db = db.Order(sortBy)
	}

	if page != 0 && limit != 0 {
		offset := (page - 1) * limit
		db = db.Offset(offset).Limit(limit)
	}

	if err := db.Find(&rentals).Error; err != nil {
		return nil, fmt.Errorf("unable to retrieve rental history: %v", err)
	}

	var rentalPB []*rentpb.Rental
	for _, rental := range rentals {
		rentalPB = append(rentalPB, &rentpb.Rental{
			Id:        int32(rental.ID),
			UserId:    int32(rental.UserID),
			CarId:     int32(rental.CarID),
			StartDate: timestamppb.New(rental.StartDate),
			EndDate:   timestamppb.New(rental.EndDate),
		})
	}

	return &rentpb.RentalHistoryResponse{Rentals: rentalPB}, nil
}

func main() {
	if err := database.ConnectGRPC(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	server := grpc.NewServer()
	rentpb.RegisterRentServiceServer(server, &RentServiceServerImpl{})

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen on port 50053: %v", err)
	}

	fmt.Println("Starting gRPC server on port 50053...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
