package controllers
import (
	"github.com/robfig/revel"
	"github.com/shavac/revel-auth/app/models"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Login() revel.Result {
	return c.Render()
}

func (c Auth) Logout() revel.Result {
	for k := range c.Session {
        delete(c.Session, k)
    }
	return c.Redirect(Auth.Login)
}

func (c Auth) Auth(username, password string) revel.Result {
	user := models.Authenticate(username, password)
	if user == nil {
		c.Flash.Error("Username and password does not match!")
		return c.Redirect(Auth.Login)
	} else {
		c.Session["username"] = user.Username
		return c.Redirect("/")
	}
}

func Logined(c *revel.Controller) bool {
	_, ok := c.Session["username"]
	return ok
}
