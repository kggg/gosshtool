package validate

import (
	"github.com/astaxie/beego/validation"
	"strings"
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
	u.Name = strings.TrimSpace(u.Name)
	u.Ipaddr = strings.TrimSpace(u.Ipaddr)
	u.User = strings.TrimSpace(u.User)
	u.Pass = strings.TrimSpace(u.Pass)
	u.Repass = strings.TrimSpace(u.Repass)
	u.Port = strings.TrimSpace(u.Port)
}
