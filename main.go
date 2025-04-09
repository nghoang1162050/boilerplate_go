package main

import (
	"log"

	_ "boilerplate_go/docs"
	"boilerplate_go/internal/controller"
	"boilerplate_go/internal/database"
	"boilerplate_go/internal/model"
	"boilerplate_go/internal/repository"
	"boilerplate_go/internal/router"
	"boilerplate_go/internal/usecase"
	"boilerplate_go/internal/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server boilerplate_go server.
// @termsOfService http://swagger.io/terms/
func main() {
	// gen entity models
	// cmd.GenModels()
	// Load .env file
	
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, proceeding with environment variables")
    }

	// Initialize Redis client
	cacheClient := &utils.CacheClient{}
	if err := cacheClient.NewRedisCache(); err != nil {
        log.Fatalf("Failed to initialize Redis: %v", err)
    }
	utils.RedisClient = cacheClient

	// Initialize database connection
	db, err := database.InitDbClient()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	productRepository := repository.NewRepository[model.Product](db)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUseCase)

	e := echo.New()
	router.NewProductRouter(e, productController)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Fatal(e.Start(":8080"))
}
