package userTransport

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	userBusiness "user_service/business"
	"user_service/component/app_context"
	"user_service/component/token/jwt"
	userModel "user_service/model"
	userStorage "user_service/stotage"
)

func LoginTransport(appCtx appContext.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var login userModel.UserLogin
		if err := c.BodyParser(&login); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}
		db := appCtx.GetDBConnection()

		store := userStorage.NewSqlModel(db)
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())
		biz := userBusiness.NewLoginBiz(store, tokenProvider, 60*60*24*3)

		token, err := biz.Login(context.TODO(), &login)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": token})
	}
}
