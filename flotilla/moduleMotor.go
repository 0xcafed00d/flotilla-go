package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Motor struct {
	Module
}

func (m *Motor) Type() dock.ModuleType {
	return dock.Motor
}
