APP_NAME=todo_api
DB_NAME=todo_db
ENV_FILE ?= .env

run: build up

build:
	@echo "Building the Docker image for the API..."
	docker compose build --build-arg ENV_FILE=${ENV_FILE}

up:
	@echo "Starting the Docker containers..."
	docker compose up -d

# migrate:
# 	@echo "Waiting for the database to be ready..."
# 	@docker compose exec -T $(DB_NAME) sh -c 'until pg_isready -q; do echo "Waiting for database..."; sleep 2; done;'
# 	@echo "Running database migrations..."
# 	@docker compose exec -T $(APP_NAME) /todo-api migrate # Alterar este comando conforme suas migrations

down:
	@echo "Stopping and removing Docker containers, networks and volumes..."
	docker compose down --volumes

restart: down run

logs:
	@echo "Showing logs from Docker containers..."
	docker compose logs -f

stop:
	@echo "Stopping all Docker containers..."
	docker compose stop

clean:
	@echo "Cleaning up unused Docker containers, networks and images..."
	docker system prune -f


# Testes
test:
	@echo "Running tests..."
	go test -v ./...

test-coverage:
	@echo "Generating test coverage report..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
