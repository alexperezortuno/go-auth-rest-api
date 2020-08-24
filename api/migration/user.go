package migration

import (
	"../model"
	"github.com/jinzhu/gorm"
)

func UserMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
