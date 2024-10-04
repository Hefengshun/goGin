package initialize

import (
	"ginDemo/global"
	"ginDemo/middlewares"
	"ginDemo/routers"
)

func InitRouters() {
	//ginServer.Use(xxx,xxx)里面可以配置多个全局中间件 当然路由组的中间件也可以配置
	//**** 中间件如果使用了携程 多线程 go XXX 里面的c *gin.Context 要拷贝使用
	global.GinServer.Use(middlewares.Cors())
	global.GinServer.Use(middlewares.HandleRouter)
	global.GinServer.Use(middlewares.JWTAuth()) //注册了鉴权中间键

	//routers.DemoRouter(global.GinServer)
	//routers.UserRouter(global.GinServer)
	//routers.UploadRouter(global.GinServer)

	systemRouter := routers.RouterGroupApp.System
	cesRouter := routers.RouterGroupApp.Ces

	//拿到两组app后组册两组里面所有的组
	cesRouter.InitDemoRouter(global.GinServer)
	cesRouter.InitSqlDemo(global.GinServer)
	systemRouter.InitUploadRouter(global.GinServer)
	systemRouter.InitMassageRouter(global.GinServer)
	systemRouter.InitUserRouter(global.GinServer)
}
