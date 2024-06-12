package main

import (
	"ginDemo/global"
	"ginDemo/initialize"
	"ginDemo/routers" //这里相当于自己定义了一个包 因为里面的方法名大写 所以可以根据这个包获取里面的公用方法
	"github.com/gin-gonic/gin"
)

func main() {
	global.DB = initialize.InitDB()
	// 切换到 release 模式
	gin.SetMode(gin.ReleaseMode)
	ginServer := gin.Default()
	// 设置受信任的代理（根据需要）
	ginServer.SetTrustedProxies([]string{"0.0.0.0"})
	routers.DefaultRouter(ginServer)

	routers.UserRouter(ginServer)
	// 正确的监听端口配置
	if err := ginServer.Run(":8088"); err != nil {
		panic(err)
	}
}
