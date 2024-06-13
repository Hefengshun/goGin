package user

import "github.com/gin-gonic/gin"

func SignUp(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Sign Up",
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login User",
	})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Logout User",
	})
}
