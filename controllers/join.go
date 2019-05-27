package controllers

import (
	"github.com/astaxie/beego"
)

type JoinController struct {
	beego.Controller
}

func (c *JoinController) Get() {
	c.TplName = "join.html"
	c.Render()
}
