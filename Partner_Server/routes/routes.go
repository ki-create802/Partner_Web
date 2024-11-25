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
		userRoutes.POST("register", userController.Register)  // 注册
		userRoutes.POST("login", userController.Login)        // 登录
		userRoutes.POST("guanzhu", userController.GuanZhu)    //关注用户
		userRoutes.GET("fansNum", userController.FansNum)     //粉丝数量
		userRoutes.PATCH("editInfo", userController.EditInfo) //用户编辑信息
	}

	return r
}
