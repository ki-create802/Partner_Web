package controller

import (
	"Partner_Web/Partner_Server/common"
	"Partner_Web/Partner_Server/model"
	"Partner_Web/Partner_Server/pkg/redis"
	"Partner_Web/Partner_Server/services"
	"Partner_Web/Partner_Server/utils"
	"fmt"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	//"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	DB *gorm.DB
}

type InUserController interface {
	Register(c *gin.Context)        //实现注册功能
	Login(c *gin.Context)           //实现登录功能
	Logout(c *gin.Context)          //用户登出
	GuanZhu(c *gin.Context)         //用户关注
	FansNum(c *gin.Context)         //返回粉丝数
	EditInfo(c *gin.Context)        //用户编辑信息
	Info(c *gin.Context)            //返回用户信息
	ForgotPassword(c *gin.Context)  //忘记密码
	VerifyResetCode(c *gin.Context) //验证验证码
	Schedule(c *gin.Context)        //首页行程推送
}

func NewUserController() InUserController {
	db := common.GetDB()
	db.Table("user").AutoMigrate(model.User{})
	//db.Table("admin").AutoMigrate(model.User{})
	return UserController{DB: db}
}

// 注册
func (a UserController) Register(c *gin.Context) { //未完成

	// 获取context中的参数
	var requestUser model.User
	// 将请求中的JSON数据绑定到User结构体中，方便后续操作
	c.Bind(&requestUser)
	userName := requestUser.UName
	userEmail := requestUser.UEmail
	password := requestUser.UKey
	//userremark := requestUser.URemark
	//userimage := requestUser.UImage //暂时是为空的！！！！！！！！！！！！都不需要传 来一个默认的

	//加个注册验证邮箱

	// 验证数据
	var user model.User
	a.DB.Table("user").Where("UEmail=?", userEmail).First(&user)
	if user.UID != 0 { // 该邮箱已存在，返回422，且返回提示信息
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
		//URemark: userremark,
		//UImage:  userimage,
	}
	a.DB.Table("user").Create(&newUser)
	//***********************************************************
	//密码加密功能
	//对于用户名要不要不让重复
	//默认头像图标11111111111111111111111111111111111111111111
	//************************************************************

	// 返回成功响应
	common.Success(c, gin.H{"newUser": newUser}, "注册成功")
}

