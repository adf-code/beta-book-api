# 📚 Beta Book API

A clean and modular Golang project for managing books using Clean Architecture principles. This project demonstrates best practices for HTTP API development, PostgreSQL integration, environment configuration, and migration handling.

---

## 🗂 Project Structure

```
github.com/adf-code/beta-book-api/
├── cmd/
│   ├── main.go                         # Application entry point
│   └── migrate.go                      # CLI for running migrations
├── config/
│   ├── config.go                       # Loads environment variables into a config struct
│   └── db_postgres.go                  # PostgreSQL database connection initialization
├── internal/
│   ├── delivery/
│   │   └── http/
│   │       └── book_handler.go         # HTTP handlers for book entity
│   ├── entity/
│   │   └── book.go                     # Book domain model/entity
│   ├── migration/
│   │   ├── runner.go                   # Core migration logic
│   │   └── utils.go                    # Helper functions for migration (file parsing, versioning, etc.)
│   ├── repository/
│   │   ├── book_repository.go          # Book repository interface
│   │   └── book_repository_postgres.go # PostgreSQL implementation of book repository
│   ├── usecase/
│   │   └── book_usecase.go             # Use cases for managing book entities
├── migration/
│   ├── {timestamp}_{action}_up.sql     # Timestamped UP migration scripts
│   └── {timestamp}_{action}_down.sql   # Timestamped DOWN migration scripts
├── .env.example                        # Example environment file for setup
├── .gitignore                          # Git ignore rules for files/folders
├── go.mod                              # Go module configuration file
└── README.md                           # Project documentation
```

---

## 🧼 Clean Architecture Overview

### `internal/entity/`
Defines the core business entities such as `Book`. These are simple structs and are completely independent of other layers.

### `internal/usecase/`
Contains application logic (use cases) such as `GetAllBooks`, `CreateBook`, etc. Use cases operate only on defined entities and do not depend on frameworks or external systems.

### `internal/repository/`
Defines repository interfaces. They describe how the application expects to fetch or store data but do not contain actual database logic.

### `internal/repository/book_repository_postgres.go`
Concrete implementation of the repository interface using PostgreSQL and Go's `database/sql` package.

### `internal/delivery/http/`
Implements HTTP handlers that receive requests, validate input, call use cases, and return responses.

### `internal/delivery/response/api_response.go`
Provides a consistent, reusable JSON response format for all API endpoints. The standard structure looks like this:
```json
{
  "status": "SUCCESS | FAILED",
  "entity": "books",
  "state": "getAllBooks",
  "message": "Success Get All Books",
  "data": []
}
```
This improves API consistency and simplifies client-side integration.

---

## ⚙️ Configuration

### `config/config.go`
Loads application configuration from environment variables, typically using a `.env` file.

### `config/db_postgres.go`
Initializes PostgreSQL connection using the values from configuration.

---

## 🚀 Application Entry Point

### `cmd/main.go`
Sets up the HTTP server, loads environment variables, connects to the database, injects dependencies, and starts the application.

---

## 🛠️ Migration System

### `cmd/migrate.go`
CLI entry point to run database migrations:

```bash
go run cmd/migrate.go up     # Run all pending migrations
go run cmd/migrate.go down   # Roll back the last migration
```

### `internal/migration/`
Contains core migration logic (`runner.go`) and utility functions (`utils.go`) such as version parsing and SQL execution.

### `migration/`
Holds raw SQL files for migrations:
- `20250725100000_create_books_table.up.sql`
- `20250725100000_create_books_table.down.sql`

---

## 🔐 Environment Variables

### `.env`
Environment configuration file. **Should not be committed**.

### `.env.example`
Example file with placeholder values. This should be committed to help other developers set up the project.

---

## 🔒 Git Configuration

### `.gitignore`
Ignores unnecessary files such as:
- Build artifacts
- Environment files
- IDE/editor settings
- Logs and database dumps

---

## 📦 Go Modules

### `go.mod`
Declares the module path and manages external dependencies for reproducible builds.

### `go.sum`
Records the cryptographic checksums for dependencies.

---

## 🧪 How to Run

1. Copy the example config:
```bash
cp .env.example .env
```

2. Fill in your PostgreSQL credentials in `.env`

3. Run database migrations:
```bash
go run cmd/migrate.go up
```

4. Run the web server:
```bash
go run cmd/main.go
```
---

## ✅ Output Format (Standard API Response)

All HTTP responses follow this structure:
```json
{
  "status": "success" | "failed",
  "entity": "books",
  "state": "getAllBooks",
  "message": "Success Get All Books",
  "data": []
}
```
## 📦 Standard API Response Format

This document explains the standard JSON response structure used in the Beta Book API project, following Clean Architecture principles.

## ✅ Example

```json
{
  "status": "success" | "failed",
  "entity": "books",
  "state": "getAllBooks",
  "message": "Success Get All Books",
  "data": []
}
```

---

## 🟢 `status: "success" | "failed"`

- **Description**: Represents the outcome of the HTTP request.
- **Values**:
    - `"success"`: The request was processed successfully (HTTP 2xx).
    - `"failed"`: The request failed due to client or server error (HTTP 4xx/5xx).
- **Purpose**: Allows the frontend to easily detect success or failure and handle user feedback accordingly.

---

## 🟠 `entity: "books"`

- **Description**: Indicates the entity or resource being processed.
- **Example Values**:
    - `"books"` — book resource
    - `"users"` — user resource
- **Relation to Clean Architecture**: Refers to the domain object defined in `internal/entity/`.

---

## 🔵 `state: "getAllBooks"`

- **Description**: Represents the use case that was executed.
- **Example Values**:
    - `"getAllBooks"` — fetch all books
    - `"createBook"` — create a new book
- **Relation to Clean Architecture**: Maps to the business logic function in `internal/usecase/`.

---

## 🟣 `message: "Success Get All Books"`

- **Description**: Human-readable message summarizing the result.
- **Purpose**: Shown on the client side as a notification or log.
- **Best Practice**: Keep it short, clear, and user-friendly.

---

## 🟤 `data: []`

- **Description**: Contains the actual result data of the request.
- **Type**:
    - `[]`: for list responses
    - `{}`: for single object
- **Special Rule**: Always return an empty array `[]` if no data exists, **never null** — this helps avoid null checks in frontend logic.

---

## 🧭 Summary

This response format helps ensure consistency across all API endpoints, improves developer experience, and facilitates frontend-backend integration.

---

