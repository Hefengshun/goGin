package system

import "ginDemo/service"

type ControllerGroup struct {
	UploadController
	UserController
}

// 这里实例化 service
var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
