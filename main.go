package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"os"
	"pers.study/cstest/common"
)

func main() {
	//项目启动最开始先读取配置
	InitConfig()
	//获取db，延迟关闭
	db := common.InoDB()
	defer db.Close()

	router := gin.Default()
	//router := gin.New()
	// api注册
	//post 请求
	router = CollectRoute(router)
	// 指定地址和端口号
	port := viper.GetString("server.port")
	if port != "" {
		panic(router.Run(":" + port))
	}
	panic(router.Run())
}

func InitConfig() {
	//读取yml配置文件，名称，类型，路径
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {

	}
}
