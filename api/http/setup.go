package http

import (
	"fmt"
	"tasoryx/api/http/handlers"
	"tasoryx/api/http/middlewares"
	"tasoryx/app"
	"tasoryx/config"
	log "tasoryx/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

var logger = log.Get().Named("http server")

func Run(appContainer app.App, cfg config.ServerConfig) error {
	logger.Info("Starting HTTP server...")
	app := fiber.New(fiber.Config{AppName: "Taskoryx"})
	api := app.Group("/api/v1")
	logger.Info("Setting up routes...")
	setupRoutes(api, appContainer)
	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	logger.Info(fmt.Sprintf("Listening on %s", address))
	return app.Listen(address)
}

func setupRoutes(router fiber.Router, appContainer app.App) {
	logger.Info("Setting up auth routes...")
	setupAuthRoutes(router, appContainer)
	logger.Info("Setting up user routes...")
	setupUserRoutes(router, appContainer)
	logger.Info("Setting up task routes...")
	setupTaskRoutes(router, appContainer)
}

func setupAuthRoutes(router fiber.Router, appContainer app.App) {
	auth := router.Group("/auth")
	auth.Post("/login", handlers.Login(appContainer))
	auth.Post("/Refresh", handlers.GetNewAccessToken(appContainer))
}

func setupUserRoutes(router fiber.Router, appContainer app.App) {
	user := router.Group("/users")
	logger.Info("Configuring user routes...")
	user.Get("/", handlers.GetUsers(appContainer))
	user.Get("/:id", handlers.GetUserByID(appContainer))
	user.Post("/", handlers.CreateNewUser(appContainer))
	user.Put("/:id", handlers.UpdateUser(appContainer))
	logger.Info("User routes configured.")
}

func setupTaskRoutes(router fiber.Router, appContainer app.App) {
	task := router.Group("/tasks", middlewares.RequireAuth(appContainer))
	logger.Info("Configuring task routes...")
	task.Get("/", handlers.GetTasks(appContainer))
	task.Get("/:id", handlers.GetTaskByID(appContainer))
	task.Post("/", handlers.CreateNewTask(appContainer))
	task.Put("/:id", handlers.UpdateTask(appContainer))
	task.Delete("/:id", handlers.DeleteTask(appContainer))
	logger.Info("Task routes configured.")
}
