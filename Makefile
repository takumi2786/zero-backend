IAMAGE_NAME=takumi2786/zero-backend:latest

docker/build:
	docker-compose build

docker/run:
	docker-compose up -d

build:
	go build -o zero-api ./cmd/zero-api

run:
	set -a && source ./deploy/local/.env && go run ./cmd/main.go
