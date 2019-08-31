package utils

import (
	"fmt"
	"sync"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yhcjunsul/beego_example/models"
)

var once sync.Once

func InitSql(aliasName string) {
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbname := beego.AppConfig.String("db.name")
	dbcharset := beego.AppConfig.String("db.charset")
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + dbcharset

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(aliasName, "mysql", dsn)
	if err := orm.RunSyncdb(aliasName, true, true); err != nil {
		fmt.Println(err)
	}

	orm.RunCommand()
}

func initTestSql() {
	orm.Debug = true
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8"

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("test", "mysql", dsn)
	orm.RegisterDataBase("default", "mysql", dsn)
	if err := orm.RunSyncdb("test", true, true); err != nil {
		fmt.Println(err)
	}

	orm.RunCommand()

	models.InitTestSetting()
}

func InitTestSql() {
	once.Do(func() {
		initTestSql()
	})
}
