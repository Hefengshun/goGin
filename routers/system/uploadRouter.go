package system

import (
	"ginDemo/controllers/upload"
	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

func (_this *UploadRouter) InitUploadRouter(ginServer *gin.Engine) {
	uploadRouter := ginServer.Group("/api")
	{
		uploadRouter.POST("/unifile", upload.UploadController{}.Unifile)
		uploadRouter.POST("/multifile", upload.UploadController{}.Multifile)
	}
}
