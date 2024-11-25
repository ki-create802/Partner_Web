package main

import (
	"Partner_Web/Partner_Server/common"
	"Partner_Web/Partner_Server/routes"
	"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/gin"
	//"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	////创建一个服务
	//ginServer := gin.Default()
	//
	//ginServer.Use()
	////加载静态页面
	//ginServer.LoadHTMLGlob("WebPage/*")
	//
	////访问地址，处理请求
	//ginServer.GET("/login", func(context *gin.Context) { //get请求
	//	context.HTML(http.StatusOK, "user.html", gin.H{
	//		"msg": "后台传递的数据"})
	//})
	////ginServer.POST()
	////给一个服务器端口 制定访问地址 通过访问8082来接受这个请求
	//ginServer.Run(":8082")

	// 获取初始化的数据库
	db := common.InitDB("run")
	// 延迟关闭数据库
	defer db.Close()

	//创建路由
	r := gin.Default()
	// 启动路由
	routes.CollectRoutes(r)

	// 启动服务
	panic(r.Run(":8082"))
}
