package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Touch struct {
	ModuleCommon
}

func (m *Touch) Type() dock.ModuleType {
	return dock.Touch
}
