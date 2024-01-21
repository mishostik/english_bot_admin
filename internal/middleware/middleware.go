package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Get("Access")
		if isValidAccessToken(accessToken) {
			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}
}

func isValidAccessToken(access string) bool {
	return true
}
