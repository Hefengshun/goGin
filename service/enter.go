package service

import (
	"ginDemo/service/demo"
	"ginDemo/service/system"
)

type ServiceGroup struct {
	DemoServiceGroup   demo.ServiceGroup
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
