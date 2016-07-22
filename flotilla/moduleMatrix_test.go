package flotilla

import (
	"testing"

	"github.com/simulatedsimian/assert"
)

func TestSetRowCol(t *testing.T) {
	assert := assert.Make(t)

	var m Matrix
	m.buffer = [8]byte{1, 2, 4, 8, 16, 32, 64, 128}
	assert(m.GetCol(0)).Equal(byte(1))
	assert(m.GetCol(4)).Equal(byte(16))
	assert(m.GetCol(7)).Equal(byte(128))
	assert(m.GetRow(0)).Equal(byte(1))
	assert(m.GetRow(4)).Equal(byte(16))
	assert(m.GetRow(7)).Equal(byte(128))
}
