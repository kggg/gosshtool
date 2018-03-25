package validate

import (
	"github.com/astaxie/beego/validation"
)

type HostInfo struct {
	Name   string `valid:"Required;MaxSize(100)"`
	Ipaddr string `valid:"Required;IP;MaxSize(100)"`
	User   string `valid:"Required;MaxSize(100)"`
	Pass   string `valid:"Required;MaxSize(150)"`
	Repass string `valid:"Required;MaxSize(150)"`
	Port   string `valid:"Required;MaxSize(100)"`
	Groups string `valid:"MaxSize(100)"`
}

func (u *HostInfo) Valid(v *validation.Validation) {
	if u.Pass != u.Repass {
		v.SetError("Password", "the password not same")
	}
}

type HostEdit struct {
	Name   string `valid:"Required;MaxSize(100)"`
	Ipaddr string `valid:"Required;IP;MaxSize(100)"`
	User   string `valid:"Required;MaxSize(100)"`
	Port   string `valid:"Required;MaxSize(100)"`
	Groups string `valid:"MaxSize(100)"`
}

type Groups struct {
	Gname string `valid:"Required;MaxSize(100)"`
	Info  string `valid:"MaxSize(300)"`
}

type UserAdd struct {
	Name  string `valid:"Required;MaxSize(100)"`
	Pass  string `valid:"Required;MaxSize(100)"`
	Pass2 string `valid:"Required;MaxSize(100)"`
	Email string `valid:"Required;Email;MaxSize(100)"`
}

func (u *UserAdd) Valid(v *validation.Validation) {
	if u.Pass != u.Pass2 {
		v.SetError("Password", "the password not same")
	}
}

type UserEdit struct {
	Email  string `valid:"Required;Email;MaxSize(100)"`
	Rights string `valid:"Required;Numeric;MaxSize(10)"`
}
