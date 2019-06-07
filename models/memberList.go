package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var DefaultMemberList MemberList

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/ex_login?charset=utf8")

	DefaultMemberList = NewMySQLMemberList()
}

type MemberList interface {
	Add(member *Member) error
	Find(id string) (*Member, bool)
	GetAll() ([]*Member, error)
}

type MySQLMemberList struct {
	db orm.Ormer
}

func NewMySQLMemberList() *MySQLMemberList {
	var memberList MySQLMemberList
	memberList.db = orm.NewOrm()
	memberList.db.Using("default")
	return &memberList
}

func (m *MySQLMemberList) Add(member *Member) error {
	memberForCheck := Member{ID: member.ID}

	beego.Info("add, id:", member.ID, "password:", member.Password, "name:", member.Name)

	if err := m.db.Read(&memberForCheck); err == nil {
		return fmt.Errorf("Duplicated ID")
	}

	if _, err := m.db.Insert(member); err != nil {
		beego.Error("add, failed to insert. err:", err.Error())
		return err
	}

	beego.Info("add, success")

	return nil
}

func (m *MySQLMemberList) Find(id string) (*Member, bool) {
	foundMember := Member{ID: id}

	beego.Info("find, id:", id)

	if err := m.db.Read(&foundMember); err != nil {
		beego.Error("find, failed to read. err:", err.Error())
		return nil, false
	}

	beego.Info("find, success")

	return &foundMember, true
}

func (m *MySQLMemberList) GetAll() ([]*Member, error) {
	var members []*Member
	num, err := m.db.QueryTable("member").All(&members)

	if err != nil {
		beego.Error("Failed to get all members, err:" + err.Error())
		return members, err
	}
	beego.Info(fmt.Sprintf("Returned member count:%v", num))

	return members, nil
}
