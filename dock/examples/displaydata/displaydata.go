package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/simulatedsimian/flotilla-go/dock"
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

	time.Sleep(100 * time.Millisecond)
	d.SendDockCommand('v')
	time.Sleep(100 * time.Millisecond)
	d.SendDockCommand('e')

	for {
		ev := <-d.Events
		fmt.Println(ev)
		exitOnError(ev.Error)
	}
}
