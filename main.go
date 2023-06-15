package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	appContext "user_service/component/app_context"
	"user_service/middleware"
	userTransport "user_service/transport"
)

func main() {

	//nếu sài variable env của goland
	//uri := os.Getenv("LOCAL_HOST")

	if err := godotenv.Load(); err != nil {
		log.Fatal("err:", err)
	}
	uri := os.Getenv("URI")

	//serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	//opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("err:", err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal("err:", err)
		}
	}()

	db := client.Database(os.Getenv("DATABASE_NAME"))
	col := db.Collection(os.Getenv("COLLECTION"))

	secretKey := "admin"

	app := fiber.New()

	appCtx := appContext.NewAppContext(col, secretKey)
	//app.Use(middleware.Recover(appCtx))

	app.Post("/login", userTransport.LoginTransport(appCtx))
	app.Delete("/:id/delete", middleware.RequiredAuth(appCtx), userTransport.DeleteTransport(appCtx))
	//delete
	//id := "648ad49a5528f5dfae809e71"
	//objectId, _ := primitive.ObjectIDFromHex(id)
	//
	//result, err := store.Delete(context.TODO(), &objectId)
	//if err != nil {
	//	log.Fatal("err:", err)
	//}
	//log.Print("result:", result)

	//update
	//dataUpdate := userModel.UserWrite{
	//	Name:     "Admin",
	//	Username: "admin",
	//	Password: "admin",
	//	Status:   0,
	//}
	//
	//id := "648ad485e7016172ec341380"
	//objectId, _ := primitive.ObjectIDFromHex(id)
	//
	//result, err := store.Update(context.TODO(), &objectId, &dataUpdate)
	//if err != nil {
	//	log.Fatal("err:", err)
	//}
	//log.Print("result:", result)

	//find

	//user, err := store.FindDataWithCondition(context.TODO(), map[string]interface{}{"username": "admin", "password": "admin"})
	//if err != nil {
	//	log.Fatal("err:", err)
	//}
	//log.Print("user:", user)

	//list
	//filter := userModel.Filter{
	//	Status: []int{0, 1},
	//}
	//var filter userModel.Filter
	//filter.Status = []int{0, 1}
	//list, err := store.ListDataWithFilter(context.TODO(), &filter)
	//if err != nil {
	//	log.Fatal("err:", err)
	//}
	//log.Print("list:", list)

	//create
	//data := userModel.UserWrite{
	//	Name:     "Kiên",
	//	Username: "M12uk1",
	//	Password: "@Klov3x124n",
	//	Status:   1,
	//}
	//rs, err := store.Create(context.TODO(), &data)
	//if err != nil {
	//	log.Fatal("err:", err)
	//}
	//log.Print("Inserted document with _id: ", rs)

	log.Fatal(app.Listen(":3000"))
}
