package system

import "ginDemo/global"

// 结构体里面的字段首字母大写才会导出使用，否则外部使用的时候则获取不到
type SysUser struct {
	global.MODEL
	Username string `gorm:"size:255" json:"name"`
	Password string
}

// 这里不能更改这个结构体 （数据库表）所有不带指针 TableName 方法本质上是一个只读操作 返回表明
func (SysUser) TableName() string {
	return "sys_users"
}
