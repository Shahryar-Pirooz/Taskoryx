package handlers

import (
	"tasoryx/app"
	"tasoryx/internal/user/domain"
	"tasoryx/pkg/context"
	appLogger "tasoryx/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

var logger = appLogger.Get().Named("handlers")

func GetUserByID(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		response := new(Res)
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		id := c.Params("id")
		user, err := service.GetUserInfo(ctx, id)
		if err != nil {
			logger.Error(err.Error())
			response = &Res{
				Status: fiber.StatusBadRequest,
				Msg:    err.Error(),
				Data:   nil,
			}
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
		response = &Res{
			Status: fiber.StatusOK,
			Msg:    "success",
			Data:   user,
		}
		logger.Info("Get data success : " + user.ID)
		return c.Status(fiber.StatusOK).JSON(response)
	}
}

func GetUsers(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		response := new(Res)
		filters := new(domain.FilterUser)
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		if err := c.Bind().Query(filters); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid query parameters",
			})
		}
		var (
			users []domain.User
			err   error
		)
		if *filters != (domain.FilterUser{}) {
			users, err = service.GetUsers(ctx, *filters)
		} else {
			users, err = service.GetUsers(ctx)
		}
		if err != nil {
		}
		response = &Res{
			Status: fiber.StatusOK,
			Msg:    "success",
			Data:   users,
		}
		return c.Status(fiber.StatusOK).JSON(response)
	}
}

func CreateNewUser(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(domain.User)
		response := new(Res)
		if err := c.Bind().Body(request); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		userID, err := service.CreateUser(ctx, *request)
		if err != nil {
			response.Msg = err.Error()
			response.Status = fiber.StatusBadGateway
			return c.Status(fiber.StatusBadGateway).JSON(response)
		}
		request.ID = userID
		response = &Res{
			Status: fiber.StatusOK,
			Msg:    "success",
			Data:   request,
		}
		return c.Status(fiber.StatusOK).JSON(response)
	}
}

func UpdateUser(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		return nil
	}
}
