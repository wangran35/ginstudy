package models

import (
	_ "database/sql"
	"fmt"

	// "time"
	// _ "ginstudy/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//读数据
var orm *xorm.Engine

// //写数据
// var DB_Write *xorm.Engine

// var engine *xorm.Engine
var err error

func init() {
	orm, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/ginstudy?charset=utf8")
	// engine, err := xorm.NewEngine("mysql", "2343432:122222@/(http://127.0.0.1:3306)/ginstudy?charset=utf8")
	// db, err = xorm.NewEngine("mysql", "username:password@tcp(host:3306)/dbname?charset=utf8")
	if err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
		return
	} else {
		orm.Ping()
		//是否显示sql语句
		orm.ShowSQL(true)
		if err = orm.Sync2(new(Admin), new(User), new(Authgroup), new(Authrule)); err != nil {
			fmt.Println(err)
		} else {
			// orm.ShowWarn(true)
			// engine.ShowWarn(true)
			fmt.Print("自动生成表成功！")
		}

	}

	// fmt.Println(err)
}
