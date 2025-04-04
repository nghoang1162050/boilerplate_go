package dbc

import "gorm.io/gorm"

type DBClient struct {
	name       string
	silent     bool
	DB         *gorm.DB
	gormConfig *gorm.Config
	driver     gorm.Dialector
	dsn        string
}