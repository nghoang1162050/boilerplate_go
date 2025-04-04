package main

import (
	"boilerplate_go/internal/models"
	"boilerplate_go/pkg/pub_proc"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	// prepare database connection (mocked for this example)
	dbUser := "root"
	dbPass	:= "my-secret-pw"
	dbHost := "localhost"
	dbPort := "3307"
	dbName := "boilerplate_go"
	db, err := gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPass, dbHost, dbPort, dbName)),
		&gorm.Config{},
	)
	models.Init(db)

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// prepare repository and services
	// productService := services.InitProductService(nil)
	pub_proc.PublicRouter(e)

	log.Fatal(e.Start(":8080"))
}
