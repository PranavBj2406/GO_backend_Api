package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const RequestIDKey = "request_id"
const RequestIDHeader = "X-Request-ID"

// RequestID assigns a request ID to each incoming request and propagates it in the response.
func RequestID(c *fiber.Ctx) error {
	requestID := c.Get(RequestIDHeader)
	if strings.TrimSpace(requestID) == "" {
		requestID = uuid.NewString()
	}

	c.Locals(RequestIDKey, requestID)
	c.Set(RequestIDHeader, requestID)

	return c.Next()
}
