package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zvodd/example-go-command/internal/gui"
	"github.com/zvodd/example-go-command/util"
)

func main() {
	var useGUI = flag.Bool("gui", false, "flag for gui launch")
	flag.Parse()

	port, err := util.FindFreePort()
	if err != nil {
		log.Fatalln(err)
	}

	if *useGUI {
		gui.StartGUI(port)
	} else {
		fmt.Println("Found free port:", port)
	}
}
