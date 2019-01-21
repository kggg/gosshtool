package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "steven:steven123@tcp(192.168.81.100:3306)/gosshtool?charset=utf8")
	//orm.Debug = true
}
