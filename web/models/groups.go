package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Groups))
}

type Groups struct {
	Id    int
	Gname string
	Info  string
}

func FindAllGroups() {

}
