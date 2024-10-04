package system

import (
	"time"
)

type SysConversations struct {
	ID          uint      `gorm:"primaryKey"`     // 会话ID
	LastMessage string    `gorm:"type:text"`      // 最后一条消息的内容
	UpdatedAt   time.Time `gorm:"autoUpdateTime"` // 更新时间
}

func (SysConversations) TableName() string {
	return "sys_conversations"
}
