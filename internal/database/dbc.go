package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/prometheus"
)

func InitDbClient() (*gorm.DB, error) {
	dbUser := "root"
	dbPass	:= "my-secret-pw"
	dbHost := "localhost"
	dbPort := "3307"
	dbName := "ecommerce_db"
	db, err := gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPass, dbHost, dbPort, dbName)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil, err
	}
	
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

	return db, err
}