package system

import (
	"ginDemo/controllers/user"
	"ginDemo/middlewares"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

// InitUserRouter 外部使用方法首字母必须大写公有方法，  内部使用方法首字母小写
func (_this *UserRouter) InitUserRouter(ginServer *gin.Engine) {
	userController := &user.UserController{} //因为是指针接受者的方法 所有使用时要&引用地址 引用的地址上面挂载了这个方法
	//uc.Demo ｜｜ (&user.UserController{}).Demo // 这可以正常工作
	userRouter := ginServer.Group("/user", middlewares.PrintOne)
	//userRouter := ginServer.Group("/user","xxx中间件")  方法1
	//userRouter.Use("xxx中间件","xxx中间件2")   方法2
	{
		userRouter.POST("/demo", userController.Demo)
		userRouter.POST("/signup", userController.SignUp)
		userRouter.POST("/login", userController.Login)
		userRouter.POST("/logout", userController.Logout)
	}
}
