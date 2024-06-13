package routers

import (
	"ginDemo/controllers/user"
	"ginDemo/middlewares"
	"github.com/gin-gonic/gin"
)

// UserRouter 外部使用方法首字母必须大写公有方法，  内部使用方法首字母小写
func UserRouter(ginServer *gin.Engine) {
	userRouter := ginServer.Group("/user", middlewares.PrintOne)
	//userRouter := ginServer.Group("/user","xxx中间件")  方法1
	//userRouter.Use("xxx中间件","xxx中间件2")   方法2
	{
		userRouter.POST("/demo", user.UserController{}.Demo)
		userRouter.POST("/signup", user.UserController{}.SignUp)
		userRouter.POST("/login", user.UserController{}.Login)
		userRouter.POST("/logout", user.UserController{}.Logout)
	}
}
