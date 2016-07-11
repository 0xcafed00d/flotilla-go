package main

import (
	"time"

	"github.com/simulatedsimian/flotilla-go/flotilla"
)

// build a struct that has all the modules you need.
var modules struct {
	flotilla.Dial
	flotilla.Number
	flotilla.Rainbow
}

func main() {
	// connect to the dock
	client, err := flotilla.ConnectToDock("/dev/ttyACM0")
	flotilla.ExitOnError(err)

	// wait for all modules to be connected
	client.AquireModules(&modules)

	modules.OnChange(func(val int) {
		modules.Number.SetInteger(val)
		modules.Rainbow.SetVU(val)
	})

	// go!!
	client.Run(time.Millisecond * 100)
}
