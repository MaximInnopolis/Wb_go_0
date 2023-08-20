include .env

build:
	docker-compose build wbl0-app

run:
	docker-compose up

run_sender:
	go run ./cmd/main.go

migrate_up:
	migrate -path ./migration -database "postgresql://${DBUSERNAME}:${DBPASSWORD}@localhost:5432/${DBNAME}?sslmode=disable" -verbose up

migrate_down:
	migrate -path ./migration -database "postgresql://${DBUSERNAME}:${DBPASSWORD}@localhost:5432/${DBNAME}?sslmode=disable" -verbose down