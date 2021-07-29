package svc

import (
	"gitee.com/fireflylove/user-svc/internal/config"
	"gitee.com/fireflylove/user-svc/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.Datasource)),
	}
}
