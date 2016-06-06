package main

import (
	"time"

	"github.com/simulatedsimian/flotilla/flotilla"
)

// build a struct that has all the modules you need.
var modules struct {
	flotilla.Matrix
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
	modules.Matrix.Plot(0, 4, 1)
	modules.Matrix.Plot(1, 4, 1)
	modules.Matrix.Plot(2, 4, 1)
	modules.Matrix.SetBrightness(12)

	flicker := 0

	client.OnTick(func(t time.Time) {
		modules.ScrollUp(flicker)
		flicker++
	})

	// go!!
	client.Run(time.Millisecond * 50)
}
