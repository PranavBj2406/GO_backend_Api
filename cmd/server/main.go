package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yourusername/go-backend-task/config"
	"github.com/yourusername/go-backend-task/db/sqlc/generated"
	"github.com/yourusername/go-backend-task/internal/handler"
	"github.com/yourusername/go-backend-task/internal/logger"
	"github.com/yourusername/go-backend-task/internal/repository"
	"github.com/yourusername/go-backend-task/internal/routes"
	"github.com/yourusername/go-backend-task/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	log := logger.GetLogger()
	defer func() {
		_ = log.Sync()
	}()

	db, err := config.NewDB(*cfg)
	if err != nil {
		log.Fatal("failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	queries := generated.New(db)
	repo := repository.NewUserRepository(queries)
	svc := service.NewUserService(repo)
	handler := handler.NewUserHandler(svc)

	app := fiber.New()
	routes.SetupRoutes(app, handler)

	addr := fmt.Sprintf(":%d", cfg.AppPort)
	log.Info("server starting", zap.String("addr", addr))

	if err := app.Listen(addr); err != nil {
		log.Fatal("failed to start server", zap.Error(err))
	}
}
