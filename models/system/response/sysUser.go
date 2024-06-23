package response

import "ginDemo/models/system"

type Login struct {
	User  system.SysUser `json:"user"`
	Token string         `json:"token" binding:"required"`
}
