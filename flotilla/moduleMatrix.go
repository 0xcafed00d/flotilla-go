package flotilla

import (
	"fmt"

	"github.com/simulatedsimian/flotilla-go/dock"
)

type MatrixBuffer struct {
	buffer [8]byte
	dirty  bool
}

type Matrix struct {
	ModuleCommon
	MatrixBuffer
	brightness byte
}

func (m *Matrix) Construct() {
	m.brightness = 64
	m.dirty = true
}

func (m *Matrix) String() (s string) {
	for i := range m.buffer {
		s += fmt.Sprintf("%08b\n", m.buffer[i])
	}
	return
}

func (m *Matrix) Set(d *dock.Dock) error {
	if m.dirty && m.address != nil {
		m.dirty = false
		return d.SetModuleData(m.address.channel, m.Type(),
			int(m.buffer[0]), int(m.buffer[1]), int(m.buffer[2]), int(m.buffer[3]),
			int(m.buffer[4]), int(m.buffer[5]), int(m.buffer[6]), int(m.buffer[7]),
			int(m.brightness))
	}
	return nil
}

func (m *Matrix) Type() dock.ModuleType {
	return dock.Matrix
}

func (m *Matrix) SetBrightness(b int) {
	m.brightness = byte(b)
	m.dirty = true
}

func (m *MatrixBuffer) Plot(col, row, v int) {
	col = 7 - col&7
	row = row & 7

	if v == 0 {
		m.buffer[col] &= ^(1 << uint(row))
	} else {
		m.buffer[col] |= (1 << uint(row))
	}

	m.dirty = true
}

func (m *MatrixBuffer) SetPattern(values []string) {
	for y := range values {
		for x, r := range values[y] {
			if r != ' ' {
				m.Plot(x, y, 1)
			} else {
				m.Plot(x, y, 0)
			}
		}
	}
}

func (m *MatrixBuffer) DrawBarGraph(values []int, min, max int) {
	m.Clear()

	bars := MinInt(8, len(values))

	for i := 0; i < bars; i++ {
		y := Map(values[i], min, max, 0, 7)
		for n := 0; n <= y; n++ {
			m.Plot(i, 7-n, 1)
		}
	}
}

func (m *MatrixBuffer) Clear() {
	m.buffer = [8]byte{}
}

func (m *MatrixBuffer) GetRow(row int) byte {
	var v byte
	for i := range m.buffer {
		v <<= 1
		v |= (m.buffer[i] >> byte(7-row)) & 1
	}
	return v
}

func (m *MatrixBuffer) SetRow(row int, v byte) {
	for i := range m.buffer {
		//mask := 1 << byte(7-row)
		v |= (m.buffer[i] >> byte(7-row)) & 1
	}
}

func (m *MatrixBuffer) GetCol(col int) byte {
	return m.buffer[col]
}

func (m *MatrixBuffer) SetCol(col int, v byte) {
	m.buffer[col] = v
}

func (m *MatrixBuffer) Scroll(dir Direction, fill byte) {
	if dir&DirLeft != 0 {
		m.ScrollLeft(fill)
	}
	if dir&DirRight != 0 {
		m.ScrollRight(fill)
	}
	if dir&DirUp != 0 {
		m.ScrollUp(fill)
	}
	if dir&DirDown != 0 {
		m.ScrollDown(fill)
	}
}

func (m *MatrixBuffer) ScrollRight(fill byte) {
	copy(m.buffer[:], m.buffer[1:])
	m.buffer[7] = fill
	m.dirty = true
}

func (m *MatrixBuffer) ScrollLeft(fill byte) {
	copy(m.buffer[1:], m.buffer[:])
	m.buffer[0] = fill
	m.dirty = true
}

func (m *MatrixBuffer) ScrollDown(fill byte) {
	for i := range m.buffer {
		m.buffer[i] = (m.buffer[i] << 1) | (fill>>byte(7-i))&1
	}
	m.dirty = true
}

func (m *MatrixBuffer) ScrollUp(fill byte) {
	for i := range m.buffer {
		m.buffer[i] = (m.buffer[i] >> 1) | ((byte(fill)>>byte(7-i))&1)<<7
	}
	m.dirty = true
}

func (m *MatrixBuffer) Roll(dir Direction) {
	if dir&DirLeft != 0 {
		m.RollLeft()
	}
	if dir&DirRight != 0 {
		m.RollRight()
	}
	if dir&DirUp != 0 {
		m.RollUp()
	}
	if dir&DirDown != 0 {
		m.RollDown()
	}
}

func (m *MatrixBuffer) RollUp() {
	row := m.GetRow(7)
	m.ScrollUp(row)
}

func (m *MatrixBuffer) RollDown() {
	row := m.GetRow(0)
	m.ScrollDown(row)
}

func (m *MatrixBuffer) RollLeft() {
	col := m.GetCol(7)
	m.ScrollLeft(col)
}

func (m *MatrixBuffer) RollRight() {
	col := m.GetCol(0)
	m.ScrollRight(col)
}
