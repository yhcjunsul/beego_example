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

	//DefaultMemberList = NewInMemoryMemberList()
	DefaultMemberList = NewMySQLMemberList()
}

type MemberList interface {
	Add(member *Member) error
	Find(id string) (*Member, bool)
}

type InMemoryMemberList struct {
	members map[string]*Member
}

func NewInMemoryMemberList() *InMemoryMemberList {
	var memberList InMemoryMemberList
	memberList.members = make(map[string]*Member)
	return &memberList
}

func (m *InMemoryMemberList) Add(member *Member) error {
	if m.members[member.ID] != nil {
		return fmt.Errorf("Duplicated ID")
	}

	m.members[member.ID] = member
	return nil
}

func (m *InMemoryMemberList) Find(id string) (*Member, bool) {
	foundMember, ok := m.members[id]

	if ok == false {
		return nil, false
	}

	return foundMember, true
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
