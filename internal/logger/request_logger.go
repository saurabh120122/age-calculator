package logger

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestLogger(log *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		reqID, _ := c.Locals("requestId").(string)

		log.Info("request",
			zap.String("requestId", reqID),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Duration("duration", duration),
		)

		return err
	}
}
