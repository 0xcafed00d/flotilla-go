package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Joystick struct {
	ModuleCommon
}

func (m *Joystick) Type() dock.ModuleType {
	return dock.Joystick
}
