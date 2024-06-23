package demo

import (
	"ginDemo/global"
	"ginDemo/models/demo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

type Login struct {
	Name     string `json:"name"`                        // 用户名
	Password string `json:"password" binding:"required"` // 密码 结构体绑定校验
}

func (_this *UserController) LoginDemo(c *gin.Context) {
	user := demo.SysDemo{}
	reqUser := new(Login)

	if err := c.ShouldBind(reqUser); err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"url":  c.Request.RequestURI,
			"data": "密码未填写！",
			"err":  err.Error(),
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
