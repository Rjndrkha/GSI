package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rjndrkha/gsitest/pkg/util"
)

func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"status": 401, "error": true, "message": "Missing token"})
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return util.JWT_SECRET, nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"status": 401, "error": true, "message": "Invalid or expired token"})
	}

	claims := token.Claims.(jwt.MapClaims)

	c.Locals("user_id", claims["user_id"])

	return c.Next()
}
