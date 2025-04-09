package cmd

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenModels() {
	// TODO: modify the DSN to your database connection string by env variable
	dsn := "root:my-secret-pw@tcp(localhost:3307)/ecommerce_db?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

	g := gen.NewGenerator(gen.Config{
        OutPath: "internal/model",
    })

	g.UseDB(db)
	g.GenerateAllTable()
	g.Execute()
}