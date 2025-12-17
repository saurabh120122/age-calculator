package logger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := uuid.New().String()
		c.Set("X-Request-ID", id)
		c.Locals("requestId", id)
		return c.Next()
	}
}
