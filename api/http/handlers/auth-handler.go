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
			return handleError(err, c, fiber.StatusBadRequest)
		}
		ctx := context.NewAppContext(c.Context())
		service := appContainer.UserService(ctx)
		user, err := service.GetUserByEmail(ctx, request.Email)
		if err != nil {
			return handleError(err, c, fiber.StatusBadGateway)
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			return handleError(err, c, fiber.StatusUnauthorized)
		}
		accToken, err := jwt.GenerateToken(user.ID, appContainer.Config().Jwt.Access_key, time.Minute*15)
		if err != nil {
			return handleError(err, c, fiber.StatusInternalServerError)
		}
		refToken, err := jwt.GenerateToken(user.ID, appContainer.Config().Jwt.Refresh_key, time.Hour*24*7)
		if err != nil {
			return handleError(err, c, fiber.StatusInternalServerError)
		}
		data := struct {
			user        UserRes
			accessToken string
		}{
			user: UserRes{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
				Role:  int8(user.Role),
			},
			accessToken: accToken,
		}
		c.Cookie(&fiber.Cookie{
			Expires:  time.Now().Add(time.Hour * 24 * 7),
			Name:     "ref",
			Value:    refToken,
			Secure:   true,
			HTTPOnly: true,
		})
		return handleSuccess(c, data, "Login successful")
	}
}

// TODO: Store refresh tokens in DB or Redis to support revocation and logout
