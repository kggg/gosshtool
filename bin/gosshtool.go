package main

import (
	"gosshtool/utils/command"
	"gosshtool/utils/handler"
	"log"
)

func main() {
	cmd := command.ParseCommand()
	err := handler.Related(cmd)
	if err != nil {
		log.Fatal("Error: ", err)
	}

}
