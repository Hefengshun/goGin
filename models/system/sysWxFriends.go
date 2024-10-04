package system

import "ginDemo/global"

type SysWxFriends struct {
	global.MODEL
	UserOpenid   string `from:"user_openid" json:"user_openid"`
	FriendName   string `from:"friend_name" json:"friend_name"`
	FriendOpenid string `form:"friend_openid" json:"friend_openid"`
	Status       string `form:"status" json:"status" gorm:"comment:三种状态pending, accepted, reject"`
}

func (SysWxFriends) TableName() string {
	return "sys_wx_friends"
}
