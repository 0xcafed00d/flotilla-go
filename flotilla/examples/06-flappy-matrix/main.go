package main

import (
	"math/rand"
	"time"

	"github.com/simulatedsimian/flotilla-go/flotilla"
)

// build a struct that has all the modules you need.
var modules struct {
	flotilla.Matrix
	flotilla.Touch
	flotilla.Number
}

func draw(m *flotilla.Matrix, h int) {
	for y := 0; y < h; y++ {
		m.Plot(7, 7-y, 1)
	}
}

type gameState struct {
	playerPos int
	deltaPos  int
	scrollPos int
	score     int
	hiscore   int
	currentH  int
}

func main() {
	// connect to the dock
	client, err := flotilla.ConnectToDock("/dev/ttyACM0")
	flotilla.ExitOnError(err)

	// wait for all modules to be connected
	client.AquireModules(&modules)
	modules.Matrix.SetBrightness(32)

	gs := gameState{playerPos: 256 * 3}

	client.OnTick(func(t time.Time) {
		modules.Matrix.Plot(0, gs.playerPos/256, 0)

		if gs.scrollPos&7 == 0 {
			modules.Matrix.ScrollLeft(0)
			if gs.scrollPos&31 == 0 {
				gs.currentH = rand.Intn(5)
			}
			modules.Matrix.Plot(7, 7, 1)
			draw(&modules.Matrix, gs.currentH)
		}
		gs.scrollPos++

		maxDelta := 32
		gs.deltaPos = flotilla.Limit(gs.deltaPos, -maxDelta, maxDelta)

		gs.playerPos += gs.deltaPos
		gs.deltaPos += 12
		modules.Matrix.Plot(0, gs.playerPos/256, gs.scrollPos&1)
	})

	modules.Touch.OnChange(func(button int, pressed bool) {
		if pressed {
			gs.deltaPos -= 98
		}
	})

	// go!!
	client.Run(time.Millisecond * 50)
}
