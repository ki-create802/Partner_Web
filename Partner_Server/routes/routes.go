package routes

import (
	"Partner_Web/Partner_Server/controller"
	//"GolandProject/routes/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	//// 允许跨域访问 还没学明白
	//r.Use(middleware.CORSMiddleware())

	// 用户操作
	userController := controller.NewUserController()
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("register", userController.Register)                 // 注册
		userRoutes.POST("login", userController.Login)                       // 登录
		userRoutes.POST("logout", userController.Logout)                     //登出
		userRoutes.POST("guanzhu", userController.GuanZhu)                   //关注用户
		userRoutes.POST("fansNum", userController.FansNum)                   //粉丝数量
		userRoutes.PATCH("editInfo", userController.EditInfo)                //用户编辑信息
		userRoutes.POST("info", userController.Info)                         //返回用户信息
		userRoutes.POST("forgot-password", userController.ForgotPassword)    //忘记密码
		userRoutes.POST("verify-reset-code", userController.VerifyResetCode) //验证验证码
		userRoutes.POST("schedule", userController.Schedule)                 //推送行程
		//userRoutes.POST("addReview", userController.AddReview)               // 添加评分
		//userRoutes.POST("getUserReviews", userController.GetUserReviews)     // 获取用户评分
		userRoutes.POST("getUserLevel", userController.GetUserLevel) // 获取用户等级称号
	}

	chatController := controller.NewChatController()
	chatRoutes := r.Group("/chat")
	{
		chatRoutes.POST("searchChatList", chatController.SearchChatList)   //返回聊天室匹配结果
		chatRoutes.POST("searchUser", chatController.SearchUser)           //返回用户匹配结果
		chatRoutes.POST("createChat", chatController.CreateChat)           //创建聊天室
		chatRoutes.POST("chatLists", chatController.ChatLists)             //返回聊天室界面的聊天室列表
		chatRoutes.POST("chatRecords", chatController.ChatRecords)         //返回聊天室id对应聊天记录
		chatRoutes.POST("saveChatRecords", chatController.SaveChatRecords) //保存前端聊天记录

		chatRoutes.POST("disbandChatRoom", chatController.DisbandChatRoom) // 房主解散群聊
		chatRoutes.POST("leaveChatRoom", chatController.LeaveChatRoom)     //退出群聊
		chatRoutes.POST("addMember", chatController.AddMember)             //添加成员到uc_match表中
		chatRoutes.POST("getSpeakers", chatController.GetSpeakers)         //返回发言成员列表（uc_match信息）
		chatRoutes.POST("successMatch", chatController.SuccessMatch)       //添加成功匹配成员信息

		chatRoutes.POST("getPendingChats", chatController.GetPendingChats) // 获取主页'等待中'板块聊天室列表
		chatRoutes.POST("getHistoryChats", chatController.GetHistoryChats) // 获取主页'历史'板块聊天室列表

	}

	fileController := controller.NewFileController()
	fileRoutes := r.Group("/file")
	{
		fileRoutes.POST("uploadAvatar", fileController.UploadAvatar) // 上传头像
	}
	return r
}
