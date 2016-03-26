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

type LifeBoard [8]byte

func (l *LifeBoard) Set(x, y, v int) {
	x = x & 7
	y = y & 7

	if v == 0 {
		l[x] = l[x] & ^(1 << uint(y))
	} else {
		l[x] = l[x] | (1 << uint(y))
	}
}

func (l *LifeBoard) Get(x, y int) int {
	x = x & 7
	y = y & 7

	if l[x]&(1<<uint(y)) != 0 {
		return 1
	}
	return 0
}

func (l *LifeBoard) writeBoard(port int, d *dock.Dock) error {
	return d.SetModuleData(port, dock.Matrix, int(l[0]), int(l[1]), int(l[2]), int(l[3]),
		int(l[4]), int(l[5]), int(l[6]), int(l[7]), 32)
}

func (l *LifeBoard) randomPopulation(rng *rand.Rand) {
	for i := range l {
		l[i] = byte(rng.Uint32())
	}
}

func (l *LifeBoard) clear() {
	for i := range l {
		l[i] = 0
	}
}

func (l *LifeBoard) fill() {
	for i := range l {
		l[i] = 255
	}
}

func (l *LifeBoard) makeGlider() {
	l.clear()
	l.Set(3, 3, 1)
	l.Set(4, 4, 1)
	l.Set(4, 5, 1)
	l.Set(3, 5, 1)
	l.Set(2, 5, 1)
}

func (l *LifeBoard) countNeighbours(xx, yy int) int {
	cnt := 0
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x != 0 || y != 0 {
				cnt += l.Get(x+xx, y+yy)
			}
		}
	}
	return cnt
}

func (l *LifeBoard) generation() {
	var dest LifeBoard
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			cnt := l.countNeighbours(x, y)

			if l.Get(x, y) == 1 {
				if cnt < 2 || cnt > 3 {
					dest.Set(x, y, 0)
				} else {
					dest.Set(x, y, 1)
				}
			} else {
				if cnt == 3 {
					dest.Set(x, y, 1)
				}
			}
		}
	}
	*l = dest
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

	matrixIdx := -1

	for {
		select {
		case ev := <-d.Events:
			if ev.ModuleType == dock.Matrix {
				if ev.EventType == dock.Connected {
					matrixIdx = ev.Port
				}
				if ev.EventType == dock.Disconnected {
					matrixIdx = -1
				}
			}
			if ev.ModuleType == dock.Touch && ev.EventType == dock.Update {
				if ev.Params[0] == 1 {
					board.randomPopulation(rng)
					board.writeBoard(matrixIdx, d)
				}
				if ev.Params[1] == 1 {
					board.makeGlider()
					board.writeBoard(matrixIdx, d)
				}
			}

			fmt.Println(ev)
			exitOnError(ev.Error)

		case <-time.After(100 * time.Millisecond):
			if matrixIdx != -1 {
				err := board.writeBoard(matrixIdx, d)
				exitOnError(err)
				board.generation()
			}
		}
	}
}
