IAMAGE_NAME=takumi2786/zero-backend:latest

docker/build:
	docker-compose build

docker/run:
	docker-compose up -d

go/build:
	wire ./cmd
	go build -o zero_api ./cmd

go/run:
	go run ./cmd/main.go ./cmd/wire_gen.go

mysql:
	docker-compose exec mysql mysql zero_system
