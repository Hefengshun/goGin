package system

import (
	"fmt"
	"ginDemo/global"
	"ginDemo/models/system"
)

type UserService struct {
}

func (_this *UserService) SignUp(u system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.DB {
		return nil, fmt.Errorf("db not init")
	}
	user := system.SysUser{
		Username: u.Username,
		Password: u.Password,
	}
	err = global.DB.Create(&user).Error
	return &user, err
}

func (_this *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.DB {
		return nil, fmt.Errorf("db not init")
	}
	var user system.SysUser
	//err = global.DB.Where("username = ? and password = ?", u.Username, u.Password).Find(&user).Error // Find找不到也返回空数据
	err = global.DB.Where("username = ? and password = ?", u.Username, u.Password).First(&user).Error //Fist找不到则返回 record nor found

	//if err == nil {
	//

	//}
	//if user {
	//
	//}
	return &user, err
}
