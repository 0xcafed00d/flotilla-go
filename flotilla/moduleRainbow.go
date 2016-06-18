package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Rainbow struct {
	ModuleCommon

	colours [5]RGB
	dirty   bool
	allsame bool
}

func (m *Rainbow) Type() dock.ModuleType {
	return dock.Rainbow
}

func (m *Rainbow) Set(d *dock.Dock) error {
	if m.dirty {
		if m.allsame {
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

func (m *Rainbow) SetAll(rgb RGB) {
	m.dirty = true
	m.allsame = true
	m.colours[0] = rgb
}

func (m *Rainbow) SetBlend(rgb1, rgb2 RGB) {
	m.dirty = true
	m.allsame = false
	m.colours[0] = rgb1
	m.colours[4] = rgb2

	m.colours[2] = Blend(rgb1, rgb2)
	m.colours[1] = Blend(rgb1, m.colours[2])
	m.colours[3] = Blend(rgb2, m.colours[2])
}

func (m *Rainbow) SetBlend3(rgb1, rgb2, rgb3 RGB) {
	m.dirty = true
	m.allsame = false
	m.colours[0] = rgb1
	m.colours[2] = rgb2
	m.colours[4] = rgb3

	m.colours[1] = Blend(rgb1, rgb2)
	m.colours[3] = Blend(rgb2, rgb3)
}
