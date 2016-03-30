package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	serialcfg := serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	port, err := serial.OpenPort(&serialcfg)

	exitOnError(err)

	log.Println("connecting to dock")
	d := dock.ConnectDock(port)

	time.Sleep(200 * time.Millisecond)
	d.SendDockCommand('e')

	numberModule := dock.Module{ModuleType: dock.Number, Dock: d}

	for {
		select {
		case ev := <-d.Events:
			numberModule.ProcessEvent(ev)
			fmt.Println(ev)
			exitOnError(ev.Error)

		case <-time.After(500 * time.Millisecond):

			if numberModule.Connected() {
				now := time.Now()
				hour := now.Hour()
				minute := now.Minute()
				second := now.Second()

				err := numberModule.Set(
					dock.GetDigitPattern(hour/10, false),
					dock.GetDigitPattern(hour%10, false),
					dock.GetDigitPattern(minute/10, false),
					dock.GetDigitPattern(minute%10, false),
					second%2)

				exitOnError(err)

			}
		}
	}
}
