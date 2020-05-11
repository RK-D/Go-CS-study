package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gopm/modules/log"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
		context.JSON(
			http.StatusUnprocessableEntity,
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
		context.JSON(
			http.StatusUnprocessableEntity,
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
		context.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code ": 422,
				"msg":   "用户已存在",
			})
		return
	}
	//不存在创建用户
	//4.创建用户， 查库 查询用户是否重复看手机号是否重复
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(password),
		bcrypt.DefaultCost)
	if err != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code ": 500,
				"msg":   "加密错误",
			})
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasPassword),
	}
	DB.Create(&newUser)

	//5.返回结果
	context.JSON(
		http.StatusUnprocessableEntity,
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
	DB := common.GetDB()
	//获取参数
	telephone := context.PostForm("telephone")
	password := context.PostForm("password")
	//数据验证
	//手机
	if len(telephone) != 11 {
		context.JSON(
			http.StatusUnprocessableEntity,
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
		context.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": 422,
				"msg":  "密码不能少于六位",
			})
		return
	}
	//判断手机号
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		context.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": 422,
				"msg":  "用户不存在",
			})
	}
	//判断密码

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		context.JSON(
			http.StatusUnprocessableEntity, gin.H{
				"code": 400,
				"msg":  "密码错误",
			})
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"code": 500,
				"msg":  "系统异常",
			})
		log.Error("token error : %v", err)
		return
	}
	//返回结果
	context.JSON(
		http.StatusUnprocessableEntity,
		gin.H{
			"code ": 200,
			"data ": gin.H{"token ": token},
			"msg":   "登录成功",
		})
}
