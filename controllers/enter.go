package controllers

import (
	"ginDemo/controllers/demo"
	"ginDemo/controllers/system"
)

type ControllerGroup struct {
	DemoControllerGroup   demo.ControllerGroup
	SystemControllerGroup system.ControllerGroup
}

var ControllerGroupApp = new(ControllerGroup)
