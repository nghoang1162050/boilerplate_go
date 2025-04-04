package main

import (
	"boilerplate_go/internal/cache"
	"boilerplate_go/pkg/api/middlewares"
	"boilerplate_go/pkg/pub_proc"
	"log"

	"github.com/labstack/echo/v4"
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
	
	// Routes
	pub_proc.InitDbClient()
	pub_proc.PublicRouter(e)

	log.Fatal(e.Start(":8080"))
}
