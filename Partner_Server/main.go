package main

import (
	"Partner_Web/Partner_Server/common"
	"Partner_Web/Partner_Server/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/gin"
	//"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 获取初始化的数据库
	db := common.InitDB("run")
	// 延迟关闭数据库
	defer db.Close()

	//初始化会话存储
	store := cookie.NewStore([]byte("secret"))

	//store.Options(sessions.Options{
	//	Path:     "/",
	//	MaxAge:   -1,
	//	HttpOnly: true,
	//})
	//创建路由
	r := gin.Default()
	//注册回会话中间件
	r.Use(sessions.Sessions("mysession", store))
	// 启动路由
	routes.CollectRoutes(r)

	// 启动服务
	panic(r.Run(":8082"))
}
