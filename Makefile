DB_URL=postgres://user:pass@localhost:5432/mydb?sslmode=disable
MIGRATIONS_DIR=./internal/infra/db/migrations

.PHONY: migrate-up migrate-down migrate-force migrate-new dev

# Run all up migrations
migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

# Roll back one migration
migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

# Force set version (careful!)
migrate-force:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force 1

# Create a new migration file (name from NAME= argument)
migrate-new:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name
dev:
	go run cmd/web/main.go