package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	_ "boilerplate_go/docs"
	"boilerplate_go/internal/controller"
	"boilerplate_go/internal/database"
	"boilerplate_go/internal/middleware"
	"boilerplate_go/internal/model"
	"boilerplate_go/internal/repository"
	"boilerplate_go/internal/router"
	"boilerplate_go/internal/usecase"
	"boilerplate_go/internal/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleWare "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server boilerplate_go server.
// @termsOfService http://swagger.io/terms/
// @BasePath /api
func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

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

	// Initialize Casbin enforcer
	enforcer := utils.InitCasbin(db)
	
	// user repository
	userRepository := repository.NewUserRepository(db)
	
	// auth repository, usecase, and controller
	authRepository := repository.NewRepository[model.User](db)
	authUseCase := usecase.NewAuthUseCase(authRepository, userRepository)
	authController := controller.NewAuthController(authUseCase)

	// product repository, usecase, and controller
	productRepository := repository.NewRepository[model.Product](db)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUseCase)

	// echo app instance
	e := echo.New()
	e.Use(echoMiddleWare.LoggerWithConfig(echoMiddleWare.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","remote_ip":"${remote_ip}","host":"${host}","method":"${method}","uri":"${uri}","status":${status}}` + "\n",
		Output: os.Stdout,
	}))

	apiGroup := e.Group("/api")

	// Authentication middleware
	e.Use(middleware.JWTMiddleware())

	// Authorization middleware
	e.Use(middleware.CasbinMiddleware(enforcer))

	router.NewProductRouter(apiGroup, productController)
	router.NewAuthRouter(apiGroup, authController)

	// Swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// sync routes to Casbin
	utils.SyncRoutersToCasbin(e, enforcer)

	log.Fatal(e.Start(":8080"))
}
