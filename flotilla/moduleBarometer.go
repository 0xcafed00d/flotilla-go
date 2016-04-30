package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Barometer struct {
	Module
}

func (m *Barometer) Type() dock.ModuleType {
	return dock.Barometer
}
