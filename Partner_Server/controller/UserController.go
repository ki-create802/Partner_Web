package controller

import (
	"GolandProject/common"
	"GolandProject/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	//"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	DB *gorm.DB
}

type InUserController interface {
	Register(c *gin.Context) //实现注册功能
	Login(c *gin.Context)    //实现登录功能
}

func NewUserController() InUserController {
	db := common.GetDB()
	db.Table("user").AutoMigrate(model.User{})
	//db.Table("admin").AutoMigrate(model.User{})
	return UserController{DB: db}
}

// 注册
func (a UserController) Register(c *gin.Context) {

	// 获取context中的参数
	var requestUser model.User
	// 将请求中的JSON数据绑定到User结构体中，方便后续操作
	c.Bind(&requestUser)
	userName := requestUser.UName
	userEmail := requestUser.UEmail
	password := requestUser.UKey

	// 验证数据
	var user model.User
	a.DB.Table("user").Where("UEmail=?", userEmail).First(&user)
	if user.UID != "" { // 该邮箱已存在，返回422，且返回提示信息
		common.Fail(c, 422, nil, "该邮箱已注册")
		return
	}

	//// 密码加密
	//hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// 创建用户
	newUser := model.User{
		UName:  userName,
		UEmail: userEmail,
		//Password:    string(hashedPassword),
		UKey: password,
	}
	a.DB.Table("user").Create(&newUser)
	//***********************************************************
	//还缺少的功能 分配用户id
	//密码加密功能
	//对于用户名要不要不让重复
	//默认头像图标
	//************************************************************

	// 返回成功响应
	common.Success(c, nil, "注册成功")
}

// Login 登录
func (a UserController) Login(c *gin.Context) {

	// 获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	userEmail := requestUser.UEmail
	password := requestUser.UKey

	// 数据验证
	var user model.User
	a.DB.Table("user").Where("UEmail=?", userEmail).First(&user)
	if user.UID == "" {
		common.Fail(c, 422, nil, "用户不存在")
		return
	}

	if user.UKey != password {
		common.Fail(c, 422, nil, "密码错误")
		return
	}

	//// 判断密码是否正确
	//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
	//	common.Fail(c, 422, nil, "密码错误")
	//	return
	//}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		common.Fail(c, 500, nil, "系统异常")
		return
	}

	// 返回结果
	common.Success(c, gin.H{"token": token}, "登录成功")
}
