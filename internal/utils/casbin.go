package utils

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func InitCasbin(db *gorm.DB) *casbin.Enforcer {
	adapter, _ := gormadapter.NewAdapterByDB(db)
	e, _ := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	e.LoadPolicy()

	return e
}