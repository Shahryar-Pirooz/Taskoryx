package middlewares

import (
	"errors"
	"strings"
	"tasoryx/api/http/handlers"
	"tasoryx/app"
	"tasoryx/pkg/jwt"

	"github.com/gofiber/fiber/v3"
)

func RequireAuth(appContainer app.App, minRole int8) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			return handlers.HandleError(errors.New("missing or invalid Authorization header"), c, fiber.StatusUnauthorized)
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ValidationToken(tokenStr, appContainer.Config().Jwt.AccessKey)
		if err != nil {
			return handlers.HandleError(err, c, fiber.StatusUnauthorized)
		}
		if claims.Role < minRole {
			return handlers.HandleError(errors.New("insufficient permissions"), c, fiber.StatusForbidden)
		}
		c.Locals("user_id", claims.UserID)
		return c.Next()
	}
}
