package handlers

import (
	"tasoryx/app"

	"github.com/gofiber/fiber/v3"
)

var jwtSecret []byte

func Login(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		jwtSecret = []byte(appContainer.Config().JwtSecret)
		// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
		//TODO: add user id to the token
		return nil
	}
}
