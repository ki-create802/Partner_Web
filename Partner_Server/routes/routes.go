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
		userRoutes.GET("fansNum", userController.FansNum)                    //粉丝数量
		userRoutes.PATCH("editInfo", userController.EditInfo)                //用户编辑信息
		userRoutes.GET("info", userController.Info)                          //返回用户信息
		userRoutes.POST("forgot-password", userController.ForgotPassword)    //忘记密码
		userRoutes.POST("verify-reset-code", userController.VerifyResetCode) //验证验证码
	}

	//chatController := controller.NewChatController()
	//chatRoutes := r.Group("/chat")
	//{
	//	chatRoutes.GET("searchChatList", chatController.SearchChatList) //返回聊天室匹配结果
	//	chatRoutes.GET("searchUser", chatController.SearchUser)         //返回用户匹配结果
	//}

	return r
}
