include .env
MIGRATIONS_PATH = ./cmd/migrate/migrations

migration:
	~/go/bin/migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	~/go/bin/migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

migrate-down:
	~/go/bin/migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

migrate-drop:
	~/go/bin/migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) drop

seed:
	@go run ./cmd/migrate/seed/main.go
