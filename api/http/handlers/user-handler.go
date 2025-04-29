package handlers

import (
	"tasoryx/app"
	"tasoryx/internal/user/domain"
	"tasoryx/pkg/context"
	appLogger "tasoryx/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

var logger = appLogger.Get().Named("handlers")

func handleError(err error, c fiber.Ctx) error {
	response := &Res{
		Status: fiber.StatusBadRequest,
		Msg:    err.Error(),
		Data:   nil,
	}
	logger.Error(err.Error())
	return c.Status(fiber.StatusBadRequest).JSON(response)
}

func handleSuccess(c fiber.Ctx, data any, msg string) error {
	response := &Res{
		Status: fiber.StatusOK,
		Msg:    "success",
		Data:   data,
	}
	logger.Info(msg)
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetUserByID retrieves a user by ID
func GetUserByID(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		id := c.Params("id")
		user, err := service.GetUserInfo(ctx, id)
		if err != nil {
			return handleError(err, c)
		}
		return handleSuccess(c, user, "User retrieved successfully")
	}
}

// GetUsers retrieves a list of users based on filters
func GetUsers(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		filters := new(domain.FilterUser)
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		if err := c.Bind().Query(filters); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid query parameters",
			})
		}
		users, err := service.GetUsers(ctx, *filters)
		if err != nil {
			return handleError(err, c)
		}
		return handleSuccess(c, users, "Users retrieved successfully")
	}
}

// CreateUser creates a new user
func CreateNewUser(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(domain.User)
		if err := c.Bind().Body(request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		userID, err := service.CreateUser(ctx, *request)
		if err != nil {
			return handleError(err, c)
		}
		request.ID = userID
		return handleSuccess(c, request, "User created successfully")
	}
}

// UpdateUser updates an existing user
func UpdateUser(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(domain.User)
		response := new(Res)
		if err := c.Bind().Body(request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		id := c.Params("id")
		err := service.UpdateUser(ctx, *request, id)
		if err != nil {
			return handleError(err, c)
		}
		return handleSuccess(c, response, "User updated successfully")
	}
}
