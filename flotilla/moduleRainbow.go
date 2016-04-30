package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Rainbow struct {
	Module
}

func (m *Rainbow) Type() dock.ModuleType {
	return dock.Rainbow
}
