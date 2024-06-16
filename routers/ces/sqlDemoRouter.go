package ces

import (
	"ginDemo/controllers/demo"
	"github.com/gin-gonic/gin"
)

type SqlDemoRouter struct{}

func (_this *SqlDemoRouter) InitSqlDemo(ginServer *gin.Engine) {
	demoControllers := &demo.DemoControllers{}
	//sqlDemoRouter := ginServer.Group("/sqlDemo")  路径区分大小写 ！！！！
	sqlDemoRouter := ginServer.Group("/demo")

	{
		sqlDemoRouter.GET("returnoneform", (&demo.SqlController{}).ReturnOneForm)
		sqlDemoRouter.GET("/createonedata", demoControllers.CreateOneData)
		sqlDemoRouter.GET("/deleteonedata", demoControllers.DeleteOneData)
		sqlDemoRouter.GET("/selectiddata", demoControllers.SelectIdData)
		sqlDemoRouter.POST("/login", demoControllers.LoginDemo)

	}
}
