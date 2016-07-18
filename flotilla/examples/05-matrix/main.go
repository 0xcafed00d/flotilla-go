package main

import (
	"time"

	"github.com/simulatedsimian/flotilla-go/flotilla"
)

// build a struct that has all the modules you need.
var modules struct {
	flotilla.Matrix
	flotilla.Joystick
}

func main() {
	// connect to the dock
	client, err := flotilla.ConnectToDock("/dev/ttyACM0")
	flotilla.ExitOnError(err)

	// wait for all modules to be connected
	client.AquireModules(&modules)
	modules.Matrix.SetBrightness(16)

	counter := 0

	client.OnTick(func(t time.Time) {
		x, y, _ := modules.Joystick.GetValue()
		if x < 250 {
			modules.Matrix.ScrollLeft(counter)
			counter++
		}
		if x > 750 {
			modules.Matrix.ScrollRight(counter)
			counter++
		}
		if y < 250 {
			modules.Matrix.ScrollUp(counter)
			counter++
		}
		if y > 750 {
			modules.Matrix.ScrollDown(counter)
			counter++
		}

	})

	// go!!
	client.Run(time.Millisecond * 100)
}
