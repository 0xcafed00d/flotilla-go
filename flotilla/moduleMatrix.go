package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Matrix struct {
	Module
}

func (m *Matrix) Type() dock.ModuleType {
	return dock.Matrix
}
