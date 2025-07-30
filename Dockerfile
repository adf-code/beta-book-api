# syntax=docker/dockerfile:1
FROM golang:1.23-alpine as builder

WORKDIR /app

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Generate Swagger docs
RUN swag init -g cmd/main.go -o ./docs

# Build main binary
RUN go build -o /bin/beta-book-api ./cmd/main.go

# Build migrate binary
RUN go build -o /bin/migrate ./cmd/migrate.go

# Final minimal image
FROM alpine:latest

# Install bash & curl for wait-for-it
RUN apk add --no-cache bash curl

WORKDIR /app

# Copy env file
COPY .env.docker .env

# Copy migration files
COPY migration ./migration

# Copy built binaries
COPY --from=builder /bin/beta-book-api .
COPY --from=builder /bin/migrate .

# Copy Swagger docs
COPY --from=builder /app/docs ./docs

# Copy wait-for-it script
COPY scripts/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

EXPOSE 8080

# Run: wait for postgres and minio, then migrate, then start API
CMD sh -c "/wait-for-it.sh postgres:5432 -- /wait-for-it.sh minio:9000 -- ./migrate up && ./beta-book-api"
