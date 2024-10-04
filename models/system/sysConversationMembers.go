package system

type SysConversationMember struct {
	ID             uint `gorm:"primaryKey"` // 记录ID
	ConversationID uint // 关联的会话ID
	UserID         uint // 用户ID
}
