package controllers

import (
	"net/http"
	"strconv"

	"github.com/yhcjunsul/beego_example/models"
	"github.com/yhcjunsul/beego_example/utils"

	"github.com/astaxie/beego"
)

type ReportController struct {
	beego.Controller
}

type PostReportCreateParam struct {
	Detail         string `json:"detail"`
	ReportReasonId int    `json:"report_reason_id"`
	PostId         int    `json:"post_id"`
}

type PostCommentReportCreateParam struct {
	Detail         string `json:"detail"`
	ReportReasonId int    `json:"report_reason_id"`
	PostCommentId  int    `json:"post_comment_id"`
}

type PostCommentReplyReportCreateParam struct {
	Detail             string `json:"detail"`
	ReportReasonId     int    `json:"report_reason_id"`
	PostCommentReplyId int    `json:"post_comment_reply_id"`
}

func (this *ReportController) URLMapping() {
	this.Mapping("CreatePostReport", this.CreatePostReport)
	this.Mapping("CreatePostCommentReport", this.CreatePostCommentReport)
	this.Mapping("CreatePostCommentReplyReport", this.CreatePostCommentReplyReport)
	this.Mapping("GetReportsByPost", this.GetReportsByPost)
	this.Mapping("GetReportsByPostComment", this.GetReportsByPostComment)
	this.Mapping("GetReportsByPostCommentReply", this.GetReportsByPostCommentReply)
}

// @Title Create post report
// @Summary Create post report
// @Param   detail                  body    string  false   "detail description of report"
// @Param   report_reason_id	    body	int    	true	"id of report reason"
// @Param   post_id                 body    int     true    "post id"
// @Success 200
// @Failure 400 Bad Request, invalid body contents
// @Failure 500 Internal server error
// @Accept json
// @router /report/post [post]
func (this *ReportController) CreatePostReport() {
	param := PostReportCreateParam{}
	report := models.Report{}

	if err := utils.UnmarshalRequestJson(this.Ctx.Input.RequestBody, &param); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid body contents")
		return
	}

	report.Detail = param.Detail

	ip := this.Ctx.Input.IP()
	report.Ip = ip

	reason, err := models.FindReportReasonById(param.ReportReasonId)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	report.ReportReason = reason

	post, err := models.FindPostById(param.PostId)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	report.Post = post

	if err = models.AddReport(&report); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	beego.Info("new post report, post id:%d, report reason id:%d", param.PostId, param.ReportReasonId)
}

// @Title Create post comment report
// @Summary Create post comment report
// @Param   detail                  body    string  false   "detail description of report"
// @Param   report_reason_id	    body	int    	true	"id of report reason"
// @Param   post_comment_id         body    int     true    "post comment id"
// @Success 200
// @Failure 400 Bad request, invalid body contents
// @Failure 500 Internal server error
// @Accept json
// @router /report/post_comment [post]
func (this *ReportController) CreatePostCommentReport() {
	param := PostCommentReportCreateParam{}
	report := models.Report{}

	if err := utils.UnmarshalRequestJson(this.Ctx.Input.RequestBody, &param); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid body contents")
		return
	}

	report.Detail = param.Detail

	ip := this.Ctx.Input.IP()
	report.Ip = ip

	reason, err := models.FindReportReasonById(param.ReportReasonId)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	report.ReportReason = reason

	comment, err := models.FindPostCommentById(param.PostCommentId)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	report.PostComment = comment

	if err = models.AddReport(&report); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	beego.Info("new post comment report, post comment id:%d, report reason id:%d", param.PostCommentId, param.ReportReasonId)
}

// @Title Create post comment reply report
// @Summary Create post comment reply report
// @Param   detail                  body    string  false   "detail description of report"
// @Param   report_reason_id	    body	int    	true	"id of report reason"
// @Param   post_comment_reply_id         body    int     true    "post comment reply id"
// @Success 200
// @Failure 400 Bad request, invalid body contents
// @Failure 500 Internal server error
// @Accept json
// @router /report/post_comment [post]
func (this *ReportController) CreatePostCommentReplyReport() {
	param := PostCommentReplyReportCreateParam{}
	report := models.Report{}

	if err := utils.UnmarshalRequestJson(this.Ctx.Input.RequestBody, &param); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid body contents")
		return
	}

	report.Detail = param.Detail

	ip := this.Ctx.Input.IP()
	report.Ip = ip

	reason, err := models.FindReportReasonById(param.ReportReasonId)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	report.ReportReason = reason

	reply, err := models.FindPostCommentReplyById(param.PostCommentReplyId)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	report.PostCommentReply = reply

	if err = models.AddReport(&report); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	beego.Info("new post comment reply report, post comment reply id:%d, report reason id:%d",
		param.PostCommentReplyId, param.ReportReasonId)
}

// @Title Get reports by post
// @Summary Get reports by post
// @Param id		path	int		true	"post id"
// @Success 200 {array} models.Report
// @Failure 400 Bad request, invalid post id
// @Failure 500 Internal server error
// @Accept json
// @router /reports/post/:id:int [get]
func (this *ReportController) GetReportsByPost() {
	post_id_param := this.Ctx.Input.Param(":id")

	post_id, err := strconv.Atoi(post_id_param)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid post id")
		return
	}

	post, err := models.FindPostById(post_id)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	reports, err := models.GetReportsByPost(post)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	this.Data["json"] = reports
	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Get reports by post comment
// @Summary Get reports by post comment
// @Param id		path	int		true	"post comment id"
// @Success 200 {array} models.Report
// @Failure 400 Bad Request, invalid comment id
// @Failure 500 Internal server error
// @Accept json
// @router /reports/post_comment/:id:int [get]
func (this *ReportController) GetReportsByPostComment() {
	comment_id_param := this.Ctx.Input.Param(":id")

	comment_id, err := strconv.Atoi(comment_id_param)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid comment id")
		return
	}

	comment, err := models.FindPostCommentById(comment_id)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	reports, err := models.GetReportsByPostComment(comment)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	this.Data["json"] = reports
	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Get reports by post comment reply
// @Summary Get reports by post comment reply
// @Param id		path	int		true	"post comment reply id"
// @Success 200 {array} models.Report
// @Failure 400 Bad request, invalid reply id
// @Failure 500 Internal server error
// @Accept json
// @router /reports/post_comment_reply/:id:int [get]
func (this *ReportController) GetReportsByPostCommentReply() {
	reply_id_param := this.Ctx.Input.Param(":id")

	reply_id, err := strconv.Atoi(reply_id_param)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid reply id")
		return
	}

	reply, err := models.FindPostCommentReplyById(reply_id)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	reports, err := models.GetReportsByPostCommentReply(reply)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	this.Data["json"] = reports
	beego.Info(this.Data["json"])
	this.ServeJSON()
}