// Login 登录
func (a UserController) Login(c *gin.Context) {

	// 获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	userEmail := requestUser.UEmail
	password := requestUser.UKey

	//fmt.Println("email:", userEmail, "  password:", password)
	// 数据验证
	var user model.User
	a.DB.Table("user").Where("UEmail=?", userEmail).First(&user)
	if user.UID == 0 {
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

	// 返回用户详细信息
	userInfo := gin.H{
		"UID":     user.UID,
		"UName":   user.UName,
		"UEmail":  user.UEmail,
		"UImage":  user.UImage,
		"URemark": user.URemark,
		"token":   token,
	}

	// 返回结果
	common.Success(c, gin.H{"userInfo": userInfo}, "登录成功")
}

// 关注行为添加到配对表
func (a UserController) GuanZhu(c *gin.Context) {
	var requestUser model.Gz
	c.Bind(&requestUser)
	gzid := requestUser.Gzid
	bgzid := requestUser.Bgzid

	// 判断 gzid 是否为 0
	if gzid == 0 {
		common.Fail(c, 400, nil, "输入Gzid 为零或为字符串")
		return
	}
	// 判断 bgzid 是否为 0
	if bgzid == 0 {
		common.Fail(c, 400, nil, "输入Bgzid 为零或为字符串")
		return
	}

	// 创建用户
	newMatch := model.Gz{
		Gzid:  gzid,
		Bgzid: bgzid,
	}
	a.DB.Table("gzmatch").Create(&newMatch)

	// 返回成功响应
	common.Success(c, nil, "关注配对成功")
}

func (a UserController) Logout(c *gin.Context) { //用户登出
	session := sessions.Default(c)
	session.Clear() // 清除所有会话数据
	session.Save()  // 保存更改
	common.Success(c, nil, "用户已登出")
}

// 返回粉丝数量  前端传入的数据是用户id 是可以只返回一个Uid的 用全局变量
func (a UserController) FansNum(c *gin.Context) {
	var UserID struct {
		Uid uint `json:"userID"`
	}
	c.Bind(&UserID)
	userid := UserID.Uid
	fmt.Println(UserID.Uid)
	fmt.Println(userid)
	// 数据验证
	var user model.User
	a.DB.Table("user").Where("uid=?", userid).First(&user)
	if user.UID == 0 {
		common.Fail(c, 422, nil, "用户不存在")
		return
	}

	// 查找 gzmatch 表中 bgzid 等于 userid 的条目数量
	var count int64
	a.DB.Table("gzmatch").Where("bgz_id=?", userid).Count(&count)

	// 返回结果
	common.Success(c, gin.H{"fansNum": count}, "获取粉丝数成功")
}

func (a UserController) EditInfo(c *gin.Context) { //修改未测试

	var requestUser model.User //改什么传什么 用json
	c.Bind(&requestUser)
	userid := requestUser.UID
	userName := requestUser.UName
	userEmail := requestUser.UEmail
	password := requestUser.UKey
	userRemark := requestUser.URemark
	userImage := requestUser.UImage

	// 数据验证
	var user model.User
	a.DB.Table("user").Where("uid=?", userid).First(&user)
	if user.UID == 0 {
		common.Fail(c, 422, nil, "用户不存在")
		return
	}

	// 更新用户信息
	updateData := model.User{
		UName:   userName,
		UEmail:  userEmail,
		UKey:    password,
		URemark: userRemark,
		UImage:  userImage,
	}
	a.DB.Table("user").Where("uid=?", userid).Updates(updateData)

	// 返回成功响应
	common.Success(c, nil, "用户信息更新成功")

}

func (a UserController) Info(c *gin.Context) { //返回用户信息
	var UserID struct {
		UID uint `json:"userID"`
	}
	c.Bind(&UserID)
	userid := UserID.UID

	if userid == 0 {
		common.Fail(c, 422, nil, "用户未登录")
		return
	}

	var user model.User
	a.DB.Table("user").Where("uid=?", userid).First(&user)
	// 更新用户信息
	infoOfUser := model.User{
		UID:     user.UID,
		UName:   user.UName,
		UEmail:  user.UEmail,
		UKey:    user.UKey,
		URemark: user.URemark,
		UImage:  user.UImage,
	}

	// 返回成功响应
	common.Success(c, gin.H{"user": infoOfUser}, "用户信息传递成功")
}

// ForgotPasswordHandler 处理忘记密码请求
func (a UserController) ForgotPassword(c *gin.Context) {
	var Message model.EmailMessage
	c.Bind(&Message)
	email := Message.Email
	// 如果电子邮件地址为空，返回400错误
	if email == "" {
		common.Fail(c, 400, gin.H{"error": "Email is required"}, "电子邮件地址为空")
		return
	}

	//调用GenerateRandomCode生成随机六位数
	code := utils.GenerateRandomCode(6)
	//调用SendVerificationCode函数发送验证码到用户的电子邮件。
	err := services.SendVerificationCode(email, code)
	if err != nil {
		common.Fail(c, 500, gin.H{"error": "Failed to send verification code"}, "服务器错误")
		return
	}
	//初始化redis客户端
	redisClient := redis.NewRedisClient()
	err = redisClient.SetCode(email, code, 5*time.Minute) // 设置验证码有效期为5分钟
	if err != nil {
		//存储验证码失败
		common.Fail(c, 500, gin.H{"error": "Failed to store verification code"}, "服务器错误")
		return
	}
	common.Success(c, gin.H{"message": "Verification code sent to your email."}, "成功发送验证码到邮箱")
}

// VerifyResetCodeHandler 验证重置密码的验证码
func (a UserController) VerifyResetCode(c *gin.Context) {
	var message model.EmailMessage
	c.Bind(&message)
	email := message.Email
	code := message.Code

	//fmt.Println("email: ", email, " code: ", code)

	redisClient := redis.NewRedisClient()
	storedCode, err := redisClient.GetCode(email)
	if err != nil {
		// 如果从Redis获取验证码失败，返回500错误
		common.Fail(c, 500, gin.H{"error": "Failed to retrieve verification code"}, "获取验证码失败")
		return
	}

	if storedCode != code {
		//fmt.Println("fail!------", storedCode, " : ", code)
		//验证码不匹配
		common.Fail(c, 400, gin.H{"error": "Invalid verification code"}, "验证码错误")
		return
	} else {
		// 如果验证码匹配，返回成功消息
		//fmt.Println("success!------", storedCode, " : ", code)
		common.Success(c, gin.H{"message": "Verification successful. You can reset your password now."}, "验证码匹配成功")
	}
}

func (a UserController) Schedule(c *gin.Context) { //返回用户行程

	var UserID struct {
		UID uint `json:"userID"`
	}
	c.Bind(&UserID)
	userid := UserID.UID

	if userid == 0 {
		common.Fail(c, 422, nil, "用户未登录")
		return
	}
	//查询加入的房间
	var cidList1 []struct {
		Cid uint `gorm:"column:cid"`
	}
	var cidList2 []struct {
		Cid uint `gorm:"column:cid"`
	}

	result := a.DB.Table("success_match").
		Select("cid").
		Where("uid = ?", userid).
		Find(&cidList1)
	if result.Error != nil {
		common.Fail(c, 500, nil, "数据库查询房间号错误")
		return
	}

	//查询为房主的房间信息
	result = a.DB.Table("chatinfo").
		Select("cid").
		Where("uid=?", userid).
		Where("cstate = ?", 1). //进行状态
		Find(&cidList2)
	if result.Error != nil {
		common.Fail(c, 500, nil, "数据库查询房间号错误")
		return
	}

	var cidList []uint

	//合并两组房间号
	for _, v := range cidList1 {
		cidList = append(cidList, v.Cid)
	}
	for _, v := range cidList2 {
		cidList = append(cidList, v.Cid)
	}

	if len(cidList) == 0 {
		common.Success(c, nil, "暂无行程")
		return
	}

	var ChatList []struct { //先取出版块id和日期
		Bid      int    `gorm:"column:bid"`
		CYueDate string `gorm:"column:c_yue_date"`
	}

	result = a.DB.Table("chatinfo").
		Select("bid, c_yue_date").
		Where("cstate = ?", 1). //进行状态
		Where("cid IN (?)", cidList).
		Find(&ChatList)
	if result.Error != nil {
		common.Fail(c, 500, nil, "数据库查询聊天室信息错误")
		return
	}
	if len(ChatList) == 0 {
		common.Success(c, nil, "暂无进行中的行程")
		return
	}

	var ScheduleList []model.ScheduleItem
	for _, chat := range ChatList {
		var Block struct {
			Bname string `gorm:"column:bname"`
		}
		result = a.DB.Table("block").
			Select("bname").
			Where("bid = ?", chat.Bid).
			First(&Block)
		if result.Error != nil {
			common.Fail(c, 500, nil, "数据库查询错误")
			return
		}

		scheduleItem := model.ScheduleItem{
			Date:    chat.CYueDate,
			Content: Block.Bname,
		}
		ScheduleList = append(ScheduleList, scheduleItem) //在上面房主的基础上加
	}
	common.Success(c, gin.H{"ScheduleList": ScheduleList}, "返回行程成功")
}
