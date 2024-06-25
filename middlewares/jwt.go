package middlewares

import (
	"fmt"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		for i := range utils.NoVerify {
			fmt.Println(utils.NoVerify[i], c.Request.RequestURI)
			if utils.NoVerify[i] == c.Request.RequestURI {
				return
			}
		}
		tokenString := c.GetHeader("Authorization")

		fmt.Print("请求token", tokenString)

		//验证前端传过来的token格式，不为空，开头为Bearer
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(401, gin.H{
				"data": gin.H{},
				"meta": gin.H{
					"msg":  "权限不足",
					"code": 401,
				},
			})
			c.Abort()
		}

		//验证通过，提取有效部分（除去Bearer)
		tokenString = tokenString[7:] //截取字符
		token, claims, err := utils.ParseToken(tokenString)
		//解析失败||解析后的token无效
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{
				"data": gin.H{},
				"meta": gin.H{
					"msg":  "权限不足",
					"code": 401,
				},
			})
			c.Abort()
		}

		//token通过验证, 获取claims中的UserID
		UserName := claims.UserName
		fmt.Println(UserName)

		c.Next()
	}
}
