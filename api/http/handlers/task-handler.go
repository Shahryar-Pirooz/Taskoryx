package handlers

import (
	"tasoryx/app"
	"tasoryx/internal/task/domain"
	"tasoryx/pkg/context"

	"github.com/gofiber/fiber/v3"
)

func GetTasks(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		filters := new(domain.FilterTask)
		ctx := context.NewAppContext(c.Context())
		service := appContainer.TaskService(ctx)
		if err := c.Bind().Query(filters); err != nil {
			return HandleError(err, c, fiber.StatusBadRequest)
		}
		tasks, err := service.GetTasks(ctx, *filters)
		if err != nil {
			return HandleError(err, c, fiber.StatusNotFound)
		}
		return HandleSuccess(c, tasks, "Tasks retrieved successfully")
	}
}
func GetTaskByID(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := context.NewAppContext(c.Context())
		service := appContainer.TaskService(ctx)
		id := c.Params("id")
		task, err := service.GetTaskByID(ctx, id)
		if err != nil {
			return HandleError(err, c, fiber.StatusNotFound)
		}
		return HandleSuccess(c, task, "Task retrieved successfully")
	}
}
func CreateNewTask(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(domain.Task)
		ctx := context.NewAppContext(c.Context())
		service := appContainer.TaskService(ctx)
		if err := c.Bind().Body(request); err != nil {
			return HandleError(err, c, fiber.StatusBadRequest)
		}
		taskID, err := service.CreateTask(ctx, *request)
		if err != nil {
			return HandleError(err, c, fiber.StatusInternalServerError)
		}
		return HandleSuccess(c, taskID, "Task created successfully")
	}
}
func UpdateTask(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(domain.Task)
		ctx := context.NewAppContext(c.Context())
		service := appContainer.TaskService(ctx)
		id := c.Params("id")
		if err := c.Bind().Body(request); err != nil {
			return HandleError(err, c, fiber.StatusBadRequest)
		}
		if err := service.UpdateTask(ctx, *request, id); err != nil {
			return HandleError(err, c, fiber.StatusInternalServerError)
		}
		return HandleSuccess(c, nil, "Task updated successfully")
	}
}
func DeleteTask(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := context.NewAppContext(c.Context())
		service := appContainer.TaskService(ctx)
		id := c.Params("id")
		if err := service.DeleteTask(ctx, id); err != nil {
			return HandleError(err, c, fiber.StatusInternalServerError)
		}
		return HandleSuccess(c, nil, "Task deleted successfully")
	}
}
