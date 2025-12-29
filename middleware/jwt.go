package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/indrariksa/contactsAPI/module"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "missing Authorization header",
			})
		}

		// Format: "Bearer <token>"
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid Authorization format. use: Bearer <token>",
			})
		}

		claims, err := module.ParseAndValidateJWT(parts[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid/expired token",
				"error":   err.Error(),
			})
		}

		// Simpan ke locals biar bisa dipakai di handler
		c.Locals("user_id", claims.UserID)
		c.Locals("username", claims.Username)

		return c.Next()
	}
}
