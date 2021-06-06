package mysql

import (
	"fmt"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	"go-admin/global"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// init 注册mysql
func init() {
	debug, err := beego.AppConfig.Bool("debug")
	if err == nil {
		orm.Debug = debug
	}

	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		beego.Error("mysql register driver error:", err)
	}

	//dataSource := "root:root@tcp(127.0.0.1:3306)/test"
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		global.BA_CONFIG.Mysql.Username,
		global.BA_CONFIG.Mysql.Password,
		global.BA_CONFIG.Mysql.Host,
		global.BA_CONFIG.Mysql.Port,
		global.BA_CONFIG.Mysql.Database,
	)

	err = orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		beego.Error("mysql register database error:", err)
	}
}
