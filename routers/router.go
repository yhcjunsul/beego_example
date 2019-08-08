// @APIVersion 1.0.0
// @Title SUGO.EE API
// @Description SUGO.EE API
// @Contact yhcjunsul@gmail.com
package routers

import (
	"github.com/yhcjunsul/beego_example/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	ns :=
		beego.NewNamespace("/v1",
			beego.NSInclude(
				&controllers.BoardController{},
				&controllers.BoardCategoryController{},
				&controllers.PostController{},
				&controllers.PostCommentController{},
				&controllers.PostCommentReplyController{},
				&controllers.ReportController{},
			),
		)

	beego.AddNamespace(ns)
}
