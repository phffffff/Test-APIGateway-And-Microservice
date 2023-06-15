package userTransport

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	userBusiness "user_service/business"
	"user_service/common"
	appContext "user_service/component/app_context"
	userStorage "user_service/stotage"
)

func DeleteTransport(appCtx appContext.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}
		db := appCtx.GetDBConnection()
		log.Print(c.Locals("UserCurrent"))
		req := c.Locals("UserCurrent").(common.Requester)
		store := userStorage.NewSqlModel(db)
		biz := userBusiness.NewDeleteUserBiz(store, req)

		result, err := biz.DeleteUser(c.Context(), objectId)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"msg": result})
	}
}
