package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Engin
	router := gin.Default()
	//router := gin.New()
	// api注册
	//post 请求
	router.POST("/api/demo/register", func(context *gin.Context) {
		//1.获取参数
		name := context.PostForm("name")
		telephone := context.PostForm("telephone")
		password := context.PostForm("password")

		//2.数据验证
		//手机
		if len(telephone) != 11 {
			context.JSON(http.StatusUnprocessableEntity,
				// map[string] interface{}
				//gin源码：type H map[string]interface{}
				//H is a shortcut for map[string]interface{}
				gin.H{
					"code": 422,
					"msg":  "手机号必须为11位",
				})
			return
		}
		//密码
		if len(password) < 6 {
			context.JSON(http.StatusUnprocessableEntity,
				gin.H{
					"code": 422,
					"msg":  "密码不能少于六位",
				})
			return
		}
		//名称验证，没有传一个随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}

		//3.手机验证 ,查库
		log.Println("tel")

		//4.创建用户， 查库
		log.Println("user")

		//5.返回结果

		log.Println(">>>> hello gin start <<<<")
		context.JSON(200, gin.H{
			"code":    200,
			"msg":     "注册成功",
			"success": true,
		})
	})
	// 指定地址和端口号
	router.Run("localhost:9090")
}

func RandomString(i int) string {
	var letters = []byte("qKLewFZXrtYIOyuROPAioYUIpaABNsDCdfEghUjkQHJzxRcPvVbSlnMmW")
	result := make([]byte, i)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
