package handlers

import (
	"tasoryx/app"
	"tasoryx/internal/user/domain"
	"tasoryx/pkg/context"

	"github.com/gofiber/fiber/v3"
)

// GetUserByID retrieves a user by ID
func GetUserByID(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		id := c.Params("id")
		user, err := service.GetUserInfo(ctx, id)
		if err != nil {
			return HandleError(err, c, fiber.StatusNotFound)
		}
		return HandleSuccess(c, user, "User retrieved successfully")
	}
}

// GetUsers retrieves a list of users based on filters
func GetUsers(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		filters := new(domain.FilterUser)
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		if err := c.Bind().Query(filters); err != nil {
			return HandleError(err, c, fiber.StatusBadRequest)
		}
		users, err := service.GetUsers(ctx, *filters)
		if err != nil {
			return HandleError(err, c, fiber.StatusNotFound)
		}
		return HandleSuccess(c, users, "Users retrieved successfully")
	}
}

func GetUserByEmail(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(domain.User)
		if err := c.Bind().Body(request); err != nil {
			return HandleError(err, c, fiber.StatusBadRequest)
		}
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		user, err := service.GetUserByEmail(ctx, request.Email)
		if err != nil {
			return HandleError(err, c, fiber.StatusInternalServerError)
		}
		return HandleSuccess(c, user, "User retrieved successfully")
	}
}

// CreateUser creates a new user
func CreateNewUser(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(domain.User)
		if err := c.Bind().Body(request); err != nil {
			return HandleError(err, c, fiber.StatusBadRequest)
		}
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		userID, err := service.CreateUser(ctx, *request)
		if err != nil {
			return HandleError(err, c, fiber.StatusInternalServerError)
		}
		request.ID = userID
		return HandleSuccess(c, request, "User created successfully")
	}
}

// UpdateUser updates an existing user
func UpdateUser(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(domain.User)
		response := new(Res)
		if err := c.Bind().Body(request); err != nil {
			return HandleError(err, c, fiber.StatusBadRequest)
		}
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		id := c.Params("id")
		err := service.UpdateUser(ctx, *request, id)
		if err != nil {
			return HandleError(err, c, fiber.StatusInternalServerError)
		}
		return HandleSuccess(c, response, "User updated successfully")
	}
}
