package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const SecretKey = "secret"

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Query("token")

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Token missing",
			})
		}

		claims, err := parseJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid Token",
			})
		}

		fmt.Println("Token Claims:", claims)

		userRole, ok := claims["role"].(string)
		if !ok || userRole != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden: Insufficient Permissions",
			})
		}
		c.Locals("userClaims", claims)
		return c.Next()
	}
}

func parseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}

func RoleMiddleware(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userClaims, ok := c.Locals("userClaims").(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
		userRole, ok := userClaims["role"].(string)
		if !ok || userRole != role {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden: Insufficient Permissions",
			})
		}

		return c.Next()
	}
}
