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
	flotilla.Dial
}

func main() {
	// connect to the dock
	client, err := flotilla.ConnectToDock("/dev/ttyACM0")
	flotilla.ExitOnError(err)

	// wait for all modules to be connected
	client.AquireModules(&modules)

	usage := cpuusage.Usage{}
	modules.SetBrightness(2)

	client.OnTick(func(t time.Time) {
		if err := usage.Measure(); err != nil {
			log.Println(err)
		} else {
			modules.Matrix.Clear()
			for i, v := range usage.Cores {
				modules.Matrix.Plot(i, flotilla.Map(v, 0, 100, 0, 7), 1)
			}
		}
	})

	// go!!
	client.Run(time.Millisecond * 100)
}
