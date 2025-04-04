package adapters

import "gorm.io/gorm"

type IAdapter interface {
	// SetConfig(config features.DatabaseConfig)
	GetDriver() (gorm.Dialector, error)
	GetServerDriver() (gorm.Dialector, error)
	GetDSN() (string, error)
	GetServerDSN() (string, error)
	GetDbCreateStatement() (string, error)
	GetDbDropStatement() (string, error)
	ValidateConfig() error
}

type Adapter struct {
	IAdapter
	adapters        map[string]IAdapter
	defaultPlatform string
	currentPlatform string
	// config          features.DatabaseConfig
}


var Adapters = &Adapter{
	defaultPlatform: "mysql",
	currentPlatform: "mysql",
	adapters:        make(map[string]IAdapter),
}
