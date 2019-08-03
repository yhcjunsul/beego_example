package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yhcjunsul/beego_example/models"

	"github.com/astaxie/beego"
)

type PostCommentReplyController struct {
	beego.Controller
}

// @Title Create post comment reply
// @Summary Create post comment reply
// @Param   contents	body	string	true	"Contents of comment reply"
// @Param   ip			body	string	true	"IP of user"
// @Param  	comment_id	path	int		true	"comment id"
// @Success 200
// @Failure 400 Bad Request
// @Failure 404 Not found
// @Accept json
// @router /post_comment/:id/post_comment_reply(/*/post_comment_reply) [post]
func (this *PostCommentReplyController) CreatePostCommentReply() {
	reply := models.PostCommentReply{}

	splat := this.Ctx.Input.Param(":splat")
	// splat == "post_comment/:id"
	splat_splits := strings.Split(splat, "/")

	if len(splat_splits) != 2 {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	if splat_splits[0] != "post_comment" {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	comment_id, err := strconv.Atoi(splat_splits[1])
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &reply); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	reply.PostComment, err = models.FindPostCommentById(comment_id)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Not found"))
		return
	}

	if err = models.AddPostCommentReply(&reply); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	beego.Info("new post comment reply, user ip:%s", reply.Ip)
}

// @Title Get post comment replies by comment
// @Summary Get post comment replies by comment
// @Param comment_id 	path	int	true	"comment id"
// @Success 200 {object}array models.PostCommentReply
// @Failure 400 Bad request
// @Failure 404 Not found
// @Accept json
// @router /post_comment/:id/post_comment_replies(/*/post_comment_replies) [get]
func (this *PostCommentReplyController) GetPostCommentRepliesByPostComment() {
	splat := this.Ctx.Input.Param(":splat")
	// splat == "post_comment/:id"
	splat_splits := strings.Split(splat, "/")

	if len(splat_splits) != 2 {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	if splat_splits[0] != "post_comment" {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	comment_id, err := strconv.Atoi(splat_splits[1])
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	comment := models.PostComment{Id: comment_id}
	replies, err := models.GetPostCommentRepliesByPostComment(&comment)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Not found post"))
		return
	}

	this.Data["json"] = replies
	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Delete post comment reply
// @Summary Delete post comment reply by ID
// @Success 200
// @Failure 404 Not found
// @Accept json
// @router /post_comment_reply/:id [delete]
func (this *PostCommentReplyController) DeletePostCommentReply() {
	id_param := this.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad request"))
	}

	if err := models.SetPostCommentReplyDeleteFlag(id, true); err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Delete post comment reply error:" + err.Error()))
		return
	}
}
