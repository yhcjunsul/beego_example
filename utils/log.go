package utils

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func InitFileLog() {
	logFileName := fmt.Sprintf("log_%d_%d_%d.log", time.Now().Year(), time.Now().Month(), time.Now().Day())
	beego.SetLogger(logs.AdapterFile, fmt.Sprintf("{\"filename\":\"%s/%s.log\"}", "logs", logFileName))
}
