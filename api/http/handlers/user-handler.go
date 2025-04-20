package handlers

import (
	"net/http"
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
				Status: http.StatusBadRequest,
				Msg:    err.Error(),
				Data:   nil,
			}
			return c.Status(http.StatusBadRequest).JSON(response)
		}
		response = &Res{
			Status: http.StatusOK,
			Msg:    "success",
			Data:   user,
		}
		logger.Info("Get data success : " + user.ID)
		return c.Status(http.StatusOK).JSON(response)
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
			Status: http.StatusOK,
			Msg:    "success",
			Data:   users,
		}
		return c.Status(http.StatusOK).JSON(response)
	}
}
