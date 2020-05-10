package main

import (
	"github.com/gin-gonic/gin"
	"pers.study/cstest/common"
)

func main() {
	//获取db，延迟关闭
	db := common.InoDB()
	defer db.Close()

	router := gin.Default()
	//router := gin.New()
	// api注册
	//post 请求
	router = CollectRoute(router)
	// 指定地址和端口号
	panic(router.Run())
}
