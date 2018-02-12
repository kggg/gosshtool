package models

import (
	"github.com/astaxie/beego/orm"
)

type Hostinfo struct {
	Id         int
	Name       string
	Ip         string
	User       string
	Pass       string
	Port       int
	Groups     string
	Created_at string
}

func init() {
	orm.RegisterModel(new(Hostinfo))
}

func FindAllHostinfo() ([]Hostinfo, error) {
	var hostinfo []Hostinfo
	o := orm.NewOrm()
	//_, err := o.QueryTable("hostinfo").RelatedSel().OrderBy("-id").Limit(5, start).All(&hostinfo)
	_, err := o.QueryTable("hostinfo").RelatedSel().OrderBy("-id").All(&hostinfo)
	return hostinfo, err
}

func FindHostByName(str string) (Hostinfo, error) {
	o := orm.NewOrm()
	var hostinfo Hostinfo
	err := o.QueryTable("hostinfo").Filter("name", str).One(&hostinfo)
	return hostinfo, err
}

func FindHostByIp(str string) (Hostinfo, error) {
	o := orm.NewOrm()
	var hostinfo Hostinfo
	err := o.QueryTable("hostinfo").Filter("Ip", str).One(&hostinfo)
	return hostinfo, err
}

func AddHost(name, ip, user, pass string, port int, group string) (int64, error) {
	o := orm.NewOrm()
	sql := "insert into hostinfo (name, ip, user, pass,port,groups) values( ?, ?, ?, ?,?,?)"
	res, err := o.Raw(sql, name, ip, user, pass, port, group).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func UpdateHost(name, ip, user, pass string, port int, group string, id int) (int64, error) {
	o := orm.NewOrm()
	sql := "update hostinfo  set name=?,ipaddr=?, user=?, pass=?,port=?, groups=? where id=?"
	res, err := o.Raw(sql, name, ip, user, pass, port, group, id).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func DeleteHost(id int) (int64, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Hostinfo{Id: id}); err == nil {
		return num, err
	} else {
		return 0, err
	}
}

func NameExistCheck(name string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("hostinfo").Filter("name", name).Exist()
	return exist
}
