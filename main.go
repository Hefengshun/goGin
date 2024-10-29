package main

import (
	"ginDemo/global"
	"ginDemo/initialize"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8089" // Default port if not specified
	}

	global.DB = initialize.InitDB()
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 切换到 release 模式
	gin.SetMode(gin.ReleaseMode)
	ginServer := gin.Default() //这里面已经加了 Logger(), Recovery() 默认是中间件的F12进去看 不想用就engine := New()手动创建路由引擎
	// 设置受信任的代理（根据需要）
	ginServer.SetTrustedProxies([]string{"0.0.0.0"})
	// 提供静态文件服务
	ginServer.Static("/static", "./static")
	global.GinServer = ginServer
	initialize.InitRouters()
	// 正确的监听端口配置
	if err := ginServer.Run(":" + port); err != nil {
		panic(err)
	}
}
