# --- Сборка приложения ---
FROM golang:1.24.3-alpine AS builder

RUN apk add --no-cache bash

WORKDIR /app

COPY ../ ./

# Загружаем зависимости и билдим
RUN go mod download
RUN go build -o app .

# --- Финальный образ ---
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app ./
COPY --from=builder /app/config/configd.yaml ./config/configd.yaml
COPY --from=builder /app/migrations/ ./migrations/
COPY --from=builder /app/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["sh", "./entrypoint.sh"]


