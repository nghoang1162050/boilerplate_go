package main

import (
	"boilerplate_go/pkg/api/middlewares"
	"boilerplate_go/pkg/pub_proc"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	
	// Middleware
	e.Use(middlewares.AuthenticationMiddleware())
	
	// Routes
	pub_proc.InitDbClient()
	pub_proc.PublicRouter(e)

	log.Fatal(e.Start(":8080"))
}
