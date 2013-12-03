package app

import (
	"github.com/robfig/revel"
)

func init() {
	revel.InterceptFunc(LoginCheck, revel.BEFORE, revel.ALL_CONTROLLERS)
}
