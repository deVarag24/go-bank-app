DB_URL=postgres://postgres:postgres@localhost:5432/go_bank_app_db?sslmode=disable
MIGRATION_PATH=app/db/migration
MIGRATE=migrate -path ${MIGRATION_PATH} -database $(DB_URL)

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

migrate-new:
	migrate create -ext sql -dir ${MIGRATION_PATH} -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrate-up migrate-down sqlc migrate-new test