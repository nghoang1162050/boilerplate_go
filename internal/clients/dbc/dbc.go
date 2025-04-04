package dbc

import (
	"boilerplate_go/internal/clients/dbc/adapters"
	"boilerplate_go/pkg/config"
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

type DBClient struct {
	name       string
	config		config.DatabaseConfig
	silent     bool
	DB         *gorm.DB
	gormConfig *gorm.Config
	adapter    *adapters.Adapter
	driver     gorm.Dialector
	dsn        string
}

func (dbc *DBClient) Connect() {
	var err error
	dbc.DB, err = gorm.Open(dbc.driver, dbc.gormConfig)
	if err != nil {
		fmt.Println(err)
		panic("failed to establish database connection")
	}
}

func (dbc *DBClient) Ping() error {
	var err error
	var d *sql.DB
	d, err = dbc.DB.DB()
	if err != nil {
		return err
		// fmt.Println(err)
		// panic("gorm error")
	}
	err = d.Ping()
	if err != nil {
		return err
		// fmt.Println(err)
		// panic("failed to ping database")
	}
	return nil
}

func (dbc *DBClient) CreateDatabase() error {
	sql, err := dbc.adapter.GetDbCreateStatement()
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}
	err = dbc.DB.Exec(sql + dbc.config.Name + ";").Error
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}
	return nil
}

func (dbc *DBClient) DropDatabase() error {
	sql, err := dbc.adapter.GetDbDropStatement()
	if err != nil {
		return fmt.Errorf("failed to drop database: %w", err)
	}
	err = dbc.DB.Exec(sql + dbc.config.Name + ";").Error
	if err != nil {
		return fmt.Errorf("failed to drop database: %w", err)
	}
	return nil
}
