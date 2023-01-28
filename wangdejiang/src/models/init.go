package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Db 数据库实例
var Db = Init()

func Init() *gorm.DB {
	dsn := "root:imAei@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil || db.Error != nil {
		panic("connect error")
	}
	return db
}
