package system

import (
	"ginDemo/controllers/system"
	"github.com/gin-gonic/gin"
)

type MassageRouter struct{}

func (m *MassageRouter) InitMassageRouter(ginServer *gin.Engine) {

	MassageController := &system.MassageController{}
	massageRouter := ginServer.Group("/massage")
	{
		massageRouter.POST("/createConversation", MassageController.CreateConversation)
		massageRouter.POST("/getConversationsWithUnreadCount", MassageController.GetConversationsWithUnreadCount)
		massageRouter.POST("/getMessagesForConversation", MassageController.GetMessagesForConversation)
		massageRouter.POST("/sendMessage", MassageController.SendMessage)
	}
}
