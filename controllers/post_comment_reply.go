package controllers

import (
	"net/http"
	"strconv"

	"github.com/yhcjunsul/beego_example/models"
	"github.com/yhcjunsul/beego_example/utils"

	"github.com/astaxie/beego"
)

type PostCommentReplyController struct {
	beego.Controller
}

func (this *PostCommentReplyController) URLMapping() {
	this.Mapping("CreatePostCommentReply", this.CreatePostCommentReply)
	this.Mapping("GetPostCommentRepliesByPostComment", this.GetPostCommentRepliesByPostComment)
	this.Mapping("DeletePostCommentReply", this.DeletePostCommentReply)
}

// @Title Create post comment reply
// @Summary Create post comment reply
// @Param   contents	body	string	true	"Contents of comment reply"
// @Param  	comment_id	path	int		true	"comment id"
// @Success 200 {object} models.PostCommentReply
// @Failure 400 Bad request, invalid comment id
// @Failure 400 Bad request, invalid body contents
// @Failure 404 Not found comment
// @Failure 500 Internal server error
// @Accept json
// @router /post_comment/:commentId:int/post_comment_reply [post]
func (this *PostCommentReplyController) CreatePostCommentReply() {
	reply := models.PostCommentReply{}

	comment_id, err := strconv.Atoi(this.Ctx.Input.Param(":commentId"))
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid comment id")
		return
	}

	if err := utils.UnmarshalRequestJson(this.Ctx.Input.RequestBody, &reply); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid body contents")
		return
	}

	reply.PostComment, err = models.FindPostCommentById(comment_id)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusNotFound, "Not found comment")
		return
	}

	ip := this.Ctx.Input.IP()
	reply.Ip = ip

	if err = models.AddPostCommentReply(&reply); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	this.Data["json"] = reply
	this.ServeJSON()

	beego.Info("new post comment reply, user ip:%s", reply.Ip)
}

// @Title Get post comment replies by comment
// @Summary Get post comment replies by comment
// @Param comment_id 	path	int	true	"comment id"
// @Success 200 {array} models.PostCommentReply
// @Failure 400 Bad request, invalid comment id
// @Failure 500 Internal server error
// @Accept json
// @router /post_comment/:commentId:int/post_comment_replies [get]
func (this *PostCommentReplyController) GetPostCommentRepliesByPostComment() {
	comment_id, err := strconv.Atoi(this.Ctx.Input.Param(":commentId"))
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid comment id")
		return
	}

	comment := models.PostComment{Id: comment_id}
	replies, err := models.GetPostCommentRepliesByPostComment(&comment)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	this.Data["json"] = replies
	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Delete post comment reply
// @Summary Delete post comment reply by ID
// @Success 200
// @Failure 400 Bad request, invalid reply id
// @Failure 500 Internal server error
// @Accept json
// @router /post_comment_reply/:id:int [delete]
func (this *PostCommentReplyController) DeletePostCommentReply() {
	id_param := this.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid reply id")
		return
	}

	if err := models.SetPostCommentReplyDeleteFlag(id, true); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server erro")
		return
	}
}
