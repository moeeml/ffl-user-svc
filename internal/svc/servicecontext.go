package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"user-svc/internal/config"
	"user-svc/model"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.Datasource)),
	}
}
