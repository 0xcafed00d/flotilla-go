package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Motion struct {
	Module
}

func (m *Motion) Type() dock.ModuleType {
	return dock.Motion
}
