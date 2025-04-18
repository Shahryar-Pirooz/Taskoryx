package http

import (
	"fmt"
	"tasoryx/api/http/handlers"
	"tasoryx/app"
	"tasoryx/config"

	"github.com/gofiber/fiber/v3"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	app := fiber.New(fiber.Config{AppName: "Taskoryx"})
	api := app.Group("/api/v1")
	setupRoutes(api, appContainer)
	return app.Listen(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
}

func setupRoutes(router fiber.Router, appContainer app.App) {
	task := router.Group("/tasks")
	task.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("All Tasks are return")
	})
	task.Get("/:id", handlers.GetUserByID(appContainer))
}
