package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"backend/internal/api/routes"
	"backend/internal/pkg/auth"
	"backend/internal/pkg/class"
	"backend/internal/pkg/image"
	"backend/internal/pkg/project"
	"backend/internal/pkg/user"

  "backend/internal/pkg/image/repositories"

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

func imageBaseConnection() (*minio.Client, error) {
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "password",  ""),
		Secure: false,
	 })
	 
	if err != nil {
		log.Fatal(err)
	}

	return minioClient, nil
}

func main() {
	db, cancel, err := databaseConnection()

	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!") 

  ib, err := imageBaseConnection()

	if err != nil {
		log.Fatal("Imagebase Connection Error $s", err)
	}
	fmt.Println("Imagebase connection success!") 

	userCollection := db.Collection("users")
	userRepository := user.NewRepository(userCollection)
	userService    := user.NewService(userRepository)

  authRepository := auth.NewRepository(userCollection)
  authService    := auth.NewService(authRepository)

	projectCollection := db.Collection("project")
	projectRepository := project.NewRepository(projectCollection)
	projectService    := project.NewService(projectRepository)

  classRepository   := class.NewRepository(projectCollection) 
  classService      := class.NewService(classRepository)

	imageCollection := db.Collection("image")
  imageImagebaseRepository := repositories.NewImageRepository(ib)
  imageDatabaseRepository := repositories.NewRepository(imageCollection)
  imageService := image.NewService(imageImagebaseRepository, imageDatabaseRepository)

	app := fiber.New()

	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("MAIN PAGE")
	})

	api := app.Group("/api")

  routes.AuthRoutes(api, authService)

  app.Use(jwtware.New(jwtware.Config {
    SigningKey: jwtware.SigningKey{Key: []byte("secret")},
    Filter: func (c *fiber.Ctx) bool {
      return c.Path() == "/api/user/" && c.Method() == "POST"
    },
  }))

	routes.UserRouter(api, userService)
	routes.ProjectRouter(api, projectService)
	routes.ClassRouter(api, classService)
  routes.ImageRouter(api, imageService)

	defer cancel()

	log.Fatal(app.Listen(":8000"))
}
