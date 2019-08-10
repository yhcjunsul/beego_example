package controllers

import (
	"strconv"

	"github.com/yhcjunsul/beego_example/models"
	"github.com/yhcjunsul/beego_example/utils"

	"github.com/astaxie/beego"
)

type PostCommentController struct {
	beego.Controller
}

func (this *PostCommentController) URLMapping() {
	this.Mapping("CreatePostComment", this.CreatePostComment)
	this.Mapping("GetPostCommentsByPost", this.GetPostCommentsByPost)
	this.Mapping("DeletePostComment", this.DeletePostComment)
}

// @Title Create post comment
// @Summary Create post comment
// @Param   contents	body	string	true	"Contents of comment"
// @Param  	post_id		path	int		true	"post id"
// @Success 200
// @Failure 400 Bad Request
// @Failure 404 Not found
// @Accept json
// @router /post/:postId:int/post_comment [post]
func (this *PostCommentController) CreatePostComment() {
	comment := models.PostComment{}

	post_id, err := strconv.Atoi(this.Ctx.Input.Param(":postId"))
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	if err := utils.UnmarshalRequestJson(this.Ctx.Input.RequestBody, &comment); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	comment.Post, err = models.FindPostById(post_id)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Not found"))
		return
	}

	ip := this.Ctx.Input.IP()
	comment.Ip = ip

	if err = models.AddPostComment(&comment); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	beego.Info("new post comment, user ip:%s", comment.Ip)
}

// @Title Get post comments by post
// @Summary Get post comments by post
// @Param post_id 	path	int	true	"post id"
// @Success 200 {array} models.PostComment
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router /post/:postId:int/post_comments [get]
func (this *PostCommentController) GetPostCommentsByPost() {
	post_id, err := strconv.Atoi(this.Ctx.Input.Param(":postId"))
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	post := models.Post{Id: post_id}
	comments, err := models.GetPostCommentsByPost(&post)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Not found post"))
		return
	}

	this.Data["json"] = comments
	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Delete post comment
// @Summary Delete post comment by ID
// @Success 200
// @Failure 404 Not found
// @Accept json
// @router /post_comment/:id:int [delete]
func (this *PostCommentController) DeletePostComment() {
	id_param := this.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad request"))
	}

	if err := models.SetPostCommentDeleteFlag(id, true); err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Delete post comment error:" + err.Error()))
		return
	}
}
