package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Colour struct {
	ModuleCommon
}

func (m *Colour) Type() dock.ModuleType {
	return dock.Colour
}
