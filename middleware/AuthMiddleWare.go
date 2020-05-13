package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pers.study/cstest/common"
	"pers.study/cstest/model"
	"strings"
)

//中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取authorization header
		tokenString := context.GetHeader("Authorization")

		//验证格式 不能为空，且头部为Bearer ，否则认为没有传token ! 和空格不要漏
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnprocessableEntity,
				gin.H{
					"code ": 401,
					"msg":   "权限不足",
				})
			context.Abort()
			return
		}
		//"Bearer "带上空格7位，从第七位开始往后提取有效部分
		tokenString = tokenString[7:]

		// 解析token 无效的token才显示权限不足
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnprocessableEntity,
				gin.H{
					"code ": 401,
					"msg ":  "权限不足",
				})
			context.Abort()
			return
		}
		//通过验证 ，获取claims 的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		//验证用户首存在用户
		if user.ID == 0 {
			context.JSON(http.StatusUnprocessableEntity,
				gin.H{
					"code ": 401,
					"msg ":  "权限不足",
				})
			context.Abort()
			return
		}
		//用户信息存在
		context.Set("user", user)
		context.Next()
	}
}
