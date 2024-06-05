.PHONY: all auth-service cars-service rent-service test build docker-up docker-down

all: build docker-up

build: auth-service cars-service rent-service

auth-service:
	cd auth-service && go build -o auth-service

cars-service:
	cd cars-service && go build -o cars-service

rent-service:
	cd rent-service && go build -o rent-service

test:
	cd auth-service && go test ./...
	cd cars-service && go test ./...
	cd rent-service && go test ./...

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down
