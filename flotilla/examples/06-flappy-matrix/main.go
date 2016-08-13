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

func drawPipes(m *flotilla.Matrix) {
	gap := rand.Intn(7)
	for y := 0; y < gap; y++ {
		m.Plot(7, y, 1)
	}
	for y := gap + 2; y < 8; y++ {
		m.Plot(7, y, 1)
	}
}

type gameState struct {
	playerPos int
	deltaPos  int
	scrollPos int
	score     int
	hiscore   int
}

func main() {
	// connect to the dock
	client, err := flotilla.ConnectToDock("/dev/ttyACM0")
	flotilla.ExitOnError(err)

	// wait for all modules to be connected
	client.AquireModules(&modules)
	modules.Matrix.SetBrightness(32)

	gs := gameState{playerPos: 16 * 3}

	client.OnTick(func(t time.Time) {
		modules.Matrix.Plot(0, gs.playerPos/16, 0)

		if gs.scrollPos&7 == 0 {
			modules.Matrix.ScrollLeft(0)
			if gs.scrollPos&31 == 0 {
				drawPipes(&modules.Matrix)
			}
		}
		gs.scrollPos++

		gs.playerPos += gs.deltaPos
		gs.deltaPos++
		modules.Matrix.Plot(0, gs.playerPos/16, gs.scrollPos&1)
	})

	modules.Touch.OnChange(func(button int, pressed bool) {
		if pressed {
			gs.deltaPos -= 8
		}
	})

	// go!!
	client.Run(time.Millisecond * 50)
}
