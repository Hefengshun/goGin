package system

import (
	"ginDemo/global"
	"github.com/gofrs/uuid/v5"
)

// 结构体里面的字段首字母大写才会导出使用，否则外部使用的时候则获取不到
type SysUser struct {
	global.MODEL
	//UserName 在数据库字段为 user_name 而后面的json则是 JSON编码和解码时指定字段的名称
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`    // 用户UUID
	UserName string    `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	Password string    `json:"-"  gorm:"comment:用户登录密码"`
}

// 这里不能更改这个结构体 （数据库表）所有不带指针 TableName 方法本质上是一个只读操作 返回表明
func (SysUser) TableName() string {
	return "sys_users"
}
