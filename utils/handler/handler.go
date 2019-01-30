package handler

import (
	"errors"
	"gosshtool/bin/command"
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
func (c Handle) Cmd(com *command.Command) {
	c.ConnRemote(com)
}

//复制和发送文件
func (c Handle) Copy(com *command.Command) {
	c.ConnRemote(com)
}

func (c Handle) ConnRemote(com *command.Command) {
	var wg sync.WaitGroup
	for _, value := range com.Host {
		wg.Add(1)
		go func(ipaddr, user, pass string, port int, hostname string, skey int, module, act string) {
			depass, err := msgcrypt.AesDecrypt(pass)
			if err != nil {
				log.Fatal(err)
			}
			var sskey bool
			if skey == 1 {
				sskey = true
			} else {
				sskey = false
			}
			sshClient, err := sshclient.NewClient(ipaddr, user, depass, port, sskey, hostname)
			if err != nil {
				log.Fatal(err)
			}
			err = sshClient.Execute(module, act)
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}(value.Ip, value.User, value.Pass, value.Port, value.Name, value.Skey, com.Module, com.Action)
	}
	wg.Wait()
}

func HandleFunc(com *command.Command) error {
	fn := NewHandle()
	err := fn.AddHandleFunc(com.Module, com)
	if err != nil {
		return err
	}
	return nil
}
