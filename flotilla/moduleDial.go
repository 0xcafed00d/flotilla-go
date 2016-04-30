package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Dial struct {
	Module
}

func (m *Dial) Type() dock.ModuleType {
	return dock.Dial
}
