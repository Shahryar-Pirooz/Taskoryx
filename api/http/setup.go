package http

import (
	"fmt"
	"tasoryx/app"
	"tasoryx/config"

	"github.com/gofiber/fiber/v3"
)

type HTTPService struct {
	app *app.App
}

func Run(cfg config.ServerConfig) error {
	app := fiber.New(fiber.Config{AppName: "Taskoryx"})
	return app.Listen(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
}
