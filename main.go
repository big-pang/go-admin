package main

import (
	beego "github.com/beego/beego/v2/adapter"
	_ "go-admin/initialize/conf"
	_ "go-admin/initialize/mysql"
	_ "go-admin/initialize/session"
	_ "go-admin/routers"
	_ "go-admin/utils/template"
)

func main() {

	//输出文件名和行号
	beego.SetLogFuncCall(true)

	//启动beego
	beego.Run()
}
