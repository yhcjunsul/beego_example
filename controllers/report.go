package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/yhcjunsul/beego_example/models"

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
// @Failure 400 Bad Request
// @Accept json
// @router /report/post [post]
func (this *ReportController) CreatePostReport() {
	param := PostReportCreateParam{}
	report := models.Report{}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &param); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	report.Detail = param.Detail

	ip := this.Ctx.Input.IP()
	report.Ip = ip

	reason, err := models.FindReportReasonById(param.ReportReasonId)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	report.ReportReason = reason

	post, err := models.FindPostById(param.PostId)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	report.Post = post

	if err = models.AddReport(&report); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
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
// @Failure 400 Bad Request
// @Accept json
// @router /report/post_comment [post]
func (this *ReportController) CreatePostCommentReport() {
	param := PostCommentReportCreateParam{}
	report := models.Report{}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &param); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	report.Detail = param.Detail

	ip := this.Ctx.Input.IP()
	report.Ip = ip

	reason, err := models.FindReportReasonById(param.ReportReasonId)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	report.ReportReason = reason

	comment, err := models.FindPostCommentById(param.PostCommentId)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	report.PostComment = comment

	if err = models.AddReport(&report); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
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
// @Failure 400 Bad Request
// @Accept json
// @router /report/post_comment [post]
func (this *ReportController) CreatePostCommentReplyReport() {
	param := PostCommentReplyReportCreateParam{}
	report := models.Report{}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &param); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	report.Detail = param.Detail

	ip := this.Ctx.Input.IP()
	report.Ip = ip

	reason, err := models.FindReportReasonById(param.ReportReasonId)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	report.ReportReason = reason

	reply, err := models.FindPostCommentReplyById(param.PostCommentReplyId)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	report.PostCommentReply = reply

	if err = models.AddReport(&report); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	beego.Info("new post comment reply report, post comment reply id:%d, report reason id:%d",
		param.PostCommentReplyId, param.ReportReasonId)
}

// @Title Get reports by post
// @Summary Get reports by post
// @Param id		path	int		true	"post id"
// @Success 200 {array} models.Report
// @Failure 400 Bad Request
// @Failure 404 Not found
// @Accept json
// @router /reports/post/:id:int [get]
func (this *ReportController) GetReportsByPost() {
	post_id_param := this.Ctx.Input.Param(":id")

	post_id, err := strconv.Atoi(post_id_param)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad request"))
	}

	post, err := models.FindPostById(post_id)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	reports, err := models.GetReportsByPost(post)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Not Found"))
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
// @Failure 400 Bad Request
// @Failure 404 Not found
// @Accept json
// @router /reports/post_comment/:id:int [get]
func (this *ReportController) GetReportsByPostComment() {
	comment_id_param := this.Ctx.Input.Param(":id")

	comment_id, err := strconv.Atoi(comment_id_param)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad request"))
	}

	comment, err := models.FindPostCommentById(comment_id)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	reports, err := models.GetReportsByPostComment(comment)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Not Found"))
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
// @Failure 400 Bad Request
// @Failure 404 Not found
// @Accept json
// @router /reports/post_comment_reply/:id:int [get]
func (this *ReportController) GetReportsByPostCommentReply() {
	reply_id_param := this.Ctx.Input.Param(":id")

	reply_id, err := strconv.Atoi(reply_id_param)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad request"))
	}

	reply, err := models.FindPostCommentReplyById(reply_id)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	reports, err := models.GetReportsByPostCommentReply(reply)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Not Found"))
		return
	}

	this.Data["json"] = reports
	beego.Info(this.Data["json"])
	this.ServeJSON()
}
