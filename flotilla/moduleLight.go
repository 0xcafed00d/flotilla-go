package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Light struct {
	Module
}

func (m *Light) Type() dock.ModuleType {
	return dock.Light
}
