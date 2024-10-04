package system

import (
	"errors"
	"fmt"
	"ginDemo/global"
	"ginDemo/models/system"
	"ginDemo/models/system/request"
	"ginDemo/utils"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
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
	u.UUID = utils.GenerateUuid()

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

func (_this *UserService) WxLogin(code string) (userInter *system.SysWxUser, err error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=wxe0c70520ee6a99ff&secret=ca073ce7bf15019d8c6f794c361f4363&js_code=" + code + "&grant_type=authorization_code"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() // 确保函数结束时关闭 body

	// 读取响应 body 数据
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// 使用 map[string]string 来解析 JSON
	var body map[string]string
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return nil, err
	}

	// 使用解析后的 openid 来查询数据库
	err = global.DB.Where("openid = ?", body["openid"]).First(&userInter).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("Record not found, creating new user")

		// 创建新用户
		var user = system.SysWxUser{
			UserName: "Super" + strconv.Itoa(utils.GenerateSimpleRandomNumber()),
			Password: "",
			UUID:     utils.GenerateUuid(), // 生成 UUID
			Openid:   body["openid"],       // 使用解析到的 openid
		}

		// 插入新用户
		err = global.DB.Create(&user).Error
		if err != nil {
			return nil, err // 如果插入出错，返回错误
		}

		return &user, nil // 返回新创建的用户
	} else if err != nil {
		return nil, err // 如果查询出错，返回错误
	}

	// 如果查询成功，返回已找到的用户
	return userInter, nil
}

func (_this *UserService) UpdateUser(reqUser request.UpdateUser) (message string, err error) {
	var user system.SysWxUser
	err = global.DB.Where(`openid = ?`, reqUser.UserOpenid).Find(&user).Error
	if err != nil {
		return err.Error(), nil
	} else {
		user.UserName = reqUser.UserName
		err = global.DB.Save(&user).Error
		if err != nil {
			return err.Error(), nil
		}
	}
	return "修改成功", nil
}

func (_this *UserService) WxAddFriends(userOpenid string, friendOpenid string) (msg string, err error) {
	var user system.SysWxUser
	// 先去重用户表里面找是是否纯在查找用户
	err = global.DB.Where("openid = ?", friendOpenid).First(&user).Error
	if err != nil {
		return "未查找到该用户！", err
	} else {
		// 如果查找到
		var friendData system.SysWxFriends
		err = global.DB.Where("user_openid = ? AND friend_openid = ?", userOpenid, friendOpenid).First(&friendData).Error
		// 则判断朋友表里是否有这条数据
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				//没有则往朋友表里面添加一条数据
				var wxFriend = system.SysWxFriends{
					UserOpenid:   userOpenid,
					FriendOpenid: friendOpenid, //找到朋友id添加
					FriendName:   user.UserName,
					Status:       "pending",
				}
				if err := global.DB.Create(&wxFriend).Error; err != nil {
					// 处理插入错误
					return "数据添加失败了！", err
				}
				return "请求已发送！", nil
			} else {
				// 处理其他数据库错误
				return "数据库错误！", err
			}
		} else {
			return "请求已发送，请耐心等候！", nil
		}
	}

}

func (_this *UserService) GetUserFriends(userOpenid string, friendStatus string) (friendList []system.SysWxFriends, err error) {
	err = global.DB.Where("(user_openid = ? OR friend_openid = ?) AND status = ?", userOpenid, userOpenid, friendStatus).Find(&friendList).Error
	if err != nil {
		return nil, err
	}

	// 收集所有需要查询的 openid
	openidSet := make(map[string]bool)
	for _, friend := range friendList {
		if friend.FriendOpenid == userOpenid {
			openidSet[friend.UserOpenid] = true
		}
	}

	// 将 openid 转换为切片
	var openids []string
	for openid := range openidSet {
		openids = append(openids, openid)
	}

	// 批量查询用户信息
	var users []system.SysWxUser
	err = global.DB.Where("openid IN (?)", openids).Find(&users).Error
	if err != nil {
		return nil, err
	}

	// 创建一个映射以便快速查找用户名称
	userMap := make(map[string]string)
	for _, user := range users {
		userMap[user.Openid] = user.UserName
	}

	// 更新 friendList 中的 FriendName
	for i, friend := range friendList {
		if friend.FriendOpenid == userOpenid {
			if userName, exists := userMap[friend.UserOpenid]; exists {
				friendList[i].FriendName = userName
			}
		}
	}

	return friendList, nil
}

func (_this *UserService) HandleFriendApply(id string, status string) (message string, err error) {
	var friendApply system.SysWxFriends
	err = global.DB.Where("id = ?", id).First(&friendApply).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "未找到申请数据", nil
		}
		return err.Error(), nil
	} else {
		friendApply.Status = status
		err = global.DB.Save(&friendApply).Error
		if err != nil {
			return err.Error(), nil
		}
	}
	if status == "accepted" {
		return "加为好友！", nil
	}
	return "拒绝添加", nil
}
