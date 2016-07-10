package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Rainbow struct {
	ModuleCommon

	colours [5]RGB
	dirty   bool
	allSame bool
}

func (m *Rainbow) Type() dock.ModuleType {
	return dock.Rainbow
}

func (m *Rainbow) Set(d *dock.Dock) error {
	if m.dirty {
		if m.allSame {
			return d.SetModuleData(m.address.channel, m.Type(),
				int(m.colours[0].R), int(m.colours[0].G), int(m.colours[0].B))
		} else {
			return d.SetModuleData(m.address.channel, m.Type(),
				int(m.colours[0].R), int(m.colours[0].G), int(m.colours[0].B),
				int(m.colours[1].R), int(m.colours[1].G), int(m.colours[1].B),
				int(m.colours[2].R), int(m.colours[2].G), int(m.colours[2].B),
				int(m.colours[3].R), int(m.colours[3].G), int(m.colours[3].B),
				int(m.colours[4].R), int(m.colours[4].G), int(m.colours[4].B))
		}
		m.dirty = false
	}
	return nil
}

func (m *Rainbow) SetSame(rgb RGB) {
	m.dirty = true
	m.allSame = true
	m.colours[0] = rgb
}

func (m *Rainbow) SetAll(rgb [5]RGB) {
	m.dirty = true
	m.allSame = true
	m.colours = rgb
}

func (m *Rainbow) SetBlend(rgb1, rgb2 RGB) {
	m.dirty = true
	m.allSame = false
	m.colours[0] = rgb1
	m.colours[4] = rgb2

	m.colours[2] = Blend(rgb1, rgb2)
	m.colours[1] = Blend(rgb1, m.colours[2])
	m.colours[3] = Blend(rgb2, m.colours[2])
}

func (m *Rainbow) SetBlend3(rgb1, rgb2, rgb3 RGB) {
	m.dirty = true
	m.allSame = false
	m.colours[0] = rgb1
	m.colours[2] = rgb2
	m.colours[4] = rgb3

	m.colours[1] = Blend(rgb1, rgb2)
	m.colours[3] = Blend(rgb2, rgb3)
}

func (m *Rainbow) SetVU(value int) {
	m.dirty = true
	m.allSame = false

	colour := LerpRGB(RGB{0, 255, 0}, RGB{255, 0, 0}, float64(value)/1000.0)

	for i := range m.colours {
		if i*200 < value {
			m.colours[i] = LerpRGB(RGB{}, colour, float64(value-i*200)/1000.0)
		} else {
			m.colours[i] = RGB{}
		}
	}
}
