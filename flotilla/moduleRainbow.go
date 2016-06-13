package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Rainbow struct {
	ModuleCommon
}

func (m *Rainbow) Type() dock.ModuleType {
	return dock.Rainbow
}
