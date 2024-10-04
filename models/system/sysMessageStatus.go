package system

type SysMessageStatus struct {
	MessageID uint   `gorm:"primaryKey"` // 消息ID
	UserID    string // 用户ID
	Status    string `gorm:"size:20;default:'unread'"` // 消息状态 (unread, read)
}

func (SysMessageStatus) TableName() string {
	return "sys_message_status"
}
