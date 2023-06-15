package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
	appContext "user_service/component/app_context"
	"user_service/component/token/jwt"
	userStorage "user_service/stotage"
)

const (
	ErrWrongAuthHeader = "ErrWrongAuthHeader"
	MsgWrongAuthHeader = "wrong authen header"
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) > 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("wrong authen header")
	}
	return parts[1], nil
}

func RequiredAuth(appCtx appContext.AppContext) fiber.Handler {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

	return func(c *fiber.Ctx) error {
		token, err := extractTokenFromHeaderString(c.Get("Authorization"))

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}

		db := appCtx.GetDBConnection()
		store := userStorage.NewSqlModel(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}
		user, err := store.FindDataWithCondition(
			c.Context(),
			map[string]interface{}{"_id": payload.UserId},
		)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": errors.New("user not exists")})
		}

		if user == nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": errors.New("user not exists")})
		}

		if user.Status == 0 {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": errors.New("user deleted")})
		}

		c.Locals("UserCurrent", user)
		return c.Next()
	}
}
