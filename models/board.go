package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Board struct {
	Id            int            `orm:"pk"`
	Name          string         `orm:"size(100)"`
	BoardCategory *BoardCategory `orm:"reverse(one)"`
	IsDeleted     bool           `orm:"default(false)"`
}

func init() {
	orm.RegisterModel(new(Board))
}

func AddBoard(b *Board) error {
	o := orm.NewOrm()
	board := Board{Name: b.BoardName, BoardCategory: b.BoardCategory}

	if isCreated, id, err := o.ReadOrCreate(&board, "Name", "BoardCategory"); err != nil {
		return err
	}

	if isCreated == false {
		return fmt.Errorf("Board already exists, id:%d, board name:%s, category name:%s",
			id, b.Name, b.BoardCategory.Name)
	}

	beego.Info(fmt.Sprintf("Success to add board, id:%d, board name:%s, category name:%s",
		id, b.Name, b.BoardCategory.Name))

	return nil
}

func FindBoardById(id int) (*Board, error) {
	o := orm.NewOrm()
	board := Board{Id: id}

	if err := o.Read(&board); err != nil {
		return nil, err
	}

	beego.Info("Success to find board, id:" + id + ", name:" + board.Name)

	return &board, nil
}

func UpdateBoard(b *Board) error {
	o := orm.NewOrm()
	tmpForChecking := Board{Id: b.Id}

	if err := o.Read(&tmpForChecking); err != nil {
		return err
	}

	if num, err := o.Update(b); err != nil {
		return err
	}

	if num == 0 {
		return fmt.Errorf("Failed to update board, board not found, id:%d", b.Id)
	}

	beego.Info(fmt.Sprint("Success to update board, id:%d, name:%s, category name:%s",
		b.Id, b.Name, b.BoardCategory.Name))

	return nil
}

func DeleteBoard(id int) error {
	o := orm.NewOrm()

	if _, err := o.Delete(&Board{Id: id}); err != nil {
		return err
	}

	beego.Info("Success to delete board, id:" + id)

	return nil
}

func GetBoardsByCategory(category *BoardCategory) ([]*Board, error) {
	o := orm.NewOrm()
	var bs []*Board

	num, err := o.QueryTable(new(Board)).Filter("BoardCategory__Id", category.Id).All(&bs)
	if err != nil {
		return nil, err
	}

	beego.Info("Success to get board by category, category name:%s, num:%d",
		category.Name, num)

	return bs, nil
}

func GetAllBoards() ([]*Board, error) {
	o := orm.NewOrm()
	var bs []*Board

	num, err := o.QueryTable(new(Board)).All(&bs)
	if err != nil {
		return nil, err
	}

	beego.Info("Success to get all board, board count:" + num)

	return bs, nil
}
