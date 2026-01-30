# Run the application
run:
	@go run cmd/api/main.go
# Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v
# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

# Create Migration
install-migration:
	@curl -fsSL https://packagecloud.io/golang-migrate/migrate/gpgkey | sudo gpg --dearmor -o /etc/apt/keyrings/migrate.gpg && echo "deb [signed-by=/etc/apt/keyrings/migrate.gpg] https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list && apt-get update && apt-get install -y migrate

create-dump:
	DB_USER=${DB_USER} DB_NAME=${DB_NAME} sh ./internal/database/scripts/create_dump.sh

apply-dump:
	DB_USER=${DB_USER} DB_NAME=${DB_NAME} sh ./internal/database/scripts/apply_dump.sh

migration-down: create-dump
	migrate -path ./internal/database/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose \
	down

create-migration:
	@read -p "Please provide migration name: " name && \
		echo $$name && \
		migrate create -ext sql -dir ./internal/database/migrations/ -seq $$name

.PHONY: run test clean watch docker-run docker-down itest templ-install create-dump create-migration migration-down apply-dump install-migration
