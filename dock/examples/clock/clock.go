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

const (
	DecPnt = 1
	HMid   = 2
	VLTop  = 4
	VLBot  = 8
	HBot   = 16
	VRBot  = 32
	VRTop  = 64
	HTop   = 128
)

var Digits = []int{
	HTop | HBot | VLBot | VLTop | VRBot | VRTop,
	VRBot | VRTop,
	HTop | HBot | HMid | VLBot | VRTop,
	HTop | HBot | HMid | VRBot | VRTop,
	HMid | VRTop | VRTop | VLTop | VRBot,
	HTop | HBot | HMid | VRBot | VLTop,
	HTop | HBot | HMid | VRBot | VLTop | VLBot,
	HTop | VRBot | VRTop | VRBot,
	HTop | HMid | HBot | VLBot | VLTop | VRBot | VRTop,
	HTop | HMid | HBot | VLTop | VRBot | VRTop,
}

func main() {

	serialcfg := serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	port, err := serial.OpenPort(&serialcfg)

	exitOnError(err)

	log.Println("connecting to dock")
	d := dock.ConnectDock(port)

	time.Sleep(200 * time.Millisecond)
	d.SendDockCommand('e')

	numberIdx := -1

	for {
		select {
		case ev := <-d.Events:
			if ev.ModuleType == dock.Number {
				if ev.EventType == dock.Connected {
					numberIdx = ev.Port
				}
				if ev.EventType == dock.Disconnected {
					numberIdx = -1
				}
			}

			fmt.Println(ev)
			exitOnError(ev.Error)

		case <-time.After(500 * time.Millisecond):

			if numberIdx != -1 {
				now := time.Now()
				hour := now.Hour()
				minute := now.Minute()
				second := now.Second()

				err := d.SetModuleData(numberIdx, dock.Number,
					Digits[hour/10], Digits[hour%10],
					Digits[minute/10], Digits[minute%10],
					second%2)

				exitOnError(err)

			}
		}
	}
}
