package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"backend/internal/api/routes"
	"backend/internal/pkg/auth"
	"backend/internal/pkg/user"

	jwtware "github.com/gofiber/contrib/jwt"
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

  authRepository := auth.NewRepository(userCollection)
  authService    := auth.NewService(authRepository)

	app := fiber.New()

	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("MAIN PAGE")
	})

  routes.AuthRoutes(app, authService)

  app.Use(jwtware.New(jwtware.Config {
    SigningKey: jwtware.SigningKey{Key: []byte("secret")},
    Filter: func (c *fiber.Ctx) bool {
      return c.Path() == "/api/user/" && c.Method() == "POST"
    },
  }))

	api := app.Group("/api")
	routes.UserRouter(api, userService)

	defer cancel()

	log.Fatal(app.Listen(":8000"))
}
