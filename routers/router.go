package routers

import (
	"github.com/yhcjunsul/beego_example/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/board_categories", &controllers.BoardCategoryController{}, "get:GetAllBoardCategories")
	beego.Router("/board_category", &controllers.BoardCategoryController{}, "post:CreateBoardCategory")
	beego.Router("/board_category/:id:int", &controllers.BoardCategoryController{}, "delete:DeleteBoardCategory")

	beego.Router("/boards", &controllers.BoardController{}, "get:GetAllBoards")
	beego.Router("/*/board", &controllers.BoardController{}, "post:CreateBoard")
	beego.Router("/*/boards", &controllers.BoardController{}, "get:GetBoardsByCategory")
	beego.Router("/board/:id:int", &controllers.BoardController{}, "delete:DeleteBoard")

	beego.Router("/*/post", &controllers.PostController{}, "post:CreatePost")
	beego.Router("/*/posts", &controllers.PostController{}, "get:GetPostsByBoard")
	beego.Router("/post/:id:int", &controllers.PostController{}, "get:GetPostById;delete:DeletePost")

	beego.Router("/*/post_comment", &controllers.PostCommentController{}, "post:CreatePostComment")
	beego.Router("/*/post_comments", &controllers.PostCommentController{}, "get:GetPostCommentsByPost")
	beego.Router("/post_comment/:id:int", &controllers.PostCommentController{}, "delete:DeletePostComment")

	beego.Router("/*/post_comment_reply", &controllers.PostCommentReplyController{}, "post:CreatePostCommentReply")
	beego.Router("/*/post_comment_replies", &controllers.PostCommentReplyController{}, "get:GetPostCommentRepliesByPostComment")
	beego.Router("/post_comment_reply/:id:int", &controllers.PostCommentReplyController{}, "delete:DeletePostCommentReply")
}
