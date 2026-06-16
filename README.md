# Go Backend Task

A REST API built with Go and GoFiber.

## Project Structure

```
.
├── cmd/server/           # Application entry point
├── config/               # Configuration management
├── db/                   # Database-related files
│   ├── migrations/       # Database migrations
│   └── sqlc/             # SQLC generated code
├── internal/             # Private application code
│   ├── handler/          # HTTP request handlers
│   ├── repository/       # Data access layer
│   ├── service/          # Business logic layer
│   ├── routes/           # Route definitions
│   ├── middleware/       # Middleware functions
│   ├── models/           # Data models
│   ├── logger/           # Logging utilities
│   └── utils/            # Utility functions
└── go.mod               # Go module definition
```

## Getting Started

1. Update the `go.mod` file with your module name
2. Install dependencies: `go mod download`
3. Implement your handlers, services, and routes
4. Run: `go run ./cmd/server`

## Dependencies

Add GoFiber and other required dependencies:

```bash
go get -u github.com/gofiber/fiber/v2
```
