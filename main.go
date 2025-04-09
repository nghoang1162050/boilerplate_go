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

	"github.com/labstack/echo/v4"
)

func main() {
	// gen entity models
	// cmd.GenModels()

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

	log.Fatal(e.Start(":8080"))
}
