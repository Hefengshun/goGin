package global

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	GinServer *gin.Engine
)
