package demo

import (
	"ginDemo/global"
	"ginDemo/models/demo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SqlController struct{}

func (_this *SqlController) ReturnOneForm(c *gin.Context) {
	userList := []demo.SysDemo{}
	global.DB.Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"api":     c.Request.RequestURI,
		"data":    userList,
	})
}

func (_this *SqlController) CreateOneData(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	//user := demo.SysDemo{
	//	Name:     name,
	//	Password: password,
	//}
	//ok := global.DB.Create(&user)
	ok := global.DB.Model(&demo.SysDemo{}).Create(map[string]interface{}{
		"Name":     name,
		"Password": password,
	})
	if ok.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "CreateOneData  success",
			"api":     c.Request.RequestURI,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "CreateOneData  error",
			"api":     c.Request.RequestURI,
			"error":   ok.Error,
		})
	}

}

func (_this *SqlController) DeleteOneData(c *gin.Context) {
	userList := []demo.SysDemo{}
	global.DB.Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"api":     c.Request.RequestURI,
		"data":    userList,
	})
}

func (_this *SqlController) SelectIdData(c *gin.Context) {
	// 获取查询参数并转换为 uint64
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Invalid ID"})
		return
	}

	// 创建 SysDemo 实例并查询数据库
	//user := demo.SysDemo{}
	//result := global.DB.Find(&user)

	user := map[string]interface{}{
		"id": id,
	}
	result := global.DB.Model(&demo.SysDemo{}).Find(&user)
	//result := global.DB.Model(demo.SysDemo{}).Select("id").Where("id = ?", id)

	// 检查是否有错误
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": result.Error.Error()})
		return
	}

	// 检查是否找到了记录
	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Record not found",
			"api":     c.Request.RequestURI,
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "SelectOneData success",
		"api":     c.Request.RequestURI,
		"data":    user,
	})

}

func (_this *SqlController) SelectNameData(c *gin.Context) {
	name := c.Query("name")
	usersList := []demo.SysDemo{}
	result := global.DB.Where("name = ?", name).Find(&usersList)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "SelectNameData success",
			"api":     c.Request.RequestURI,
			"data":    usersList,
		})
	}
}

func (_this *SqlController) SelectKeyToArray(c *gin.Context) {
	name := c.PostFormArray("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "SelectNameData success",
		"api":     c.Request.RequestURI,
		"data":    name,
	})
}
func (_this *SqlController) Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
}
