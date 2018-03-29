package models

import (
	"github.com/astaxie/beego/orm"
)

type Rights struct {
	Id    int
	Rname string
}

func init() {
	orm.RegisterModel(new(Rights))
}

func FindAllRights() ([]Rights, error) {
	o := orm.NewOrm()
	var right []Rights
	_, err := o.QueryTable("rights").All(&right)
	return right, err
}
