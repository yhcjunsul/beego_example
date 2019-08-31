package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type BoardCategory struct {
	Id     int      `json:"id"`
	Name   string   `json:"name" orm:"size(100);unique"`
	Boards []*Board `orm:"reverse(many)"`
}

func AddCategory(c *BoardCategory) error {
	o := orm.NewOrm()
	category := BoardCategory{Name: c.Name}

	id, err := o.Insert(&category)

	if err != nil {
		return errors.Wrap(err, "insert fail")
	}

	*c = category
	beego.Info("Success to add category, id:%d, category name:%s", id, category.Name)

	return nil
}

func FindCategoryById(id int) (*BoardCategory, error) {
	o := orm.NewOrm()
	c := BoardCategory{Id: id}

	if err := o.Read(&c); err != nil {
		return nil, errors.Wrap(err, "Read fail")
	}

	beego.Info("Success to find board category by id, id:%d, name:%s", c.Id, c.Name)

	return &c, nil

}

func FindCategoryByName(name string) (*BoardCategory, error) {
	o := orm.NewOrm()
	c := BoardCategory{Name: name}

	if err := o.Read(&c, "Name"); err != nil {
		return nil, errors.Wrap(err, "Read fail")
	}

	beego.Info("Success to find board category by name, id:%d, name:%s", c.Id, name)

	return &c, nil
}

func UpdateCategory(c *BoardCategory) error {
	o := orm.NewOrm()
	tmpForChecking := BoardCategory{Id: c.Id}

	if err := o.Read(&tmpForChecking); err != nil {
		return errors.Wrap(err, "read fail")
	}

	num, err := o.Update(c)

	if err != nil {
		return errors.Wrap(err, "update fail")
	}

	if num == 0 {
		return fmt.Errorf("Failed to update board category, category not found, id:%d", c.Id)
	}

	beego.Info("Success to update board category, id:%d, name:%s", c.Id, c.Name)

	return nil
}

func DeleteCategoy(id int) error {
	o := orm.NewOrm()

	num, err := o.Delete(&BoardCategory{Id: id})
	if err != nil {
		return errors.Wrap(err, "delete fail")
	}

	if num == 0 {
		return errors.Errorf("Not found category for deleting category")
	}

	beego.Info("Success to delete board category, id:%d", id)

	return nil
}

func GetAllCategories() ([]*BoardCategory, error) {
	o := orm.NewOrm()
	var cs []*BoardCategory

	num, err := o.QueryTable(new(BoardCategory)).All(&cs)
	if err != nil {
		return nil, errors.Wrap(err, "query all fail")
	}

	beego.Info("Success to get all board category, category count:%d", num)

	return cs, nil
}
