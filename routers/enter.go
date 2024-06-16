package routers

import (
	"ginDemo/routers/ces"
	"ginDemo/routers/system"
)

type RouterGroup struct {
	Ces    ces.RouterGroup
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
