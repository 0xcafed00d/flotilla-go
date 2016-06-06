package main

import (
	"time"

	"github.com/simulatedsimian/flotilla/flotilla"
)

// build a struct that has all the modules you need.
var modules struct {
	flotilla.Matrix
	flotilla.Dial
}

// set up change handlers for the modules
func init() {
}

func main() {
	// connect to the dock
	client, err := flotilla.ConnectToDock("/dev/ttyACM0")
	flotilla.ExitOnError(err)

	// wait for all modules to be connected
	client.AquireModules(&modules)

	client.OnTick(func(t time.Time) {
		modules.Matrix.ScrollLeft(0)
		y := flotilla.Map(modules.Dial.GetValue(), 0, 1000, 0, 7)
		modules.Matrix.Plot(7, y, 1)
	})

	// go!!
	client.Run(time.Millisecond * 20)
}
