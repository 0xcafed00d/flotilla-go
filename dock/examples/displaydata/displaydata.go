package main

import (
	"fmt"
	"os"

	"github.com/simulatedsimian/flotilla/dock"
	"github.com/tarm/serial"
)

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {

	serialcfg := serial.Config{Name: "/dev/ttyACM0", Baud: 115200}
	port, err := serial.OpenPort(&serialcfg)
	exitOnError(err)

	d := dock.ConnectDock(port)
	d.SendDockCommand('e')

	for {
		ev := <-d.Events
		fmt.Println(ev)
		exitOnError(ev.Error)

		if ev.EventType == dock.Update {
			d.SetModuleData(1, dock.Matrix, 1, 2, 3, 4, 5, 6, 7, 8, ev.Params[0]>>4)
		}

	}
}
