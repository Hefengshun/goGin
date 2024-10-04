package request

type SignUp struct {
	UserName string `json:"UserName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	UserName string `json:"userName"`                    // 用户名
	Password string `json:"password" binding:"required"` // 密码 结构体绑定校验
}

type UpdateUser struct {
	UserName   string `json:"userName" binding:"required"`
	UserOpenid string `json:"openid" binding:"required"`
}
