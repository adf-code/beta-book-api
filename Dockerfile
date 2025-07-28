# syntax=docker/dockerfile:1
FROM golang:1.23-alpine as builder

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Generate Swagger docs
RUN swag init -g cmd/main.go -o ./docs

# Build main binary
RUN go build -o /bin/beta-book-api ./cmd/main.go

# 🔥 Build migrate binary
RUN go build -o /bin/migrate ./cmd/migrate.go

# Final minimal image
FROM alpine:latest

WORKDIR /app

# Copy env
COPY .env.docker .env

# Copy migration files
COPY migration ./migration

# Copy built binaries
COPY --from=builder /bin/beta-book-api .
COPY --from=builder /bin/migrate .

# Copy docs if needed
COPY --from=builder /app/docs ./docs

EXPOSE 8080

# Run migration, then app
CMD sh -c "./migrate up && ./beta-book-api"
