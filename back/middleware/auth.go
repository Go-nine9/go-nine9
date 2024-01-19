package middleware

import (
	"github.com/go-nine9/go-nine9/services"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("jwt")

		if token == "" {
			return c.Next()
		}

		claims, err := services.ParseJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid JWT",
			})
		}

		c.Locals("userRoles", claims["role"])

		return c.Next()
	}
}
