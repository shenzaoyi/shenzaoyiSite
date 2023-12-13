package model

import (
	"be/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败: " + fmt.Sprintf("%s", err))
	}
	Db = db
	u := User{}
	db.AutoMigrate(&u)
}
