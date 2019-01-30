package main

import (
	"gosshtool/bin/command"
	"gosshtool/utils/handler"
	"log"
)

func main() {
	cmd := command.ParseCommand()
	err := handler.HandleFunc(cmd)
	if err != nil {
		log.Fatal("Error: ", err)
	}

}
