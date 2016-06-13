package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Motion struct {
	ModuleCommon
}

func (m *Motion) Type() dock.ModuleType {
	return dock.Motion
}
