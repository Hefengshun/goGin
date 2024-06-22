package system

import (
	"ginDemo/controllers/system"
	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

func (_this *UploadRouter) InitUploadRouter(ginServer *gin.Engine) {
	uploadRouter := ginServer.Group("/api")
	{
		uploadRouter.POST("/unifile", system.UploadController{}.UniFile)
		uploadRouter.POST("/multifile", system.UploadController{}.MultiFile)
	}
}
