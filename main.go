package main

import (
	"gosshtool/lib/command"
	"gosshtool/lib/sshclient"
	"log"
	"sync"
)

func main() {
	cmd := command.ParseCommand()
	var wg sync.WaitGroup
	if cmd.Module == "cmd" || cmd.Module == "sendfile" || cmd.Module == "getfile" {
		// execute a command on remote host
		for name, value := range cmd.Host {
			wg.Add(1)
			go func(host, user, pass string, port int, hostname string, module, cmd string) {
				sshClient := sshclient.New(user, pass, host, port, hostname)
				err := sshClient.Execute(module, cmd)
				if err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}(value.Ip, value.User, value.Pass, value.Port, name, cmd.Module, cmd.Action)
		}
		wg.Wait()
	}

	if cmd.Module == "ping" {
	}
	if cmd.Module == "list" {
	}

}
