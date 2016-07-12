package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Colour struct {
	ModuleCommon
	r, g, b, c int
	fupdate    UpdateFunc
}

func (m *Colour) Construct() {
	m.fupdate = func(params []int) {
		m.r, m.g, m.b, m.c = params[0], params[1], params[2], params[3]
	}
	m.OnUpdate(m.fupdate)
}

func (m *Colour) Type() dock.ModuleType {
	return dock.Colour
}
