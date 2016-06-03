package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Matrix struct {
	ModuleCommon
}

func (m *Matrix) Type() dock.ModuleType {
	return dock.Matrix
}
