package main

import (
	"log"
	"time"

	"github.com/simulatedsimian/cpuusage"
	"github.com/simulatedsimian/flotilla-go/flotilla"
)

// build a struct that has all the modules you need.
var modules struct {
	flotilla.Matrix
	flotilla.Number
	flotilla.Rainbow
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
			modules.Matrix.DrawBarGraph(usage.Cores, 0, 100)
			modules.Number.SetInteger(usage.Overall)
			modules.Rainbow.SetBlend(flotilla.RGB{255, 0, 0}, flotilla.RGB{0, 0, 255})
		}
	})

	// go!!
	client.Run(time.Millisecond * 250)
}
