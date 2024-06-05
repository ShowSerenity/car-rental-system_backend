package main

import (
	"context"
	"log"
	"net"

	"car-rental-system/auth-service/database"
	"car-rental-system/auth-service/models"
	pb "car-rental-system/auth-service/proto"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServiceServerImpl implements the AuthServiceServer interface.
type AuthServiceServerImpl struct {
	pb.UnimplementedAuthServiceServer
}

// Register implements the Register RPC method.
func (s *AuthServiceServerImpl) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	username := req.GetUsername()
	email := req.GetEmail()
	password := req.GetPassword()
	fullName := req.GetFullName()

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash password: %v", err)
	}

	// Create a new user
	user := models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		FullName: fullName,
	}

	// Save the user to the database
	if err := database.DB.Create(&user).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create user: %v", err)
	}

	// Return success response
	return &pb.RegisterResponse{
		Status:  200,
		Message: "User registered successfully",
	}, nil
}

// Login implements the Login RPC method.
func (s *AuthServiceServerImpl) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()

	// Find the user by username
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid username or password")
	}

	// Compare the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid username or password")
	}

	// Return success response
	return &pb.LoginResponse{
		Status:  200,
		Message: "Login successful",
	}, nil
}

func main() {
	if err := database.ConnectGRPC(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, &AuthServiceServerImpl{})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("Starting gRPC server on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
