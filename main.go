package main

import (
	"github.com/astaxie/beego"

	"github.com/yhcjunsul/beego_example/models"
	_ "github.com/yhcjunsul/beego_example/routers"
	"github.com/yhcjunsul/beego_example/utils"
)

func main() {
	utils.InitSql()
	models.InitTestSetting()
	beego.Run()
}
