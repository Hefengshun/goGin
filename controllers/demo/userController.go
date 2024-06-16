package demo

import (
	"ginDemo/global"
	"ginDemo/models/demo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

type Login struct {
	Name     string `json:"name"`     // 用户名
	Password string `json:"password"` // 密码
}

func (_this *UserController) LoginDemo(c *gin.Context) {

	//name := c.PostForm("name")         //取出的是form-data数据
	//password := c.PostForm("password") //取出的是form-data数据
	userForm := demo.SysDemo{}
	var user Login
	c.ShouldBind(&user) //取出的是 json数据
	//result, err := c.MultipartForm() //取出的是form-data数据
	//if err != nil {
	//	fmt.Println(result)
	//}
	global.DB.Where("name = ? AND password = ?", userForm.Name, userForm.Password).First(&userForm)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"url":  c.Request.RequestURI,
		"data": user,
	})
}
