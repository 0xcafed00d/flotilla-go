package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Motor struct {
	ModuleCommon
}

func (m *Motor) Type() dock.ModuleType {
	return dock.Motor
}
