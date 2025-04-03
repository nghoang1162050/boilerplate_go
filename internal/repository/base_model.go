package repository

import "gorm.io/gorm"

var db *gorm.DB

func Init(database *gorm.DB) {
	db = database
}

type BaseModel struct {
	ID int `gorm:"id;type:"int";primaryKey"`
}