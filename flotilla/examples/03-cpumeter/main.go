package main

import (
	"log"
	"time"

	"github.com/simulatedsimian/cpuusage"
	"github.com/simulatedsimian/flotilla/flotilla"
)

// build a struct that has all the modules you need.
var modules struct {
	flotilla.Matrix
	flotilla.Touch
}

func main() {
	// connect to the dock
	client, err := flotilla.ConnectToDock("/dev/ttyACM0")
	flotilla.ExitOnError(err)

	// wait for all modules to be connected
	client.AquireModules(&modules)

	usage := cpuusage.Usage{}
	modules.SetBrightness(2)

	modules.Touch.OnChange(func(button int, pressed bool) {
		log.Println(button, pressed)
	})

	client.OnTick(func(t time.Time) {
		if err := usage.Measure(); err != nil {
			log.Println(err)
		} else {
			modules.Matrix.DrawBarGraph(usage.Cores, 0, 100)
		}
	})

	// go!!
	client.Run(time.Millisecond * 250)
}
