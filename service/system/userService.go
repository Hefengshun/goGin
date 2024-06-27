package system

import (
	"errors"
	"fmt"
	"ginDemo/global"
	"ginDemo/models/system"
	"ginDemo/utils"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type UserService struct {
}

func (_this *UserService) SignUp(u system.SysUser) (userInter system.SysUser, err error) {

	var user = system.SysUser{}

	if !errors.Is(global.DB.Where("user_name = ?", u.UserName).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())

	err = global.DB.Create(&u).Error
	return u, err
}

func (_this *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser

	//err = global.DB.Where("user_name = ? and password = ?", u.Username, u.Password).Find(&user).Error // Find找不到也返回空数据
	//***  表字段使用了驼峰命名  数据库生成的是下划线命名法  ***
	err = global.DB.Where("user_name = ?", u.UserName).First(&user).Error //Fist找不到则返回 record nor found
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}

	return &user, err
}
