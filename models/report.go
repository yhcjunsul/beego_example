package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Report struct {
	Id               int               `json:"id"`
	Detail           string            `json:"detail" orm:"size(4000)"`
	CreatedTime      time.Time         `json:"created_time" orm:"auto_now_add"`
	ReportReason     *ReportReason     `json:"report_reason" orm:"rel(fk)"`
	Ip               string            `json:"ip" orm:"type(char)"`
	Post             *Post             `json:"post" orm:"rel(fk);null"`
	PostComment      *PostComment      `json:"post_comment" orm:"rel(fk);null"`
	PostCommentReply *PostCommentReply `json:"post_comment_replie" orm:"rel(fk);null"`
}

func AddReport(r *Report) error {
	o := orm.NewOrm()
	report := Report{
		Detail:           r.Detail,
		ReportReason:     r.ReportReason,
		Ip:               r.Ip,
		Post:             r.Post,
		PostComment:      r.PostComment,
		PostCommentReply: r.PostCommentReply,
	}

	id, err := o.Insert(&report)

	if err != nil {
		return errors.Wrap(err, "insert fail")
	}

	*r = report

	beego.Info("Success to add report, id:%d, report reason id:%d",
		id, r.ReportReason.Id)

	return nil
}

func FindReportById(id int) (*Report, error) {
	o := orm.NewOrm()
	report := Report{Id: id}

	if err := o.Read(&report); err != nil {
		return nil, errors.Wrap(err, "read fail")
	}

	beego.Info("Success to find report, id:%d", id)

	return &report, nil
}

func UpdateReport(r *Report) error {
	o := orm.NewOrm()
	tmpForChecking := Report{Id: r.Id}

	if err := o.Read(&tmpForChecking); err != nil {
		return errors.Wrap(err, "read fail")
	}

	num, err := o.Update(r)

	if err != nil {
		return errors.Wrap(err, "update fail")
	}

	if num == 0 {
		return fmt.Errorf("Failed to update report, report not found, id:%d", r.Id)
	}

	beego.Info("Success to update report, id:%d", r.Id)

	return nil
}

func DeleteReport(id int) error {
	o := orm.NewOrm()

	num, err := o.Delete(&Report{Id: id})
	if err != nil {
		return errors.Wrap(err, "delete fail")
	}

	if num == 0 {
		return errors.Errorf("Not found report for deleting")
	}

	beego.Info("Success to delete report, id:%d", id)

	return nil
}

func GetReportsByPost(post *Post) ([]*Report, error) {
	o := orm.NewOrm()
	var rs []*Report

	num, err := o.QueryTable(new(Report)).Filter("Post__Id", post.Id).OrderBy("-created_time").All(&rs)
	if err != nil {
		return nil, errors.Wrap(err, "query by post fail")
	}

	beego.Info("Success to get report by post, post id:%d, num:%d",
		post.Id, num)

	return rs, nil
}

func GetReportsByPostComment(postComment *PostComment) ([]*Report, error) {
	o := orm.NewOrm()
	var rs []*Report

	num, err := o.QueryTable(new(Report)).Filter("PostComment__Id", postComment.Id).OrderBy("-created_time").All(&rs)
	if err != nil {
		return nil, errors.Wrap(err, "query by post comment fail")
	}

	beego.Info("Success to get report by post comment, post comment id:%d, num:%d",
		postComment.Id, num)

	return rs, nil
}

func GetReportsByPostCommentReply(postCommentReply *PostCommentReply) ([]*Report, error) {
	o := orm.NewOrm()
	var rs []*Report

	num, err := o.QueryTable(new(Report)).Filter("PostCommentReply__Id", postCommentReply.Id).OrderBy("-created_time").All(&rs)
	if err != nil {
		return nil, errors.Wrap(err, "query by post comment reply fail")
	}

	beego.Info("Success to get report by post comment reply, post id:%d, num:%d",
		postCommentReply.Id, num)

	return rs, nil
}

func GetReportsByReportReason(reportReason *ReportReason) ([]*Report, error) {
	o := orm.NewOrm()
	var rs []*Report

	num, err := o.QueryTable(new(Report)).Filter("ReportReason__Id", reportReason.Id).OrderBy("-created_time").All(&rs)
	if err != nil {
		return nil, errors.Wrap(err, "query by report reason fail")
	}

	beego.Info("Success to get report by report reason reply, report reason id:%d, num:%d",
		reportReason.Id, num)

	return rs, nil
}

func GetAllReports() ([]*Report, error) {
	o := orm.NewOrm()
	var rs []*Report

	num, err := o.QueryTable(new(Report)).All(&rs)
	if err != nil {
		return nil, errors.Wrap(err, "query all fail")
	}

	beego.Info("Success to get all report, report count:%d", num)

	return rs, nil
}
