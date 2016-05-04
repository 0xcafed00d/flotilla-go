package flotilla

import (
	"testing"

	"github.com/simulatedsimian/assert"
)

func TestLimit(t *testing.T) {
	assert := assert.Make(t)

	assert(Limit(3, 0, 5)).Equal(3)
	assert(Limit(-1, 0, 5)).Equal(0)
	assert(Limit(6, 0, 5)).Equal(5)
}

func TestMap(t *testing.T) {
	assert := assert.Make(t)

	assert(Map(100, 0, 200, 0, 100)).Equal(50)
	assert(Map(100, 0, 200, -50, 50)).Equal(0)
	assert(Map(0, 0, 200, -50, 50)).Equal(-50)
	assert(Map(200, 0, 200, -50, 50)).Equal(50)
}

func TestMapDist(t *testing.T) {
	assert := assert.Make(t)

	var dist [8]int

	for val := 0; val < 1024; val++ {
		dist[Map(val, 0, 1023, 0, 7)]++
	}

	assert(dist).Equal([8]int{128, 128, 128, 128, 128, 128, 128, 128})
}

func TestMapDist2(t *testing.T) {
	assert := assert.Make(t)

	var dist [2]int

	for val := 0; val < 1024; val++ {
		dist[Map(val, 0, 1023, 0, 1)]++
	}

	assert(dist).Equal([2]int{512, 512})
}
