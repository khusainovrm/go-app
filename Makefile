.PHONY: run build test clean docker-up docker-down migrate

# Запуск приложения
run:
	go run main.go

# Сборка приложения
build:
	go build -o bin/main main.go

# Запуск тестов
test:
	go test -v ./...

# Очистка артефактов сборки
clean:
	rm -rf bin/

# Запуск Docker Compose
docker-up:
	docker-compose up -d

# Остановка Docker Compose
docker-down:
	docker-compose down

# Пересборка и запуск Docker Compose
docker-rebuild:
	docker-compose down
	docker-compose up --build -d

# Просмотр логов
logs:
	docker-compose logs -f

# Подключение к базе данных
db-connect:
	docker-compose exec postgres psql -U postgres -d testdb

# Инициализация зависимостей
deps:
	go mod tidy
	go mod download

# Форматирование кода
fmt:
	go fmt ./...

# Линтинг кода
lint:
	golangci-lint run

# Установка зависимостей для разработки
dev-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/cosmtrek/air@latest

# Запуск с hot reload
dev:
	air