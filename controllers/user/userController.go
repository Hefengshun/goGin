package user

import "github.com/gin-gonic/gin"

//func SignUp(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"message": "Sign Up",
//	})
//}
//
//func Login(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"message": "Login User",
//	})
//}
//
//func Logout(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"message": "Logout User",
//	})
//}
//以上相当于单个抛出的函数

type UserController struct {
	BaseController //使用结构体实现控制器的继承
}

func (_this *UserController) Demo(c *gin.Context) {
	name, ok := c.Get("name") //这里忽略了name的空接口类型
	if ok {
		c.JSON(200, gin.H{
			"name": name,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Logout User",
		})
	}

}

func (_this *UserController) SignUp(c *gin.Context) {
	_this.success(c)
}

func (_this *UserController) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login User",
	})
}

func (_this *UserController) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Logout User",
	})
}
