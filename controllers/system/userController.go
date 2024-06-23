package system

import (
	"fmt"
	"ginDemo/models/common/response"
	"ginDemo/models/system"
	stytemReq "ginDemo/models/system/request"
	stytemRes "ginDemo/models/system/response"
	"github.com/gin-gonic/gin"
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

func (_this *UserController) SignUp(c *gin.Context) {
	var l stytemReq.SignUp
	err := c.ShouldBindJSON(&l)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Login User",
			"data":    l,
			"err":     err.Error(),
		})
		return
	}
	user := &system.SysUser{
		Username: l.Name,
		Password: l.Password,
	}
	userInfo, err := userService.SignUp(*user)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Login User",
			"err":     err.Error(),
			"data":    userInfo,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Login User",
			"data":    userInfo,
			"success": "注册成功",
		})
	}
}

func (_this *UserController) Login(c *gin.Context) {
	var reqUser stytemReq.Login
	if err := c.ShouldBind(&reqUser); err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"url":  c.Request.RequestURI,
			"err":  err.Error(),
		})
		return
	}
	user := &system.SysUser{
		Username: reqUser.Name,
		Password: reqUser.Password,
	}
	// 打印指针
	fmt.Printf("user pointer: %p\n", user)

	// 打印指针指向的值
	fmt.Printf("user value: %+v\n", user)

	// 直接打印指针，会自动解引用
	fmt.Println("user:", user)

	userInfo, err := userService.Login(user)
	//fmt.Println(userInfo, *userInfo)

	if err != nil {
		response.FailWithMessage("用户不存在或者密码错误", c)
		return
	}
	returnUserInfo := &stytemRes.Login{User: *userInfo, Token: "token"} // *userInfo 之所以要传这个 因为那边接受的是值

	response.OkWithDetailed(returnUserInfo, "登录成功！", c)
}

func (_this *UserController) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Logout User",
	})
}
