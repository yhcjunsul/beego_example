package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PostCommentReply struct {
	Id          int          `orm:"pk"`
	Contents    string       `orm:"size(4000)"`
	Ip          string       `orm:"type(char)"`
	IsDeleted   bool         `orm:"default(false)"`
	CreatedTime time.Time    `orm:"auto_now_add"`
	PostComment *PostComment `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(PostCommentReply))
}

func AddPostCommentReply(r *PostCommentReply) error {
	o := orm.NewOrm()
	reply := PostCommentReply{
		Contents:    r.Contents,
		Ip:          r.Ip,
		PostComment: r.PostComment,
	}

	if _, err := o.Insert(&reply); err != nil {
		return err
	}

	beego.Info(fmt.Sprintf("Success to add commment reply, id:%d, comment id:%s, ip:%s",
		r.Id, r.PostComment.Id, r.Ip))

	return nil
}

func FindPostCommentReplyById(id int) (*PostCommentReply, error) {
	o := orm.NewOrm()
	reply := PostCommentReply{Id: id}

	if err := o.Read(&reply); err != nil {
		return nil, err
	}

	beego.Info("Success to find comment reply, id:" + id)

	return &reply, nil
}

func SetPostCommentReplyDeleteFlag(id int, isDeleted bool) error {
	o := orm.NewOrm()
	reply := PostCommentReply{Id: id}

	if err := o.Read(&reply); err != nil {
		return err
	}

	reply.IsDeleted = isDeleted
	if _, err := o.Update(&reply); err != nil {
		return err
	}

	beego.Info("Success to set post comment reply delete flag, id:" + id + ", flag:" + isDeleted)

	return nil
}

func GetPostCommentRepliesByPostComment(comment *PostComment) ([]*PostCommentReply, error) {
	o := orm.NewOrm()
	var rs []*PostCommentReply

	num, err := o.QueryTable(new(PostCommentReply)).Filter("PostComment__Id", comment.Id).OrderBy("created_time").All(&ps)
	if err != nil {
		return nil, err
	}

	beego.Info("Success to get post comment replies by comment, comment id:%s, num:%d",
		comment.Id, num)

	return rs, nil
}
