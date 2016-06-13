package main

import (
	"fmt"
	"time"

	"github.com/simulatedsimian/flotilla-go/flotilla"
)

// build a struct that has all the modules you need.
var modules struct {
	flotilla.Dial
}

// set up change handlers for the modules
func init() {
	modules.Dial.OnChange(func(val int) {
		fmt.Println(val)
	})
}

func main() {
	// connect to the dock
	client, err := flotilla.ConnectToDock("/dev/ttyACM0")
	flotilla.ExitOnError(err)

	// wait for all modules to be connected
	client.AquireModules(&modules)

	client.OnTick(func(t time.Time) {
		fmt.Println(t)
	})

	// go!!
	client.Run(time.Millisecond * 50)
}
