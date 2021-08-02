package main

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "go-admin/initialize/conf"
	"go-admin/initialize/mysql"
	_ "go-admin/initialize/session"
	_ "go-admin/routers"
	_ "go-admin/utils/template"
)

func main() {

	//输出文件名和行号
	logs.SetLogFuncCall(true)
	//启动beego
	web.Run()
	mysql.CloseDB()
}
