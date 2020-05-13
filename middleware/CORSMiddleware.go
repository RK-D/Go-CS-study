package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//解决浏览器同源问题, 前端需正确安装并使用axios 和vue-axios
func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if context.Request.Method == http.MethodOptions {
			context.AbortWithStatus(200)
		} else {
			context.Next()
		}
	}
}
