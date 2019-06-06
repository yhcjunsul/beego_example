package controllers

import (
	"github.com/yhcjunsul/beego_example/models"

	"github.com/astaxie/beego"
)

type MemberController struct {
	beego.Controller
}

// Examples:
//
//   req: POST /member/ ["id": "", "password":"", "name":""]
//   res: 400 invalid id, invalid password, invalid name
//
//   req: POST /task/ ["id": "ldgmart", "password":"1234", "name":"daegyu"]
//   res: 200
func (this *MemberController) NewMember() {
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

	name := this.Ctx.Request.PostFormValue("name")
	if name == "" {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("invalid name"))
		return
	}

	m, err := models.NewMember(id, password, name)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	if err := models.DefaultMemberList.Add(m); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	beego.Info("new member, id:", id, "password:", password, "name:", name)
}

// Examples:
//
//   req: GET /member/ldgmart
//   res: 200 {"id": "ldgmart", "name": "daegyu"}
//
//   req: GET /member/dkdkdkdk
//   res: 404 task not found
func (this *MemberController) GetMember() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("ID is ", id)
	member, ok := models.DefaultMemberList.Find(id)
	beego.Info("Found", ok)
	if !ok {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("member not found"))
		return
	}
	this.Data["json"] = struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{member.ID, member.Name}
	this.ServeJSON()
}
