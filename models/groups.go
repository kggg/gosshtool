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

func FindAllGroups() ([]Groups, error) {
	o := orm.NewOrm()
	var groups []Groups
	_, err := o.QueryTable("groups").All(&groups)
	return groups, err
}

func FindGroupsById(id int) (Groups, error) {
	o := orm.NewOrm()
	var groups Groups
	err := o.QueryTable("groups").Filter("id", id).One(&groups)
	return groups, err
}

func FindGroupsByName(gname string) ([]Groups, error) {
	o := orm.NewOrm()
	var groups []Groups
	_, err := o.QueryTable("groups").Filter("gname", gname).All(&groups)
	return groups, err
}

func AddGroups(name, info string) (int64, error) {
	o := orm.NewOrm()
	sql := "insert into groups (gname, info) values( ?, ?)"
	res, err := o.Raw(sql, name, info).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func EditGroups(name, info string, id int) (int64, error) {
	o := orm.NewOrm()
	sql := "update groups  set gname=?,info=? where id=?"
	res, err := o.Raw(sql, name, info, id).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func DeleteGroups(id int) (int64, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Groups{Id: id}); err == nil {
		return num, err
	} else {
		return 0, err
	}
}

func GroupsExistCheck(name string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("groups").Filter("gname", name).Exist()
	return exist
}
