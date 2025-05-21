package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
)

func AuthMiddleware(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ContextKey: "jwt",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Return status 401 and failed authentication error.
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("jwt").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			// Extract user ID from claims["sub"]
			userIdStr := claims["sub"].(string)
			userId, err := strconv.ParseUint(userIdStr, 10, 64)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": true,
					"msg":   "Invalid token subject",
				})
			}
			c.Locals("user_id", uint(userId))
			return c.Next()
		},
	})(c)
}
