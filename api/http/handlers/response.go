package handlers

import (
	appLogger "tasoryx/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

type Res struct {
	Status int    `json:"status"`
	Msg    string `json:"message"`
	Data   any    `json:"data"`
}

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
