package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int
	Name       string
	Pass       string
	Email      string
	Rights     *Rights `orm:"rel(fk)"`
	Created_at string
}

func init() {
	orm.RegisterModel(new(User))
}

func FindAllUser() ([]User, error) {
	o := orm.NewOrm()
	var user []User
	_, err := o.QueryTable("user").RelatedSel("rights").All(&user)
	return user, err
}

func FindUserByName(name string) (User, error) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("name", name).One(&user)
	return user, err
}

func FindUserById(id int) (User, error) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("id", id).RelatedSel("rights").One(&user)
	return user, err
}

func AddUser(name, pass, email string) (int64, error) {
	o := orm.NewOrm()
	sql := "insert into user (name, pass,email) values( ?, ?, ?)"
	res, err := o.Raw(sql, name, pass, email).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func EditUser(email string, right, id int) (int64, error) {
	o := orm.NewOrm()
	sql := "update user  set email=?,rights_id=? where id=?"
	res, err := o.Raw(sql, email, right, id).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func ChangeUserPass(pass string, id int) (int64, error) {
	o := orm.NewOrm()
	sql := "update user set pass=? where id=?"
	res, err := o.Raw(sql, pass, id).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func DeleteUser(id int) (int64, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&User{Id: id}); err == nil {
		return num, err
	} else {
		return 0, err
	}
}

func UserExistCheck(name string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("user").Filter("name", name).Exist()
	return exist
}
