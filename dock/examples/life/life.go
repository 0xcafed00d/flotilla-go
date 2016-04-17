package main

import (
	"fmt"
	"log"
	"math/rand"
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
	var board LifeBoard

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	board.randomPopulation(rng)

	serialcfg := serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	port, err := serial.OpenPort(&serialcfg)

	exitOnError(err)

	log.Println("connecting to dock")
	d := dock.ConnectDock(port)

	time.Sleep(200 * time.Millisecond)
	d.SendDockCommand('e')

	matrixModule := dock.Module{ModuleType: dock.Matrix, Dock: d}
	numberModule := dock.Module{ModuleType: dock.Number, Dock: d}

	gen := 0

	for {
		select {
		case ev := <-d.Events:
			matrixModule.ProcessEvent(ev)
			numberModule.ProcessEvent(ev)

			if ev.ModuleType == dock.Touch && ev.EventType == dock.EventUpdate {
				if ev.Params[0] == 1 {
					board.randomPopulation(rng)
					board.writeBoard(&matrixModule)
					gen = 0
				}
				if ev.Params[1] == 1 {
					board.makeGlider()
					board.writeBoard(&matrixModule)
					gen = 0
				}
			}

			fmt.Println(ev)
			exitOnError(ev.Error)

		case <-time.After(100 * time.Millisecond):
			if matrixModule.Connected() {
				if numberModule.Connected() {
					err := numberModule.Set(
						dock.GetDigitPattern((gen/1000)%10, false),
						dock.GetDigitPattern((gen/100)%10, false),
						dock.GetDigitPattern((gen/10)%10, false),
						dock.GetDigitPattern(gen%10, false))
					exitOnError(err)
				}

				err := board.writeBoard(&matrixModule)
				exitOnError(err)
				board.generation()
				gen++
			}
		}
	}
}
