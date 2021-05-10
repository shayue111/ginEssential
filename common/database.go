package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shayue/ginessential/model"
)

var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() {
	host := "localhost"
	port := "3306"
	database := "gin_essential"
	username := "root"
	password := "shayue123"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	_ = db.AutoMigrate(&model.User{})

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
