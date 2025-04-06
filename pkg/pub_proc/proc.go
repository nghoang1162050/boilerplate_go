package pub_proc

import (
	"boilerplate_go/internal/models"
	"boilerplate_go/pkg/api/routers"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/prometheus"
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
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	models.Init(db)
	
	db.Use(prometheus.New(prometheus.Config{
		DBName: dbName,
		RefreshInterval: 15,
		PushAddr: "127.0.0.1", // push metrics if `PushAddr` configured
		StartServer: true,
		HTTPServerPort:  8081,  // configure http server port, default port 8080 (if you have configured multiple instances, only the first `HTTPServerPort` will be used to start server)
		MetricsCollector: []prometheus.MetricsCollector {
		  &prometheus.MySQL{
			VariableNames: []string{"Threads_running"},
		  },
		},  // user defined metrics
	  }))
}