# Go Backend Task - User Management API

A production-style REST API built using Go and GoFiber to manage users with their date of birth and dynamically calculated age.

## Features

* CRUD operations for users
* Dynamic age calculation using Go's `time` package
* PostgreSQL database
* SQLC for type-safe database access
* Uber Zap structured logging
* Request ID middleware
* Request duration logging middleware
* Input validation using `go-playground/validator`
* Docker and Docker Compose support
* Layered architecture (Handler → Service → Repository)

---

## Tech Stack

* Go
* GoFiber
* PostgreSQL
* SQLC
* pgx/v5
* Uber Zap
* go-playground/validator
* Docker & Docker Compose

---

## Project Structure

```text
.
├── cmd/server/
├── config/
├── db/
│   ├── migrations/
│   └── sqlc/
│       ├── generated/
│       └── queries/
├── internal/
│   ├── handler/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── routes/
│   ├── service/
│   ├── logger/
│   └── utils/
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

---

## API Endpoints

### Create User

```http
POST /api/v1/users
```

Request:

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

Response:

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}
```

---

### Get User by ID

```http
GET /api/v1/users/:id
```

Response:

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 36
}
```

---

### List Users

```http
GET /api/v1/users
```

---

### Update User

```http
PUT /api/v1/users/:id
```

---

### Delete User

```http
DELETE /api/v1/users/:id
```

Returns `204 No Content`.

---

## Running Locally

### Prerequisites

* Go 1.25+
* Docker Desktop

### Start PostgreSQL

```bash
docker compose up -d postgres
```

### Run migrations

Apply migration files from:

```text
db/migrations/
```

### Generate SQLC code

```bash
sqlc generate
```

### Start the server

```bash
go run ./cmd/server
```

Server runs on:

```text
http://localhost:3000
```

---

## Running with Docker

```bash
docker compose up --build
```

---

## Environment Variables

```env
APP_PORT=3000

DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=userdb
```

---

## Dynamic Age Calculation

Age is not stored in the database. It is calculated dynamically during API responses using Go's `time` package.

---

## Logging

* Structured logging using Uber Zap
* Request ID middleware
* Request duration logging
* HTTP request tracing

```
```
