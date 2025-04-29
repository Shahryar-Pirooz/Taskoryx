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
	setupUserRoutes(router, appContainer)
	setupTaskRoutes(router, appContainer)
}

func setupUserRoutes(router fiber.Router, appContainer app.App) {
	user := router.Group("/users")
	user.Get("/", handlers.GetUsers(appContainer))
	user.Get("/:id", handlers.GetUserByID(appContainer))
	user.Post("/", handlers.CreateNewUser(appContainer))
	user.Put("/:id", handlers.UpdateUser(appContainer))
	// TODO: getUserByEmail and deleteUser
}

func setupTaskRoutes(router fiber.Router, appContainer app.App) {
	task := router.Group("/tasks")
	task.Get("/", handlers.GetTasks(appContainer))
	task.Get("/:id", handlers.GetTaskByID(appContainer))
	task.Post("/", handlers.CreateNewTask(appContainer))
	task.Put("/:id", handlers.UpdateTask(appContainer))
	// TODO: deleteTask
}
