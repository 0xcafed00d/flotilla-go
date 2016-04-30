package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Slider struct {
	Module
}

func (m *Slider) Type() dock.ModuleType {
	return dock.Slider
}
