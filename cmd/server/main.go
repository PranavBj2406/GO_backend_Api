package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// TODO: Initialize application
	_ = fiber.New()
	_ = cors.New()
	_ = godotenv.Load()
	_ = validator.New()
	_ = pgxpool.Config{}
	_ = zap.NewNop()
	_ = uuid.New()
}
