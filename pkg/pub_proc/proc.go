package pub_proc

import (
	"boilerplate_go/internal/models"
	"boilerplate_go/pkg/api/routers"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func PublicRouter(e *echo.Echo) {
	routers.InitUserRouters(e)
	routers.InitProductRouter(e)
}

func InitDbClient() {
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
}