package helper

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = os.Getenv("JWT_SECRET")

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Vérifiez que le token est signé avec l'algorithme attendu
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
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

func GetToken(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")

	// Split the header into two parts: "Bearer" and the token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		// Handle error: invalid or missing Authorization header
		return "", c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid or missing Authorization header",
		})
	}

	token := parts[1]

	return token, nil
}

func GetClaims(c *fiber.Ctx) (jwt.MapClaims, error) {
	claims, ok := c.Locals("userClaims").(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Unauthorized: Claims missing or not jwt.MapClaims")
	}
	return claims, nil
}

func GeneratePassword(length int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+{}[]|;':\",.<>/?`~")

	password := make([]rune, length)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}

	for _, c := range password {
		if !unicode.IsUpper(c) && !unicode.IsLower(c) && !unicode.IsDigit(c) && c < '!' || c > '~' {
			return GeneratePassword(length)
		}
	}

	return string(password)
}

func isManager(user_role string) bool {
	if user_role == "manager" {
		return true
	}
	return false
}

func isOwner(user_role string) bool {
	if user_role == "owner" {
		return true
	}
	return false
}

// func RoleMiddleware(c *fiber.Ctx) error {
// 	authHeader := c.Get("Authorization")
// 	if authHeader == "" {
// 		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
// 	}

// 	tokenString := strings.Split(authHeader, " ")[1]
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecret, nil
// 	})

// 	if err != nil || !token.Valid {
// 		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
// 	}

// 	claims := token.Claims.(jwt.MapClaims)
// 	userRole := claims["user_role"].(string)

// 	if userRole != "admin" {
// 		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
// 	}

// 	return c.Next()
// }
