package routers

import (
	"ginDemo/controllers/user"
	"github.com/gin-gonic/gin"
)

// UserRouter 外部使用方法首字母必须大写公有方法，  内部使用方法首字母小写
func UserRouter(ginServer *gin.Engine) {
	userRouter := ginServer.Group("/user")
	{
		userRouter.POST("/signup", user.SignUp)
		userRouter.POST("/login", user.Login)
		userRouter.POST("/logout", user.Logout)
	}
}
