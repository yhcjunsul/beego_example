package main

import (
	"ex_login_chul/controllers"
	_ "ex_login_chul/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/join/", &controllers.JoinController{})
	beego.Router("/login/", &controllers.LoginController{}, "post:Login")
	beego.Router("/member/", &controllers.MemberController{}, "post:NewMember")
	beego.Router("/member/:id:string", &controllers.MemberController{}, "get:GetMember")
	beego.Run()
}
