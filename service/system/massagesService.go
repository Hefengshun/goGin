package system

import (
	"ginDemo/global"
	"ginDemo/models/system"
	"ginDemo/models/system/response"
)

type MassageService struct {
}

// FindOrCreateConversation
// 查找会话
func (m *MassageService) FindOrCreateConversation(userID string, friendID string) (uint, bool, error) {
	var conversationID uint
	// 检查是否已经存在会话
	err := global.DB.Table("sys_conversations").
		Joins("JOIN sys_conversation_members scm1 ON sys_conversations.id = scm1.conversation_id").
		Joins("JOIN sys_conversation_members scm2 ON sys_conversations.id = scm2.conversation_id").
		Where("scm1.user_id = ? AND scm2.user_id = ?", userID, friendID).
		Select("sys_conversations.id").
		Scan(&conversationID).Error

	// 如果找到了会话，返回会话ID
	if err == nil && conversationID > 0 {
		return conversationID, true, nil
	}

	// 如果没有找到会话，则创建新的会话
	conversation := system.SysConversations{}
	if err := global.DB.Create(&conversation).Error; err != nil {
		return 0, false, err
	}

	// 创建会话成员记录
	members := []system.SysConversationMembers{
		{ConversationID: conversation.ID, UserID: userID},
		{ConversationID: conversation.ID, UserID: friendID},
	}

	if err := global.DB.Create(&members).Error; err != nil {
		return 0, false, err
	}

	return conversation.ID, false, nil
}

// GetConversationsWithUnreadCount
// 获取会话列表及未读消息数量
func (m *MassageService) GetConversationsWithUnreadCount(userID string) ([]response.ConversationWithUnreadCount, error) {
	var result []response.ConversationWithUnreadCount
	err := global.DB.Table("sys_conversations").
		Select("sys_conversations.id as conversation_id, sys_conversations.last_message,sys_conversations.updated_at, COUNT(sys_message_status.message_id) as unread_count").
		Joins("JOIN sys_conversation_members ON sys_conversations.id = sys_conversation_members.conversation_id").
		Joins("LEFT JOIN sys_messages ON sys_conversations.id = sys_messages.conversation_id").
		Joins("LEFT JOIN sys_message_status ON sys_messages.id = sys_message_status.message_id AND sys_message_status.user_id = ?", userID).
		Where("sys_conversation_members.user_id = ?", userID).
		Where("sys_message_status.status = 'unread' OR sys_message_status.status IS NULL").
		Group("sys_conversations.id").
		Find(&result).Error

	if err != nil {
		return nil, err
	}

	// 收集会话 ID
	var conversationIDs []uint
	for _, conversation := range result {
		conversationIDs = append(conversationIDs, conversation.ConversationID)
	}

	// 查询会话成员
	var members []struct {
		ConversationID uint
		UserID         string
	}
	err = global.DB.Table("sys_conversation_members").
		Select("conversation_id, user_id").
		Where("conversation_id IN (?)", conversationIDs).
		Find(&members).Error

	if err != nil {
		return nil, err
	}

	// 收集用户 ID
	userIDSet := make(map[string]bool)
	for _, member := range members {
		if member.UserID != userID { // 过滤掉当前用户
			userIDSet[member.UserID] = true
		}
	}

	// 将用户 ID 转换为切片
	var userIDs []string
	for id := range userIDSet {
		userIDs = append(userIDs, id)
	}

	// 批量查询用户信息
	var users []system.SysWxUser
	err = global.DB.Where("openid IN (?)", userIDs).Find(&users).Error
	if err != nil {
		return nil, err
	}

	// 创建用户信息映射
	userMap := make(map[string]system.SysWxUser)
	for _, user := range users {
		userMap[user.Openid] = user
	}

	// 组装结果
	for i, conversation := range result {
		for _, member := range members {
			if member.ConversationID == conversation.ConversationID && member.UserID != userID { // 过滤掉当前用户
				if user, exists := userMap[member.UserID]; exists {
					result[i].FriendOpenid = user.Openid
					result[i].FriendName = user.UserName
				}
			}
		}

	}

	return result, nil
}

// GetMessagesForConversation
// 获取会话中的消息记录
func (m *MassageService) GetMessagesForConversation(conversationID uint, userID string) ([]system.SysMessages, error) {
	var messages []system.SysMessages
	// 获取最近的消息记录
	err := global.DB.Where("conversation_id = ?", conversationID).Order("created_at").Limit(20).Find(&messages).Error
	if err != nil {
		return nil, err
	}

	// 将自己未读消息标记为已读
	err = global.DB.Model(&system.SysMessageStatus{}).
		Where("message_id IN (?) AND user_id = ? AND status = 'unread'",
			global.DB.Table("sys_messages").
				Select("id").
				Where("conversation_id = ?", conversationID),
			userID).
		Update("status", "read").Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// SendMessage
// 发送消息并更新未读状态
func (m *MassageService) SendMessage(conversationID uint, senderID string, content string) (string, error) {
	// 创建新消息
	message := system.SysMessages{
		ConversationID: conversationID,
		SenderID:       senderID,
		Content:        content,
	}
	err := global.DB.Create(&message).Error
	if err != nil {
		return "添加数据失败了！", err
	}

	// 更新会话的最后一条消息
	err = global.DB.Model(&system.SysConversations{}).Where("id = ?", conversationID).
		Update("last_message", content).Error
	if err != nil {
		return "更新会话最后一条数据失败了！", err
	}

	// 获取会话中的所有成员
	var members []system.SysConversationMembers
	global.DB.Where("conversation_id = ?", conversationID).Find(&members)

	// 为每个会话成员插入未读状态
	for _, member := range members {
		if member.UserID != senderID {
			err := global.DB.Create(&system.SysMessageStatus{
				MessageID: message.ID,
				UserID:    member.UserID,
				Status:    "unread",
			}).Error
			if err != nil {
				return "会话成员插入未读状态失败！", err
			}
		}
	}

	return "数据发送成功！", nil
}
