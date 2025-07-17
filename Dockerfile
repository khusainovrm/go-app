FROM golang:1.21-alpine AS builder

WORKDIR /app

# Установка зависимостей
RUN apk add --no-cache git

# Копирование файлов go.mod и go.sum
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копирование исходного кода
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копирование бинарного файла из builder
COPY --from=builder /app/main .

# Копирование .env файла (если есть)
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]