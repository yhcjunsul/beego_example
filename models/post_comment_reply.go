package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type PostCommentReply struct {
	Id          int          `json:"id"`
	Contents    string       `json:"contents" orm:"size(4000)"`
	Ip          string       `json:"ip" orm:"type(char)"`
	IsDeleted   bool         `json:"is_deleted" orm:"default(false)"`
	CreatedTime time.Time    `json:"created_time" orm:"auto_now_add"`
	PostComment *PostComment `json:"post_comment" orm:"rel(fk)"`
	Reports     []*Report    `json:"reports" orm:"reverse(many)"`
}

func AddPostCommentReply(r *PostCommentReply) error {
	o := orm.NewOrm()
	reply := PostCommentReply{
		Contents:    r.Contents,
		Ip:          r.Ip,
		PostComment: r.PostComment,
	}

	if _, err := o.Insert(&reply); err != nil {
		return errors.Wrap(err, "insert fail")
	}

	*r = reply

	beego.Info(fmt.Sprintf("Success to add commment reply, id:%d, comment id:%s, ip:%s",
		r.Id, r.PostComment.Id, r.Ip))

	return nil
}

func FindPostCommentReplyById(id int) (*PostCommentReply, error) {
	o := orm.NewOrm()
	reply := PostCommentReply{Id: id}

	if err := o.Read(&reply); err != nil {
		return nil, errors.Wrap(err, "read fail")
	}

	beego.Info("Success to find comment reply, id:%d", id)

	return &reply, nil
}

func SetPostCommentReplyDeleteFlag(id int, isDeleted bool) error {
	o := orm.NewOrm()
	reply := PostCommentReply{Id: id}

	if err := o.Read(&reply); err != nil {
		return errors.Wrap(err, "read fail")
	}

	reply.IsDeleted = isDeleted
	if _, err := o.Update(&reply); err != nil {
		return errors.Wrap(err, "update fail")
	}

	beego.Info("Success to set post comment reply delete flag, id:%d, flag:%v", id, isDeleted)

	return nil
}

func GetPostCommentRepliesByPostComment(comment *PostComment) ([]*PostCommentReply, error) {
	o := orm.NewOrm()
	var rs []*PostCommentReply

	num, err := o.QueryTable(new(PostCommentReply)).Filter("PostComment__Id", comment.Id).OrderBy("created_time").All(&rs)
	if err != nil {
		return nil, errors.Wrap(err, "query by comment fail")
	}

	beego.Info("Success to get post comment replies by comment, comment id:%s, num:%d",
		comment.Id, num)

	return rs, nil
}
