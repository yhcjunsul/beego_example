package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BoardCategory struct {
	Id   int16  `orm:"pk"`
	Name string `orm:"size(100);unique"`
}

func init() {
	orm.RegisterModel(new(BoardCategory))
}

func AddCategory(c *BoardCategory) error {
	o := orm.NewOrm()
	category := BoardCategory{Name: c.Name}

	if isCreated, id, err := o.ReadOrCreate(&category, "Name"); err != nil {
		return err
	}

	if isCreated == false {
		return fmt.Errorf("Category name already exists, id:%d, name:%s", id, name)
	}

	beego.Info("Success to add category, id:" + id + ", category name:" + name)

	return nil
}

func FindCategoryByName(name string) (*BoardCategory, error) {
	o := orm.NewOrm()
	c := BoardCategory{Name: name}

	if err := o.Read(&c, "Name"); err != nil {
		return nil, err
	}

	beego.Info("Success to find board category, id:" + c.Id + ", name:" + name)

	return &c, nil
}

func UpdateCategory(c *BoardCategory) error {
	o := orm.NewOrm()
	tmpForChecking := BoardCategory{Id: c.Id}

	if err := o.Read(&tmpForChecking); err != nil {
		return err
	}

	if num, err := o.Update(c); err != nil {
		return err
	}

	if num == 0 {
		return fmt.Errorf("Failed to update board category, category not found, id:%d", c.Id)
	}

	beego.Info("Success to update board category, id:" + c.Id + ", name:" + c.Name)

	return nil
}

func DeleteCategoy(id int) error {
	o := orm.NewOrm()
	if _, err := o.Delete(&BoardCategory{Id: id}); err != nil {
		return err
	}

	beego.Info("Success to delete board category, id:" + id)

	return nil
}

func GetAllCategories() ([]*BoardCategory, error) {
	o := orm.NewOrm()
	var cs []*BoardCategory

	num, err := o.QueryTable(new(BoardCategory)).All(&cs)
	if err != nil {
		return nil, err
	}

	beego.Info("Success to get all board category, category count:" + num)

	return cs, nil
}
