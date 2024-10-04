package system

import (
	"time"
)

type SysMessages struct {
	ID             uint      `gorm:"primaryKey"` // 消息ID
	ConversationID uint      // 关联的会话ID
	SenderID       string    // 发送者ID
	Content        string    `gorm:"type:text"`      // 消息内容
	CreatedAt      time.Time `gorm:"autoCreateTime"` // 消息创建时间
}

func (SysMessages) TableName() string {
	return "sys_messages"
}
