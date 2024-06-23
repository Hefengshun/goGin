package request

type SignUp struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Name     string `json:"name"`                        // 用户名
	Password string `json:"password" binding:"required"` // 密码 结构体绑定校验
}
