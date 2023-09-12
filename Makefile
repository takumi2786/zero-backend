IAMAGE_NAME=takumi2786/zero-backend:latest

docker/build:
	docker-compose build

docker/run:
	docker-compose up -d

go/build:
	go build -o zero-api ./cmd/zero-api

go/run:
	set -a && source ./deploy/local/.env && go run ./cmd/main.go

mysql:
	docker-compose exec mysql mysql zero_system
