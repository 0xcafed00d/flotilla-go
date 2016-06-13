package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Light struct {
	ModuleCommon
}

func (m *Light) Type() dock.ModuleType {
	return dock.Light
}
