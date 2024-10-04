package response

import "time"

type ConversationWithUnreadCount struct {
	ConversationID uint
	LastMessage    string
	UnreadCount    uint
	UpdatedAt      time.Time
	FriendOpenid   string
	FriendName     string
}
