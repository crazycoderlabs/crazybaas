package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	return c.Render()
}

func (c App) Login() revel.Result	{
	return c.Render()
}

func (c App) DoLogin(username string, password string)	revel.Result	{
	c.Validation.Required(username).MessageKey("signin.v.username_required")
	c.Validation.MaxSize(username,20).MessageKey("signin.v.username_max_size")
	c.Validation.MinSize(password,2).MessageKey("signin.v.username_min_size")
	c.Validation.Required(password).MessageKey("signin.v.password_required")
	c.Validation.MaxSize(password,20).MessageKey("signin.v.password_max_size")
	c.Validation.MinSize(password,7).MessageKey("signin.v.password_min_size")
	if c.Validation.HasErrors()	{
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Login)
	}
	return c.RenderTemplate("Index.html")
}