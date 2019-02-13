package conf

import (
	"fmt"
	"testing"
)

func TestGetmode(t *testing.T) {
	filedir, err := Getworkdir()
	if err != nil {
		t.Error(err)
	}
	filename := filedir + "/conf.ini"
	cfg, err := Newconfig(filename)
	if err != nil {
		//fmt.Println(err)
		t.Error(err)
	}
	//fmt.Println(cfg)
	m := cfg.Getmode()
	fmt.Println(m)
}

func TestGetDBparams(t *testing.T) {
	filedir, err := Getworkdir()
	if err != nil {
		t.Error(err)
	}
	filename := filedir + "/conf.ini"
	cfg, err := Newconfig(filename)
	if err != nil {
		//fmt.Println(err)
		t.Error(err)
	}
	m := cfg.GetDBparams()
	fmt.Println(m)
}

func TestGetFileinfo(t *testing.T) {
	filedir, err := Getworkdir()
	if err != nil {
		t.Error(err)
	}
	filename := filedir + "/conf.ini"
	cfg, err := Newconfig(filename)
	if err != nil {
		t.Error(err)
	}
	m := cfg.GetFileinfo()
	fmt.Println(m)
}

func TestGetHostfile(t *testing.T) {
	filedir, err := Getworkdir()
	if err != nil {
		t.Error(err)
	}
	filename := filedir + "/conf.ini"
	cfg, err := Newconfig(filename)
	if err != nil {
		t.Error(err)
	}
	filehost, err := cfg.GetHostfile("/home/steven/go/src/gosshtool/bin/conf/host.conf")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(filehost)
}
