package main

import (
	"log"

	"github.com/bodgix/gocmdline/actions"
	"github.com/bodgix/gocmdline/cmds"
)

func main() {
	arguments, err := cmds.ParseArgs()
	if err != nil {
		log.Fatal(err)
	}
	action, err := actions.GetAction(arguments)
	if err != nil {
		log.Fatal(err)
	}
	action.Do()
}
