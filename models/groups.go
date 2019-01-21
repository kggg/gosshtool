package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Hostgroups))
}

type Hostgroups struct {
	Id    int
	Gname string
	Info  string
}

func FindAllGroups() ([]Hostgroups, error) {
	o := orm.NewOrm()
	var groups []Hostgroups
	_, err := o.QueryTable("hostgroups").All(&groups)
	return groups, err
}

func FindGroupsById(id int) (Hostgroups, error) {
	o := orm.NewOrm()
	var groups Hostgroups
	err := o.QueryTable("hostgroups").Filter("id", id).One(&groups)
	return groups, err
}

func FindGroupsByName(gname string) ([]Hostgroups, error) {
	o := orm.NewOrm()
	var groups []Hostgroups
	_, err := o.QueryTable("hostgroups").Filter("gname", gname).All(&groups)
	return groups, err
}

func AddGroups(name, info string) (int64, error) {
	o := orm.NewOrm()
	sql := "insert into hostgroups (gname, info) values( ?, ?)"
	res, err := o.Raw(sql, name, info).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func EditGroups(name, info string, id int) (int64, error) {
	o := orm.NewOrm()
	sql := "update hostgroups  set gname=?,info=? where id=?"
	res, err := o.Raw(sql, name, info, id).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func DeleteGroups(id int) (int64, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Hostgroups{Id: id}); err == nil {
		return num, err
	} else {
		return 0, err
	}
}

func GroupsExistCheck(name string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("hostgroups").Filter("gname", name).Exist()
	return exist
}
