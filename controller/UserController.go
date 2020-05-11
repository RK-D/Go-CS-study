package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"pers.study/cstest/common"
	"pers.study/cstest/model"
	"pers.study/cstest/util"
)

func Register(context *gin.Context) {

	DB := common.GetDB()
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
		name = util.RandomString(10)
	}

	//3.手机验证 ,查库
	if isTelephoneExist(DB, telephone) {
		context.JSON(http.StatusUnprocessableEntity,
			gin.H{
				"code ": 422,
				"msg":   "用户已存在",
			})
		return
	}
	//不存在创建用户
	//4.创建用户， 查库
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	//5.返回结果
	context.JSON(http.StatusUnprocessableEntity,
		gin.H{
			"code ": 200,
			"msg ":  "注册成功",
		})

}

//查询手机号实现
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Login(context *gin.Context) {
	//获取参数
	//数据验证
	//判断手机号
	//判断密码
	//发放token
	//返回结果
}
