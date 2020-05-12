package main

import (
	"github.com/gin-gonic/gin"
	"pers.study/cstest/controller"
	"pers.study/cstest/middleware"
)

//路由
func CollectRoute(router *gin.Engine) *gin.Engine {
	//注册
	router.POST("/api/demo/register", controller.Register)
	//登录
	router.POST("/api/demo/login", controller.Login)
	//用户信息,使用中间件保护信息接口
	router.GET("/api/demo/userInfo", middleware.AuthMiddleWare(), controller.UserInfo)
	//router.Run("localhost:9090")
	return router
}
