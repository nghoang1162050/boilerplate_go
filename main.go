package main

import (
	"boilerplate_go/internal/cache"
	"boilerplate_go/pkg/api/middlewares"
	"boilerplate_go/pkg/pub_proc"
	"log"

	_ "boilerplate_go/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	// Initialize configuration
	cache.RedisClient = &cache.CacheClient{}
    if err := cache.RedisClient.NewRedisCache(); err != nil {
        log.Fatalf("Failed to initialize Redis: %v", err)
    }

	e := echo.New()
	
	// Middleware
	// e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
	// 	XSSProtection:         "",
	// 	ContentTypeNosniff:    "",
	// 	XFrameOptions:         "",
	// 	HSTSMaxAge:            3600,
	// 	ContentSecurityPolicy: "default-src 'self'",
	// }))
	// custom middleware for retricted apis
	e.Use(middlewares.AuthenticationMiddleware())
	
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Routes
	pub_proc.InitDbClient()
	pub_proc.PublicRouter(e)

	log.Fatal(e.Start(":8080"))
}
