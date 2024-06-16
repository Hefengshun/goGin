package demo

type SysDemo struct {
	Id       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;" json:"name"`
	Password string `gorm:"size:255;" json:"password"`
}

func (SysDemo) TableName() string {
	return "sys_demo"
}
