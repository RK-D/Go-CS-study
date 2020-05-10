package main

import (
	"github.com/gin-gonic/gin"
	"pers.study/cstest/controller"
)

func CollectRoute(router *gin.Engine) *gin.Engine {
	router.POST("/api/demo/register", controller.Register)
	//router.Run("localhost:9090")
	return router
}
