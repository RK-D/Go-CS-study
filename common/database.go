package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"pers.study/cstest/model"
)

var DB *gorm.DB

//创建连接池
func InoDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "go_study"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}
	//自动创建数据表
	db.AutoMigrate(&model.User{})
	//一直没赋值错误 runtime error: invalid memory address or nil pointer dereference
	//panicmem: panic(memoryError)
	DB = db
	return db
}

//定义获取方法
func GetDB() *gorm.DB {
	return DB
}
