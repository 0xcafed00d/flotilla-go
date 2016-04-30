package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Joystick struct {
	Module
}

func (m *Joystick) Type() dock.ModuleType {
	return dock.Joystick
}
