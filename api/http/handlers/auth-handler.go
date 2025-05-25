package handlers

import (
	"errors"
	"fmt"
	"tasoryx/app"
	"tasoryx/pkg/context"
	"tasoryx/pkg/jwt"
	"time"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

func Login(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(LoginReq)
		if err := c.Bind().Body(request); err != nil {
			return HandleError(err, c, fiber.StatusBadRequest)
		}
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		user, err := service.GetUserByEmail(ctx, request.Email)
		if err != nil {
			return HandleError(err, c, fiber.StatusBadGateway)
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			return HandleError(err, c, fiber.StatusUnauthorized)
		}
		accToken, err := jwt.GenerateToken(user.ID, appContainer.Config().Jwt.AccessKey, int8(user.Role), time.Minute*15)
		if err != nil {
			return HandleError(err, c, fiber.StatusInternalServerError)
		}
		refToken, err := jwt.GenerateToken(user.ID, appContainer.Config().Jwt.RefreshKey, int8(user.Role), time.Hour*24*7)
		if err != nil {
			return HandleError(err, c, fiber.StatusInternalServerError)
		}
		c.Set("Authorization", "Bearer "+accToken)
		data := UserRes{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  int8(user.Role),
		}
		c.Cookie(&fiber.Cookie{
			Expires:  time.Now().Add(time.Hour * 24 * 7),
			Name:     "ref",
			Value:    refToken,
			Secure:   true,
			HTTPOnly: true,
			SameSite: "Strict",
		})
		key := fmt.Sprintf("AUTH.REFRESH_TOKEN:%s", user.ID)
		if err := appContainer.Cache().Set(ctx, key, refToken); err != nil {
			return HandleError(err, c, fiber.StatusBadGateway)
		}
		return HandleSuccess(c, data, "Login successful")
	}
}

func GetNewAccessToken(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		refToken := c.Cookies("ref")
		if refToken == "" {
			return HandleError(errors.New("refresh token missing"), c, fiber.StatusUnauthorized)
		}

		claims, err := jwt.ValidationToken(refToken, appContainer.Config().Jwt.RefreshKey)
		if err != nil {
			return HandleError(err, c, fiber.StatusUnauthorized)
		}

		ctx := context.NewAppContext(c.Context())

		key := fmt.Sprintf("AUTH.REFRESH_TOKEN:%s", claims.UserID)
		storedToken, err := appContainer.Cache().Get(ctx, key)
		if err != nil || storedToken != refToken {
			return HandleError(errors.New("refresh token revoked or invalid"), c, fiber.StatusUnauthorized)
		}

		newAccessToken, err := jwt.GenerateToken(claims.UserID, appContainer.Config().Jwt.AccessKey, claims.Role, time.Minute*15)
		if err != nil {
			return HandleError(err, c, fiber.StatusInternalServerError)
		}

		c.Set("Authorization", "Bearer "+newAccessToken)
		return HandleSuccess(c, claims.UserID, "New access token generated")
	}
}
