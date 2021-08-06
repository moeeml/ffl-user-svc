package database

import (
	"gitee.com/fireflylove/user-svc/model"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(model.User{})
}
