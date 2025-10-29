package main

import (
	"log"
	"context"
	"fmt"
	"time"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"backend/internal/pkg/user"
	"backend/internal/api/routes"
)

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:password@127.0.0.1:27017/")) // Replace with your MongoDB URI

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("bugmarker")
	return db, cancel, nil
}

func main() {
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!") 

	userCollection := db.Collection("users")
	userRepository := user.NewRepository(userCollection)
	userService    := user.NewService(userRepository)

	app := fiber.New()

	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("MAIN PAGE")
	})

	api := app.Group("/api")
	routes.UserRouter(api, userService)

	defer cancel()

	log.Fatal(app.Listen(":8000"))
}
