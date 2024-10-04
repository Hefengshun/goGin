package system

import (
	"fmt"
	"ginDemo/models/common/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MassageController struct {
}

// CreateConversation
// 创建会话
func (m *MassageController) CreateConversation(c *gin.Context) {
	userOpenid := c.PostForm("userOpenid")
	friendOpenid := c.PostForm("friendOpenid")

	conversationID, exist, err := massageService.FindOrCreateConversation(userOpenid, friendOpenid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if exist {
		response.OkWithDetailed(map[string]interface{}{
			"conversationID": conversationID,
			"exist":          exist,
		}, "打开会话", c)
		return
	}
	response.OkWithDetailed(map[string]interface{}{
		"conversationID": conversationID,
		"exist":          exist,
	}, "创建了一个对话", c)
}

// GetConversationsWithUnreadCount
// 获取会话列表及未读消息数量
func (m *MassageController) GetConversationsWithUnreadCount(c *gin.Context) {
	userID := c.PostForm("userID")
	ConversationsData, err := massageService.GetConversationsWithUnreadCount(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(ConversationsData, "返回成功", c)
}

// GetMessagesForConversation
// 获取会话中的消息记录
func (m *MassageController) GetMessagesForConversation(c *gin.Context) {
	num64, err := strconv.ParseUint(c.PostForm("conversationID"), 10, 32)
	if err != nil {
		fmt.Println("转换错误:", err)
		return
	}
	conversationID := uint(num64)
	userID := c.PostForm("userID")
	massageList, err := massageService.GetMessagesForConversation(conversationID, userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(massageList, "返回消息记录", c)
}

// SendMessage
// 发送消息并更新未读状态
func (m *MassageController) SendMessage(c *gin.Context) {
	num64, err := strconv.ParseUint(c.PostForm("conversationID"), 10, 32)
	if err != nil {
		fmt.Println("转换错误:", err)
		return
	}
	conversationID := uint(num64)
	senderID := c.PostForm("senderID")
	content := c.PostForm("content")

	massage, sendErr := massageService.SendMessage(conversationID, senderID, content)
	if sendErr != nil {
		response.FailWithMessage(sendErr.Error(), c)
		return
	}
	response.OkWithDetailed(map[string]string{}, massage, c)

}
