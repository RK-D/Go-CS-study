package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"pers.study/cstest/model"
)

var DB *gorm.DB

//创建连接池
func InoDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
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
