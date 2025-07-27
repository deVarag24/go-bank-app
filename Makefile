migrate-up:
	migrate -path app/db/migration -database "postgres://postgres:postgres@localhost:5432/go_bank_app_db?sslmode=disable" -verbose up

migrate-down:
	migrate -path app/db/migration -database "postgres://postgres:postgres@localhost:5432/go_bank_app_db?sslmode=disable" -verbose down

.PHONY: migrate-up migrate-down