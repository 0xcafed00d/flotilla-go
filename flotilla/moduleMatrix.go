package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Matrix struct {
	ModuleCommon

	buffer     [8]byte
	brightness byte
	dirty      bool
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

func (m *Matrix) Plot(x, y, v int) {
	x = x & 7
	y = y & 7

	if v == 0 {
		m.buffer[x] = m.buffer[x] & ^(1 << uint(y))
	} else {
		m.buffer[x] = m.buffer[x] | (1 << uint(y))
	}

	m.dirty = true
}
