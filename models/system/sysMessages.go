package system

import (
	"time"
)

type SysMessage struct {
	ID             uint      `gorm:"primaryKey"` // 消息ID
	ConversationID uint      // 关联的会话ID
	SenderID       uint      // 发送者ID
	Content        string    `gorm:"type:text"`      // 消息内容
	CreatedAt      time.Time `gorm:"autoCreateTime"` // 消息创建时间
}
