package app

import (
	"github.com/robfig/revel"
	"github.com/shavac/revel-auth/app/controllers"
	"strings"
)

var (
	AuthConfigLoaded   = false
	BypassControllers  []string
	BypassActions      []string
	AuthAllControllers bool   = false
	LoginURL           string = "/login"
)

func LoginCheck(c *revel.Controller) revel.Result {
	if !AuthConfigLoaded {
		LoadAuthConfig()
	}
	if AuthAllControllers {
		if 	controllers.Logined(c) ||
			c.Request.URL.Path == LoginURL ||
			revel.ContainsString(BypassControllers, c.Name) ||
			revel.ContainsString(BypassActions, c.Action) {
			return nil
		} else {
			return c.Redirect(LoginURL)
		}
	}
	return nil
}

func LoadAuthConfig() {
	if authconf, err := revel.LoadConfig("auth.conf"); err == nil {
		authconf.SetSection("auth")
		AuthAllControllers = authconf.BoolDefault("auth.controllers.all", false)
		LoginURL = authconf.StringDefault("auth.login.url", "/login")
		BypassControllers = strings.Fields(authconf.StringDefault("auth.bypass.controller", ""))
		BypassActions = strings.Fields(authconf.StringDefault("auth.bypass.action", ""))
		AuthConfigLoaded = true
	} else {
		AuthConfigLoaded = false
	}
}
