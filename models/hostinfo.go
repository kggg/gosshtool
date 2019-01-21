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
	Skey       int
	Hostgroups *Hostgroups `orm:"rel(fk)"`
	Created_at string
}

type HostAll struct {
	Hostinfo
	Gname string
}

func init() {
	orm.RegisterModel(new(Hostinfo))
}

func FindAllHostinfo() ([]HostAll, error) {
	var hostinfo []HostAll
	o := orm.NewOrm()
	//_, err := o.QueryTable("hostinfo").RelatedSel().OrderBy("-id").Limit(5, start).All(&hostinfo)
	sql := "select hostinfo.id,hostinfo.ip,hostinfo.user,hostinfo.pass,hostinfo.port,hostinfo.name, hostinfo.skey,hostgroups.gname from hostinfo left join hostgroups on hostgroups.id=hostinfo.hostgroups_id"
	_, err := o.Raw(sql).QueryRows(&hostinfo)
	return hostinfo, err
}

func FindHostByName(str string) (Hostinfo, error) {
	o := orm.NewOrm()
	var hostinfo Hostinfo
	err := o.QueryTable("hostinfo").Filter("name", str).One(&hostinfo)
	//err := o.QueryTable("hostinfo").Filter("name", str).RelatedSel("hostgroups").One(&hostinfo)
	return hostinfo, err
}

func FindHostById(id int) (Hostinfo, error) {
	o := orm.NewOrm()
	var hostinfo Hostinfo
	err := o.QueryTable("hostinfo").Filter("Id", id).RelatedSel("hostgroups").One(&hostinfo)
	return hostinfo, err
}

func FindHostByGroupname(gname string) ([]HostAll, error) {
	var hostinfo []HostAll
	o := orm.NewOrm()
	sql := "select hostinfo.id,hostinfo.ip,hostinfo.user,hostinfo.pass,hostinfo.port,hostinfo.name, hostgroups.gname from hostinfo inner join hostgroups on hostgroups.id=hostinfo.hostgroups_id where hostgroups.gname=?"
	_, err := o.Raw(sql, gname).QueryRows(&hostinfo)
	return hostinfo, err
}

func FindHostByIp(str string) (Hostinfo, error) {
	o := orm.NewOrm()
	var hostinfo Hostinfo
	err := o.QueryTable("hostinfo").Filter("Ip", str).One(&hostinfo)
	return hostinfo, err
}

func AddHost(name, ip, user, pass string, port int, group int) (int64, error) {
	o := orm.NewOrm()
	sql := "insert into hostinfo (name, ip, user, pass,port,hostgroups_id) values( ?, ?, ?, ?,?,?)"
	res, err := o.Raw(sql, name, ip, user, pass, port, group).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func EditHost(name, ip, user string, port int, group int, id int) (int64, error) {
	o := orm.NewOrm()
	sql := "update hostinfo  set name=?,ip=?, user=?,port=?, hostgroups_id=? where id=?"
	res, err := o.Raw(sql, name, ip, user, port, group, id).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func ChangeHostPass(pass string, id int) (int64, error) {
	o := orm.NewOrm()
	sql := "update hostinfo  set pass=? where id=?"
	res, err := o.Raw(sql, pass, id).Exec()
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
