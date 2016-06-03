package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Barometer struct {
	ModuleCommon
}

func (m *Barometer) Type() dock.ModuleType {
	return dock.Barometer
}
