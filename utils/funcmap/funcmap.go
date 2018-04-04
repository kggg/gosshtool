package funcmap

import (
	"errors"
	"gosshtool/utils/command"
	"gosshtool/utils/msgcrypt"
	"gosshtool/utils/sshclient"
	"log"
	"reflect"
	"sync"
)

var (
	ErrParamsNotAdapted = errors.New("The number of params is not adapted.")
)

type Funcs map[string]reflect.Value

func NewFuncs(size int) Funcs {
	return make(Funcs, size)
}

func (f Funcs) Bind(name string, fn interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(name + " is not callable.")
		}
	}()
	v := reflect.ValueOf(fn)
	v.Type().NumIn()
	f[name] = v
	return
}

func (f Funcs) Call(name string, params ...interface{}) (result []reflect.Value, err error) {
	if _, ok := f[name]; !ok {
		err = errors.New("The key [" + name + "] does not exist.")
		return
	}
	if len(params) != f[name].Type().NumIn() {
		err = ErrParamsNotAdapted
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f[name].Call(in)
	return
}

func (f Funcs) RegisterFunc() {
	err := f.Bind("cmd", cmd)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Bind("copyfile", copyfile)
	if err != nil {
		log.Fatal(err)
	}

}

func Related(com command.Command) error {
	fn := NewFuncs(3)
	fn.RegisterFunc(com.Module)
	_, err := fn.Call(com.Module, com)
	if err != nil {
		return err
	}
	return nil
}

//执行远程命令
func cmd(com command.Command) {
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
func copyfile(com command.Command) {
}
