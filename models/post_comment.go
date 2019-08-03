package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type PostComment struct {
	Id          int                 `json:"id"`
	Contents    string              `json:"contents" orm:"size(4000)"`
	Ip          string              `json:"ip" orm:"type(char)"`
	IsDeleted   bool                `json:"is_deleted" orm:"default(false)"`
	CreatedTime time.Time           `json:"created_time" orm:"auto_now_add"`
	Post        *Post               `json:"post" orm:"rel(fk)"`
	Replies     []*PostCommentReply `orm:"reverse(many)"`
}

func AddPostComment(c *PostComment) error {
	o := orm.NewOrm()
	comment := PostComment{
		Contents: c.Contents,
		Ip:       c.Ip,
		Post:     c.Post,
	}

	if _, err := o.Insert(&comment); err != nil {
		return errors.Wrap(err, "insert fail")
	}

	beego.Info(fmt.Sprintf("Success to add commment, id:%d, post id:%s, ip:%s",
		c.Id, c.Post.Id, c.Ip))

	return nil
}

func FindPostCommentById(id int) (*PostComment, error) {
	o := orm.NewOrm()
	comment := PostComment{Id: id}

	if err := o.Read(&comment); err != nil {
		return nil, errors.Wrap(err, "read fail")
	}

	beego.Info("Success to find comment, id:%d", id)

	return &comment, nil
}

func SetPostCommentDeleteFlag(id int, isDeleted bool) error {
	o := orm.NewOrm()
	comment := PostComment{Id: id}

	if err := o.Read(&comment); err != nil {
		return errors.Wrap(err, "read fail")
	}

	comment.IsDeleted = isDeleted
	if _, err := o.Update(&comment); err != nil {
		return errors.Wrap(err, "update fail")
	}

	beego.Info("Success to set post comment delete flag, id:%d, flag:%v", id, isDeleted)

	return nil
}

func GetPostCommentsByPost(post *Post) ([]*PostComment, error) {
	o := orm.NewOrm()
	var cs []*PostComment

	num, err := o.QueryTable(new(PostComment)).Filter("Post__Id", post.Id).OrderBy("created_time").All(&cs)
	if err != nil {
		return nil, errors.Wrap(err, "query by post fail")
	}

	beego.Info("Success to get post comments by post, post id:%s, num:%d",
		post.Id, num)

	return cs, nil
}
