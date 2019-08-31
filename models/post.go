package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Post struct {
	Id          int            `json:"id"`
	Title       string         `json:"title" orm:"size(1000)"`
	Contents    string         `json:"contents" orm:"type(text)"`
	ViewCount   int            `json:"view_count" orm:"default(0)"`
	Ip          string         `json:"ip" orm:"type(char)"`
	IsDeleted   bool           `json:"is_deleted" orm:"default(false)"`
	CreatedTime time.Time      `json:"created_time" orm:"auto_now_add"`
	Board       *Board         `json:"board" orm:"rel(fk)"`
	Comments    []*PostComment `json:"post_comments" orm:"reverse(many)"`
	Reports     []*Report      `json:"reports" orm:"reverse(many)"`
}

func AddPost(p *Post) error {
	o := orm.NewOrm()
	post := Post{
		Title:    p.Title,
		Contents: p.Contents,
		Ip:       p.Ip,
		Board:    p.Board,
	}

	if _, err := o.Insert(&post); err != nil {
		return errors.Wrap(err, "insert fail")
	}

	*p = post

	beego.Info(fmt.Sprintf("Success to add post, id:%d, title:%s, board name:%s, ip:%s",
		p.Id, p.Title, p.Board.Name, p.Ip))

	return nil
}

func FindPostById(id int) (*Post, error) {
	o := orm.NewOrm()
	post := Post{Id: id}

	if err := o.Read(&post); err != nil {
		return nil, errors.Wrap(err, "read fail")
	}

	beego.Info("Success to find post, id:%d, title:%s", id, post.Title)

	return &post, nil
}

func SetPostDeleteFlag(id int, isDeleted bool) error {
	o := orm.NewOrm()
	post := Post{Id: id}

	if err := o.Read(&post); err != nil {
		return errors.Wrap(err, "read fail")
	}

	post.IsDeleted = isDeleted
	if _, err := o.Update(&post); err != nil {
		return errors.Wrap(err, "update fail")
	}

	beego.Info("Success to set post delete flag, id:%d, flag:%v", id, isDeleted)

	return nil
}

func IncreasePostViewCount(id int) error {
	o := orm.NewOrm()
	post := Post{Id: id}

	if err := o.Read(&post); err != nil {
		return errors.Wrap(err, "read fail")
	}

	post.ViewCount++
	if _, err := o.Update(&post); err != nil {
		return errors.Wrap(err, "update fail")
	}

	beego.Info("Success to increse post view count, id:%d, view count:%v", id, post.ViewCount)

	return nil
}

func GetPostsByBoard(board *Board) ([]*Post, error) {
	o := orm.NewOrm()
	var ps []*Post

	num, err := o.QueryTable(new(Post)).Filter("Board__Id", board.Id).OrderBy("-created_time").All(&ps)
	if err != nil {
		return nil, errors.Wrap(err, "query by board fail")
	}

	beego.Info("Success to get post by board, board name:%s, num:%d",
		board.Name, num)

	return ps, nil
}
