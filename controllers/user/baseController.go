package user

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (_this *BaseController) success(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (_this *BaseController) error(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "error",
	})
}
