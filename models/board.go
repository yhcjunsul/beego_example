package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Board struct {
	Id            int            `json:"id"`
	Name          string         `json:"name" orm:"size(100)"`
	IsDeleted     bool           `json:"is_deleted" orm:"default(false)"`
	BoardCategory *BoardCategory `json:"board_category" orm:"rel(fk)"`
	Posts         []*Post        `orm:"reverse(many)"`
}

func AddBoard(b *Board) error {
	o := orm.NewOrm()
	board := Board{Name: b.Name, BoardCategory: b.BoardCategory}

	isCreated, id, err := o.ReadOrCreate(&board, "Name", "BoardCategory")

	if err != nil {
		return errors.Wrap(err, "read or create fail")
	}

	if isCreated == false {
		return fmt.Errorf("Board already exists, id:%d, board name:%s, category name:%s",
			id, b.Name, b.BoardCategory.Name)
	}

	*b = board

	beego.Info("Success to add board, id:", id, ", board id:", b.Id,
		", board name:", b.Name, ", category name:", b.BoardCategory.Name)

	return nil
}

func FindBoardById(id int) (*Board, error) {
	o := orm.NewOrm()
	board := Board{Id: id}

	if err := o.Read(&board); err != nil {
		return nil, errors.Wrap(err, "read fail")
	}

	beego.Info("Success to find board, id:%d, name:%s", id, board.Name)

	return &board, nil
}

func UpdateBoard(b *Board) error {
	o := orm.NewOrm()
	tmpForChecking := Board{Id: b.Id}

	if err := o.Read(&tmpForChecking); err != nil {
		return errors.Wrap(err, "read fail")
	}

	num, err := o.Update(b)

	if err != nil {
		return errors.Wrap(err, "update fail")
	}

	if num == 0 {
		return fmt.Errorf("Failed to update board, board not found, id:%d", b.Id)
	}

	beego.Info("Success to update board, id:%d, name:%s, category name:%s",
		b.Id, b.Name, b.BoardCategory.Name)

	return nil
}

func DeleteBoard(id int) error {
	o := orm.NewOrm()

	num, err := o.Delete(&Board{Id: id})
	if err != nil {
		return errors.Wrap(err, "delete fail")
	}

	if num == 0 {
		return errors.Errorf("Not found board for deleting board")
	}

	beego.Info("Success to delete board, id:%d", id)

	return nil
}

func SetBoardDeleteFlag(id int, isDeleted bool) error {
	o := orm.NewOrm()
	board := Board{Id: id}

	if err := o.Read(&board); err != nil {
		return errors.Wrap(err, "read fail")
	}

	board.IsDeleted = isDeleted
	if _, err := o.Update(&board); err != nil {
		return errors.Wrap(err, "update fail")
	}

	beego.Info("Success to set board close flag, id:%d, flag:%v", id, isDeleted)

	return nil
}

func GetBoardsByCategory(category *BoardCategory) ([]*Board, error) {
	o := orm.NewOrm()
	var bs []*Board

	num, err := o.QueryTable(new(Board)).Filter("BoardCategory__Id", category.Id).All(&bs)
	if err != nil {
		return nil, errors.Wrap(err, "query by category fail")
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
		return nil, errors.Wrap(err, "query all fail")
	}

	beego.Info("Success to get all board, board count:%d", num)

	return bs, nil
}
