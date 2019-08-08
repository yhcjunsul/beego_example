package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type ReportReason struct {
	Id       int       `json:"id"`
	Contents string    `json:"contents" orm:"size(1000)"`
	Reports  []*Report `json:"reports" orm:"reverse(many)"`
}

func AddReportReason(r *ReportReason) error {
	o := orm.NewOrm()
	reason := ReportReason{Contents: r.Contents}

	id, err := o.Insert(&reason)

	if err != nil {
		return errors.Wrap(err, "insert fail")
	}

	beego.Info("Success to add report reason, id:%d, contents:%s",
		id, r.Contents)

	return nil
}

func FindReportReasonById(id int) (*ReportReason, error) {
	o := orm.NewOrm()
	reason := ReportReason{Id: id}

	if err := o.Read(&reason); err != nil {
		return nil, errors.Wrap(err, "read fail")
	}

	beego.Info("Success to find report reason, id:%d", id)

	return &reason, nil
}

func UpdateReportReason(r *ReportReason) error {
	o := orm.NewOrm()
	tmpForChecking := ReportReason{Id: r.Id}

	if err := o.Read(&tmpForChecking); err != nil {
		return errors.Wrap(err, "read fail")
	}

	num, err := o.Update(r)

	if err != nil {
		return errors.Wrap(err, "update fail")
	}

	if num == 0 {
		return fmt.Errorf("Failed to update report reason, report reason not found, id:%d", r.Id)
	}

	beego.Info("Success to update report reason, id:%d", r.Id)

	return nil
}

func DeleteReportReason(id int) error {
	o := orm.NewOrm()

	num, err := o.Delete(&ReportReason{Id: id})
	if err != nil {
		return errors.Wrap(err, "delete fail")
	}

	if num == 0 {
		return errors.Errorf("Not found report reason for deleting")
	}

	beego.Info("Success to delete report reason, id:%d", id)

	return nil
}

func GetAllReportResons() ([]*ReportReason, error) {
	o := orm.NewOrm()
	var rs []*ReportReason

	num, err := o.QueryTable(new(ReportReason)).All(&rs)
	if err != nil {
		return nil, errors.Wrap(err, "query all fail")
	}

	beego.Info("Success to get all report reason, report reason count:%d", num)

	return rs, nil
}
