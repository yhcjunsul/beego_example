package controllers

import (
	"ex_login_chul/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

// Examples:
//
//   req: POST /login/
//	 post form : ["id" : "", "password" : ""]
//   res: 400  invalid password, invalid id
//		  404  member not found
//
//   req: POST /login/
//	 post form : ["id": "ldgmart", "password":"1234"}
//   res: 200
func (this *LoginController) Login() {
	id := this.Ctx.Request.PostFormValue("id")
	if id == "" {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("invalid id"))
		return
	}

	password := this.Ctx.Request.PostFormValue("password")
	if password == "" {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("invalid password"))
		return
	}

	member, isFound := models.DefaultMemberList.Find(id)
	if isFound == false {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("member not found"))
		return
	}

	if member.CheckPassword(password) == false {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("invalid password"))
		return
	}
}
