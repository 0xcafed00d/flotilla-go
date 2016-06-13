package main

import (
	"math/rand"

	"github.com/simulatedsimian/flotilla-go/dock"
)

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

func (l *LifeBoard) writeBoard(matrix *dock.Module) error {
	return matrix.Set(int(l[0]), int(l[1]), int(l[2]), int(l[3]),
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
