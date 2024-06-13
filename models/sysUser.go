package models

import "ginDemo/global"

// 结构体里面的字段首字母大写才会导出使用，否则外部使用的时候则获取不到
type Users struct {
	global.MODEL
	Username string `gorm:"size:255" json:"name"`
	Password string
}

func (Users) TableName() string {
	return "sys_users"
}
