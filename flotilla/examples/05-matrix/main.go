package main

import (
	"fmt"
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

	modules.Matrix.Plot(0, 0, 1)
	modules.Matrix.Plot(0, 2, 1)
	fmt.Println(modules.Matrix.String())

	client.OnTick(func(t time.Time) {
		dir, _ := modules.Joystick.GetDirection()
		if dir != flotilla.DirNone {
			modules.Matrix.Scroll(dir, counter)
			counter++
		}
	})

	// go!!
	client.Run(time.Millisecond * 100)
}
