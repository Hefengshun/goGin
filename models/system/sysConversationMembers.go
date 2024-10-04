package system

type SysConversationMembers struct {
	ID             uint   `gorm:"primaryKey"` // 记录ID
	ConversationID uint   // 关联的会话ID
	UserID         string // 用户ID
}

func (SysConversationMembers) TableName() string {
	return "sys_conversation_members"
}
