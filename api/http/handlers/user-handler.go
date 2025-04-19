package handlers

import (
	"net/http"
	"tasoryx/app"
	"tasoryx/pkg/context"
	appLogger "tasoryx/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

var logger = appLogger.Get().Named("handlers")

func GetUserByID(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		id := c.Params("id")
		user, err := service.GetUserInfo(ctx, id)
		if err != nil {
			logger.Error(err.Error())
		}
		logger.Info("Get data success : " + user.ID)
		return c.Status(http.StatusOK).JSON(user)
	}
}
