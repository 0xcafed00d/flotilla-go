package main

import (
	"fmt"

	"github.com/simulatedsimian/flotilla/flotilla"
)

// build a struct that has all the modules you need.
var modules struct {
	//flotilla.Dial
	flotilla.Touch
	flotilla.Rainbow
}

// set up change handlers for the modules
func init() {
	/*
		modules.Dial.OnChange(func(val int) {
			fmt.Println(val)
		})
	*/
}

func main() {
	// connect to the dock
	client, err := flotilla.ConnectToDock("/dev/ttyACM0")
	flotilla.ExitOnError(err)

	// wait for all modules to be connected
	client.AquireModules(&modules)
	fmt.Println("waiting for modules")
	client.WaitForModules(&modules)
	fmt.Println("modules connected")

	// go!!
	client.Run()
}
