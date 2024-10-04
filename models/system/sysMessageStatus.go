package system

type SysMessageStatus struct {
	MessageID uint   `gorm:"primaryKey"`               // 消息ID
	UserID    uint   `gorm:"primaryKey"`               // 用户ID
	Status    string `gorm:"size:20;default:'unread'"` // 消息状态 (unread, read)
}
