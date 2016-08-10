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
	modules.Matrix.SetBrightness(32)

	modules.Matrix.SetPattern([]string{
		"   OO   ",
		"  OOOO  ",
		" OOOOOO ",
		"OO OO OO",
		"OOOOOOOO",
		"  O  O  ",
		" O OO O ",
		"O O  O O"})

	client.OnTick(func(t time.Time) {
		dir, _ := modules.Joystick.GetDirection()
		if dir != flotilla.DirNone {
			modules.Matrix.Scroll(dir, 0)
		}
	})

	// go!!
	client.Run(time.Millisecond * 100)
}
