package main

import (
	"Partner_Web/Partner_Server/common"
	"Partner_Web/Partner_Server/routes"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/gin"
	//"net/http"
	_ "github.com/go-sql-driver/mysql"
)

// func main() {
// 	//定义Redis服务器路径
// 	redisServerPath := "../Redis-x64-3.2.100/redis-server.exe"
// 	// redisServerPath := "/opt/homebrew/bin/redis-server" //mac系统

// 	//启动Redis服务器
// 	redisServerCmd := exec.Command(redisServerPath)
// 	//绑定标准输出和标准错误重定向
// 	redisServerCmd.Stdout = os.Stdout
// 	redisServerCmd.Stderr = os.Stderr

// 	// 使用 context 来管理 Redis 服务器的生命周期
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel() // 确保在程序结束时取消 context

// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	// 启动 Redis 服务器
// 	go func() {
// 		defer wg.Done()
// 		err := redisServerCmd.Start()
// 		if err != nil {
// 			fmt.Printf("Failed to start Redis server: %v\n", err)
// 			return
// 		}
// 		go func() {
// 			<-ctx.Done() // 等待 context 被取消
// 			redisServerCmd.Process.Kill()
// 		}()
// 		err = redisServerCmd.Wait()
// 		if err != nil {
// 			fmt.Printf("Redis server exited with error: %v\n", err)
// 		}
// 	}()

// 	// 等待 Redis 服务器启动
// 	time.Sleep(2 * time.Second)

// 	//定义Redis客户端路径
// 	redisClientPath := "../Redis-x64-3.2.100/redis-cli.exe"
// 	// redisClientPath := "/opt/homebrew/bin/redis-cli"
// 	redisClientCmd := exec.Command(redisClientPath, "ping")
// 	redisClientCmd.Stdout = os.Stdout
// 	redisClientCmd.Stderr = os.Stderr

// 	err := redisClientCmd.Run()
// 	if err != nil {
// 		fmt.Printf("Failed to run Redis client: %v\n", err)
// 		return
// 	}
// 	// 等待 Redis 客户端执行完毕
// 	time.Sleep(1 * time.Second)

// 	// 获取初始化的数据库
// 	db := common.InitDB("run")
// 	// 延迟关闭数据库
// 	defer db.Close()

// 	// //初始化会话存储
// 	// store := cookie.NewStore([]byte("secret"))

// 	//创建路由
// 	r := gin.Default()

// 	r.Use(cors.Default())

// 	// //注册回会话中间件
// 	// r.Use(sessions.Sessions("mysession", store))
// 	// 启动路由
// 	routes.CollectRoutes(r)

// 	// 配置静态文件服务(上传头像)
// 	r.Static("/avatars", "./avatars")

// 	//// 启动服务
// 	//panic(r.Run(":8082"))

// 	// 创建 HTTP 服务器
// 	srv := &http.Server{
// 		Addr:    ":8082",
// 		Handler: r,
// 	}

// 	// 启动 HTTP 服务器
// 	go func() {
// 		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			fmt.Printf("Failed to start HTTP server: %v\n", err)
// 		}
// 	}()

// 	// 等待 Web 服务器关闭
// 	quit := make(chan os.Signal, 1)
// 	// 监听中断信号
// 	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 	<-quit

// 	// 关闭 Web 服务器
// 	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	if err := srv.Shutdown(ctx); err != nil {
// 		fmt.Printf("Failed to shutdown HTTP server: %v\n", err)
// 	}

// 	// 取消 context 以关闭 Redis 服务器
// 	cancel()

// 	// 等待 Redis 服务器关闭
// 	wg.Wait()

// }

func main() {
	//定义Redis服务器路径
	redisServerPath := "../Redis-x64-3.2.100/redis-server.exe"

	//启动Redis服务器
	redisServerCmd := exec.Command(redisServerPath)
	//绑定标准输出和标准错误重定向
	redisServerCmd.Stdout = os.Stdout
	redisServerCmd.Stderr = os.Stderr

	// 使用 context 来管理 Redis 服务器的生命周期
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 确保在程序结束时取消 context

	var wg sync.WaitGroup
	wg.Add(1)

	// 启动 Redis 服务器
	go func() {
		defer wg.Done()
		err := redisServerCmd.Start()
		if err != nil {
			fmt.Printf("Failed to start Redis server: %v\n", err)
			return
		}
		go func() {
			<-ctx.Done() // 等待 context 被取消
			redisServerCmd.Process.Kill()
		}()
		err = redisServerCmd.Wait()
		if err != nil {
			fmt.Printf("Redis server exited with error: %v\n", err)
		}
	}()

	// 等待 Redis 服务器启动
	time.Sleep(2 * time.Second)

	//定义Redis客户端路径
	redisClientPath := "../Redis-x64-3.2.100/redis-cli.exe"
	redisClientCmd := exec.Command(redisClientPath, "ping")
	redisClientCmd.Stdout = os.Stdout
	redisClientCmd.Stderr = os.Stderr

	err := redisClientCmd.Run()
	if err != nil {
		fmt.Printf("Failed to run Redis client: %v\n", err)
		return
	}
	// 等待 Redis 客户端执行完毕
	time.Sleep(1 * time.Second)

	// 获取初始化的数据库
	db := common.InitDB("run")
	// 延迟关闭数据库
	defer db.Close()

	// //初始化会话存储
	// store := cookie.NewStore([]byte("secret"))

	//创建路由
	r := gin.Default()

	r.Use(cors.Default())

	// //注册回会话中间件
	// r.Use(sessions.Sessions("mysession", store))
	// 启动路由
	routes.CollectRoutes(r)

	// 配置静态文件服务(上传头像)
	r.Static("/avatars", "./avatars")

	//// 启动服务
	//panic(r.Run(":8082"))

	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:    ":8082",
		Handler: r,
	}

	// 启动 HTTP 服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Failed to start HTTP server: %v\n", err)
		}
	}()

	// 等待 Web 服务器关闭
	quit := make(chan os.Signal, 1)
	// 监听中断信号
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 	// 关闭 Web 服务器
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Failed to shutdown HTTP server: %v\n", err)
	}

	// 取消 context 以关闭 Redis 服务器
	cancel()

	// 等待 Redis 服务器关闭
	wg.Wait()

}
