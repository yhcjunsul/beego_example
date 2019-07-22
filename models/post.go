package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id          int       `orm:"pk"`
	Title       string    `orm:"size(1000)"`
	Contents    string    `orm:"size(65535)"`
	ViewCount   int       `orm:"default(0)"`
	Ip          string    `orm:"type(char)"`
	IsDeleted   bool      `orm:"default(false)"`
	CreatedTime time.Time `orm:"auto_now_add"`
	Board       *Board    `orm:"reverse(one)"`
}

func init() {
	orm.RegisterModel(new(Post))
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
		return err
	}

	beego.Info(fmt.Sprintf("Success to add post, id:%d, title:%s, board name:%s, ip:%s",
		p.Id, p.Title, p.Board.Name, p.Ip))

	return nil
}

func FindPostById(id int) (*Post, error) {
	o := orm.NewOrm()
	post := Post{Id: id}

	if err := o.Read(&post); err != nil {
		return nil, err
	}

	beego.Info("Success to find post, id:" + id + ", title:" + post.Title)

	return &post, nil
}

func SetPostDeleteFlag(id int, isDeleted bool) error {
	o := orm.NewOrm()
	post := Post{Id: id}

	if err := o.Read(&post); err != nil {
		return err
	}

	post.IsDeleted = isDeleted
	if _, err := o.Update(&post); err != nil {
		return err
	}

	beego.Info("Success to set post delete flag, id:" + id + ", flag:" + isDeleted)

	return nil
}

func GetPostsByBoard(board *Board) ([]*Post, error) {
	o := orm.NewOrm()
	var ps []*Post

	num, err := o.QueryTable(new(Post)).Filter("Board__Id", board.Id).OrderBy("-created_time").All(&ps)
	if err != nil {
		return nil, err
	}

	beego.Info("Success to get post by board, board name:%s, num:%d",
		board.Name, num)

	return ps, nil
}
