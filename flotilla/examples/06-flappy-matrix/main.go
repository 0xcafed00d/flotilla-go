package main

import (
	"fmt"
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
	heights   [8]int
}

func (gs *gameState) init(m *flotilla.Matrix) {
	gs.playerPos = 256 * 3

	for x := 0; x < 8; x++ {
		m.Plot(x, 7, 1)
		gs.heights[x] = 1
	}

}

func main() {
	// connect to the dock
	client, err := flotilla.ConnectToDock("/dev/ttyACM0")
	flotilla.ExitOnError(err)

	// wait for all modules to be connected
	client.AquireModules(&modules)
	modules.Matrix.SetBrightness(32)

	gs := gameState{}
	gs.init(&modules.Matrix)

	client.OnTick(func(t time.Time) {
		modules.Matrix.Plot(0, gs.playerPos/256, 0)

		if gs.scrollPos&7 == 0 {
			modules.Matrix.ScrollLeft(0)
			copy(gs.heights[0:], gs.heights[1:])
			if gs.scrollPos&31 == 0 {
				gs.currentH = rand.Intn(4) + 1
			}
			modules.Matrix.Plot(7, 7, 1)
			gs.heights[7] = gs.currentH
			draw(&modules.Matrix, gs.currentH)
			fmt.Println(gs.heights)
		}
		gs.scrollPos++

		maxDelta := 64
		gs.deltaPos = flotilla.Limit(gs.deltaPos, -maxDelta, maxDelta)

		gs.playerPos += gs.deltaPos
		gs.deltaPos += 12
		if gs.playerPos > (7-gs.heights[0])*256 {
			gs.playerPos = (7 - gs.heights[0]) * 256
		}
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
