package handlers

import (
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
		accToken, err := jwt.GenerateToken(user.ID, appContainer.Config().Jwt.AccessKey, time.Minute*15)
		if err != nil {
			return HandleError(err, c, fiber.StatusInternalServerError)
		}
		refToken, err := jwt.GenerateToken(user.ID, appContainer.Config().Jwt.RefreshKey, time.Hour*24*7)
		if err != nil {
			return HandleError(err, c, fiber.StatusInternalServerError)
		}
		data := struct {
			User        UserRes `json:"user"`
			AccessToken string  `json:"accessToken"`
		}{
			User: UserRes{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
				Role:  int8(user.Role),
			},
			AccessToken: accToken,
		}
		c.Cookie(&fiber.Cookie{
			Expires:  time.Now().Add(time.Hour * 24 * 7),
			Name:     "ref",
			Value:    refToken,
			Secure:   true,
			HTTPOnly: true,
			SameSite: "Strict",
		})
		return HandleSuccess(c, data, "Login successful")
	}
}

func GetNewAccessToken(appContainer app.App) fiber.Handler {
	return func(c fiber.Ctx) error {
		request := new(RefreshTokenReq)
		if err := c.Bind().Body(request); err != nil {
			return HandleError(err, c, fiber.StatusBadRequest)
		}
		claims, err := jwt.ValidationToken(request.AccessToken, appContainer.Config().Jwt.AccessKey)
		if err != nil || claims.UserID != request.UserID {
			return HandleError(err, c, fiber.StatusNonAuthoritativeInformation)
		}
		newAccessToken, err := jwt.GenerateToken(claims.UserID, appContainer.Config().Jwt.AccessKey, time.Minute*15)
		if err != nil {
			return HandleError(err, c, fiber.StatusInternalServerError)
		}
		data := struct {
			AccessToke string `json:"access_token"`
		}{
			AccessToke: newAccessToken,
		}
		return HandleSuccess(c, data, "Refresh token generated successfully")

	}
}

// TODO: Store refresh tokens in DB or Redis to support revocation and logout
// TODO: Don't need to generate a refresh token if it doesn't expire
