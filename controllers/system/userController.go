package system

import (
	"ginDemo/global"
	"ginDemo/models/demo"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

type UserController struct{}

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
	c.JSON(200, gin.H{
		"message": "Login User",
	})
}

func (_this *UserController) Login(c *gin.Context) {
	type Login struct {
		Name     string `json:"name"`                        // 用户名
		Password string `json:"password" binding:"required"` // 密码 结构体绑定校验
	}
	user := demo.SysDemo{}
	reqUser := new(Login)
	//userService.Login(&reqUser)
	if err := c.ShouldBind(reqUser); err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"url":  c.Request.RequestURI,
			"data": "密码未填写！",
		})
		return
	}
	c.ShouldBind(&reqUser) //c.ShouldBind 使用了 c.Request.Body json数据 绑定到结构体里面 但是不可绑定多个结构体多次绑定使用c.ShouldBindBodyWith
	result := global.DB.Where("name = ? and password = ?", reqUser.Name, reqUser.Password).Find(&user)

	/*
		name := c.PostForm("name")         //取出的是form-data数据
		password := c.PostForm("password") //取出的是form-data数据
		result := global.DB.Where("name = ? AND password = ?", name, password).Find(&user)
	*/

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"url":  c.Request.RequestURI,
			"data": "登录失败！未查找到数据",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"url":  c.Request.RequestURI,
		"data": user,
	})
}

func (_this *UserController) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Logout User",
	})
}
