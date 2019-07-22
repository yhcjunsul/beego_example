package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PostComment struct {
	Id          int       `orm:"pk"`
	Contents    string    `orm:"size(4000)"`
	Ip          string    `orm:"type(char)"`
	IsDeleted   bool      `orm:"default(false)"`
	CreatedTime time.Time `orm:"auto_now_add"`
	Post        *Post     `orm:"reverse(fk)"`
}

func init() {
	orm.RegisterModel(new(PostComment))
}

func AddPostComment(c *PostComment) error {
	o := orm.NewOrm()
	comment := PostComment{
		Contents: c.Contents,
		Ip:       c.Ip,
		Post:     c.Post,
	}

	if _, err := o.Insert(&comment); err != nil {
		return err
	}

	beego.Info(fmt.Sprintf("Success to add commment, id:%d, post id:%s, ip:%s",
		c.Id, c.Post.Id, c.Ip))

	return nil
}

func FindPostCommentById(id int) (*PostComment, error) {
	o := orm.NewOrm()
	comment := PostComment{Id: id}

	if err := o.Read(&comment); err != nil {
		return nil, err
	}

	beego.Info("Success to find comment, id:" + id)

	return &comment, nil
}

func SetPostCommentDeleteFlag(id int, isDeleted bool) error {
	o := orm.NewOrm()
	comment := PostComment{Id: id}

	if err := o.Read(&comment); err != nil {
		return err
	}

	comment.IsDeleted = isDeleted
	if _, err := o.Update(&comment); err != nil {
		return err
	}

	beego.Info("Success to set post comment delete flag, id:" + id + ", flag:" + isDeleted)

	return nil
}

func GetPostCommentsByPost(post *Post) ([]*PostComment, error) {
	o := orm.NewOrm()
	var cs []*PostComment

	num, err := o.QueryTable(new(PostComment)).Filter("Post__Id", post.Id).OrderBy("created_time").All(&ps)
	if err != nil {
		return nil, err
	}

	beego.Info("Success to get post comments by post, post id:%s, num:%d",
		post.Id, num)

	return ps, nil
}
