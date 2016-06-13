package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Joystick struct {
	ModuleCommon
}

func (m *Joystick) Type() dock.ModuleType {
	return dock.Joystick
}
