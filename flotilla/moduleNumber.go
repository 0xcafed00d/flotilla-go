package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Number struct {
	ModuleCommon
}

func (m *Number) Type() dock.ModuleType {
	return dock.Number
}
