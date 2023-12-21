package router

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vaibhavm18/go-url-shortner/internal/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddHomeGroup(app *fiber.App) {
	homeGroup := app.Group("/")
	homeGroup.Get("/health", checkHealth)

}

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
}

func checkHealth(c *fiber.Ctx) error {
	col := config.GetDBCollection("Users")

	user := User{
		Username: "vaibhav1",
	}

	res, err := col.InsertOne(context.Background(), user)

	if err != nil {
		fmt.Println(err)
		return c.Status(403).JSON(fiber.Map{
			"message": "failed to create a user",
		})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Hello, World 👋!", "data": res})
}
