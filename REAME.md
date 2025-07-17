# Go API Project Template

Шаблон проекта для создания REST API на Go с использованием Gin, GORM, PostgreSQL и Docker.

## Технический стек

- **Go 1.21** - основной язык программирования
- **Gin** - веб-фреймворк для создания API
- **GORM** - ORM для работы с базой данных
- **PostgreSQL** - реляционная база данных
- **Docker & Docker Compose** - контейнеризация

## Структура проекта

```
go-app/
├── config/
│   └── database.go          # Конфигурация базы данных
├── handlers/
│   └── user.go              # Обработчики для пользователей
├── middleware/
│   ├── cors.go              # CORS middleware
│   ├── logger.go            # Логирование
│   └── error.go             # Обработка ошибок
├── models/
│   ├── user.go              # Модель пользователя
│   └── response.go          # Модели ответов
├── routes/
│   └── routes.go            # Настройка маршрутов
├── main.go                  # Точка входа
├── docker-compose.yml       # Docker Compose конфигурация
├── Dockerfile              # Docker образ
├── .env                    # Переменные окружения
├── .env.example            # Пример переменных окружения
├── go.mod                  # Go модули
├── go.sum                  # Контрольные суммы модулей
├── Makefile               # Команды для разработки
├── .air.toml              # Конфигурация hot reload
└── README.md              # Этот файл
```

## Быстрый старт

### 1. Клонирование и настройка

```bash
# Создание директории проекта
mkdir go-app
cd go-app

# Копирование файлов из шаблона
# (скопируйте все файлы из артефактов выше)

# Настройка переменных окружения
cp .env.example .env
```

### 2. Запуск с Docker Compose (рекомендуется)

```bash
# Запуск всех сервисов
make docker-up

# Или напрямую
docker-compose up -d
```

### 3. Локальный запуск

```bash
# Установка зависимостей
go mod tidy

# Запуск PostgreSQL в Docker
docker run -d \
  --name postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=testdb \
  -p 5432:5432 \
  postgres:15-alpine

# Запуск приложения
make run
```

## API Endpoints

### Пользователи

- `GET /api/v1/users` - получить всех пользователей
- `GET /api/v1/users/:id` - получить пользователя по ID
- `POST /api/v1/users` - создать нового пользователя
- `PUT /api/v1/users/:id` - обновить пользователя
- `DELETE /api/v1/users/:id` - удалить пользователя

### Health Check

- `GET /health` - проверка состояния сервера

## Примеры использования

### Создание пользователя

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30
  }'
```

### Получение всех пользователей

```bash
curl http://localhost:8080/api/v1/users
```

### Получение пользователя по ID

```bash
curl http://localhost:8080/api/v1/users/1
```

### Обновление пользователя

```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "jane@example.com",
    "age": 25
  }'
```

### Удаление пользователя

```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## Доступные команды

```bash
# Разработка
make run                 # Запуск приложения
make build              # Сборка приложения
make test               # Запуск тестов
make clean              # Очистка артефактов
make fmt                # Форматирование кода
make deps               # Установка зависимостей

# Docker
make docker-up          # Запуск Docker Compose
make docker-down        # Остановка Docker Compose
make docker-rebuild     # Пересборка и запуск
make logs               # Просмотр логов

# База данных
make db-connect         # Подключение к PostgreSQL

# Разработка с hot reload
make dev-deps           # Установка зависимостей для разработки
make dev                # Запуск с hot reload
```

## Переменные окружения

```env
DB_HOST=localhost       # Хост базы данных
DB_USER=postgres        # Пользователь базы данных
DB_PASSWORD=password    # Пароль базы данных
DB_NAME=testdb          # Имя базы данных
DB_PORT=5432           # Порт базы данных
PORT=8080              # Порт приложения
```

## Структура ответов API

### Успешный ответ

```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30,
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
}
```

### Ответ с ошибкой

```json
{
  "success": false,
  "message": "Operation failed",
  "error": "Detailed error message"
}
```

## Расширение проекта

### Добавление новой модели

1. Создайте новую модель в `models/`
2. Добавьте миграцию в `models/migrate.go`
3. Создайте обработчики в `handlers/`
4. Добавьте маршруты в `routes/routes.go`

### Добавление middleware

1. Создайте новый файл в `middleware/`
2. Добавьте middleware в `main.go`

### Добавление аутентификации

Для добавления JWT аутентификации:

```bash
go get github.com/golang-jwt/jwt/v5
```

Затем создайте middleware для проверки токенов.

## Тестирование

```bash
# Запуск всех тестов
make test

# Запуск тестов с покрытием
go test -v -cover ./...

# Запуск конкретного теста
go test -v ./handlers -run TestUserHandler
```

## Производство

### Сборка для production

```bash
# Сборка оптимизированного бинарного файла
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -s' -o main .

# Или использование Docker
docker build -t go-app .
```

### Переменные окружения для production

```env
GIN_MODE=release
DB_HOST=your-production-db-host
DB_SSL_MODE=require
```

## Полезные ресурсы

- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Docker Documentation](https://docs.docker.com/)

## Лицензия

MIT License