package validate

import (
	"github.com/astaxie/beego/validation"
)

/*
var MyMessage = map[string]string{
	"Required":     "不能为空",
	"Min":          "最小数值不能低于 %d",
	"Max":          "最大值不能大于 %d",
	"Range":        "取值范围在 %d 和 %d之间",
	"MinSize":      "Minimum size is %d",
	"MaxSize":      "Maximum size is %d",
	"Length":       "长度必须是 %d",
	"Alpha":        "Must be valid alpha characters",
	"Numeric":      "Must be valid numeric characters",
	"AlphaNumeric": "Must be valid alpha or numeric characters",
	"Match":        "必须匹配 %s",
	"NoMatch":      "Must not match %s",
	"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
	"Email":        "必须是正确的邮件地址",
	"IP":           "必须是正确的IP地址",
	"Base64":       "Must be valid base64 characters",
	"Mobile":       "Must be valid mobile number",
	"Tel":          "Must be valid telephone number",
	"Phone":        "Must be valid telephone or mobile phone number",
	"ZipCode":      "Must be valid zipcode",
}
*/

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
