package handler

import (
	"errors"
	"gosshtool/utils/command"
	"gosshtool/utils/msgcrypt"
	"gosshtool/utils/sshclient"
	"log"
	"reflect"
	"strings"
	"sync"
)

type Handle map[string]reflect.Value

func NewHandle() Handle {
	return make(Handle, 0)
}

//添加命令行模块对应的执行函数
func (c Handle) AddHandleFunc(name string, f interface{}) error {
	vf := reflect.ValueOf(c)
	vft := vf.Type()
	mNum := vf.NumMethod()
	for i := 0; i < mNum; i++ {
		methodname := strings.ToLower(vft.Method(i).Name)
		if name == methodname {
			c[name] = vf.Method(i)
		}
	}
	if _, ok := c[name]; ok {
		parms := []reflect.Value{reflect.ValueOf(f)}
		c[name].Call(parms)
	} else {
		return errors.New("The module name [" + name + "] invalid")
	}
	return nil
}

//执行远程命令
func (c Handle) Cmd(com command.Command) {
	var wg sync.WaitGroup
	for _, value := range com.Host {
		wg.Add(1)
		go func(ipaddr, user, pass string, port int, hostname string, module, c string) {
			depass, err := msgcrypt.AesDecrypt(pass)
			if err != nil {
				log.Fatal(err)
			}
			sshClient := sshclient.New(ipaddr, user, depass, port, hostname)
			err = sshClient.Execute(module, c)
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}(value.Ip, value.User, value.Pass, value.Port, value.Name, com.Module, com.Action)
	}
	wg.Wait()
}

//复制和发送文件
func (c Handle) Copyfile(com command.Command) {
	log.Println(com.Module)
}

func HandleFunc(com command.Command) error {
	fn := NewHandle()
	err := fn.AddHandleFunc(com.Module, com)
	if err != nil {
		return err
	}
	return nil
}
