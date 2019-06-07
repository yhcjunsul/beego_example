package routers

import (
	"github.com/yhcjunsul/beego_example/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/member", &controllers.MemberController{}, "get:GetAllMembers;post:NewMember")
	beego.Router("/member/:id:string", &controllers.MemberController{}, "get:GetMember")
}
