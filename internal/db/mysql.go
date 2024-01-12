package db

import (
	"awesomeProject3/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, _ := GetDB()
	db.AutoMigrate(&entity.LogEntity{})
	DB = db
}

func GetDB() (*gorm.DB, error) {
	dsn := "root:qiuyue@tcp(127.0.0.1:3306)/scp?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return DB, err
}
