package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func HandleRouter(c *gin.Context) {
	c.Request.URL.Path = strings.ToLower(c.Request.URL.Path)
	fmt.Println(c.Request.URL.Path)
	c.Next()
}
