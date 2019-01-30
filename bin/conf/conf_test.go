package conf

import (
	"fmt"
	"testing"
)

//const filepath = "/bin/conf/conf.ini"

func TestGetmode(t *testing.T) {
	cfg, err := Newconfig(filepath)
	if err != nil {
		//fmt.Println(err)
		t.Error(err)
	}
	//fmt.Println(cfg)
	m := cfg.Getmode()
	fmt.Println(m)
}

func TestGetDBparams(t *testing.T) {
	cfg, err := Newconfig(filepath)
	if err != nil {
		//fmt.Println(err)
		t.Error(err)
	}
	m := cfg.GetDBparams()
	fmt.Println(m)
}

func TestGetFileinfo(t *testing.T) {
	cfg, err := Newconfig(filepath)
	if err != nil {
		t.Error(err)
	}
	m := cfg.GetFileinfo()
	fmt.Println(m)
}
