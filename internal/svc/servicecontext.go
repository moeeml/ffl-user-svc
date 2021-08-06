package svc

import (
	"gitee.com/fireflylove/user-svc/database"
	"gitee.com/fireflylove/user-svc/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	DB *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		DB : database.NewMysql(&c),
	}
}
