package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Number struct {
	ModuleCommon

	buffer [4]int
	dirty  bool
}

func (m *Number) Type() dock.ModuleType {
	return dock.Number
}

func (m *Number) Set(d *dock.Dock) error {
	if m.dirty && m.address != nil {
		m.dirty = false
		return d.SetModuleData(m.address.channel, m.Type(),
			int(m.buffer[0]), int(m.buffer[1]), int(m.buffer[2]), int(m.buffer[3]), 0)
	}
	return nil
}

func (m *Number) SetInteger(i int) {
	m.dirty = true
	if i > 9999 {
		i = 9999
	}

	m.buffer[3] = dock.GetDigitPattern(i%10, false)
	i = i / 10
	m.buffer[2] = dock.GetDigitPattern(i%10, false)
	i = i / 10
	m.buffer[1] = dock.GetDigitPattern(i%10, false)
	i = i / 10
	m.buffer[0] = dock.GetDigitPattern(i%10, false)
}
