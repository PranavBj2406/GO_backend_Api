package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/go-backend-task/internal/logger"
	"go.uber.org/zap"
)

// RequestLogger logs one structured entry per request.
func RequestLogger(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)

	requestID, _ := c.Locals(RequestIDKey).(string)
	if requestID == "" {
		requestID = c.Get(RequestIDHeader)
	}

	logger.GetLogger().Info("http request",
		zap.String("request_id", requestID),
		zap.String("method", c.Method()),
		zap.String("path", c.OriginalURL()),
		zap.Int("status_code", c.Response().StatusCode()),
		zap.String("duration", duration.String()),
		zap.String("client_ip", c.IP()),
	)

	return err
}
