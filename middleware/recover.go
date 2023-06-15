package middleware

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	appContext "user_service/component/app_context"
)

func Recover(appCtx appContext.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				c.Locals("Content-Type", "application/json")

				if err != nil {
					c.Status(http.StatusInternalServerError).JSON(fiber.Map{"err:": "something error with server"})
					return
				}

				c.Status(http.StatusInternalServerError).JSON(fiber.Map{"err:": "something error with server"})
				return
			}
		}()

		return c.Next()
	}
}
