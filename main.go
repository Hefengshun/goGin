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
	ginServer := gin.Default() //这里面已经加了 Logger(), Recovery() 默认是中间件的F12进去看 不想用就engine := New()手动创建路由引擎
	// 设置受信任的代理（根据需要）
	ginServer.SetTrustedProxies([]string{"0.0.0.0"})
	//ginServer.Use(xxx,xxx)里面可以配置多个全局中间件 当然路由组的中间件也可以配置
	//**** 中间件如果使用了携程 多线程 go XXX 里面的c *gin.Context 要拷贝使用
	routers.DemoRouter(ginServer)
	routers.UserRouter(ginServer)
	// 正确的监听端口配置
	if err := ginServer.Run(":8088"); err != nil {
		panic(err)
	}
}
