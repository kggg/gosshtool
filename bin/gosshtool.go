package main

import (
	"gosshtool/utils/command"
	"gosshtool/utils/funcmap"
	"log"
)

func main() {
	cmd := command.ParseCommand()
	err := funcmap.Related(cmd)
	if err != nil {
		log.Fatal("Error: ", err)
	}

}
