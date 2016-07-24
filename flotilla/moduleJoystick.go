package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Joystick struct {
	ModuleCommon

	valueX int
	valueY int
	button bool

	fupdateDefault UpdateFunc
	fupdateUser    func(valueX int, valueY int, button bool)
}

func (m *Joystick) Construct() {
	m.fupdateDefault = func(params []int) {
		m.valueX = Map(params[1], 0, 1023, 0, 1000)
		m.valueY = Map(params[2], 0, 1023, 0, 1000)
		m.button = IntToBool(params[0])

		if m.fupdateUser != nil {
			m.fupdateUser(m.valueX, m.valueY, m.button)
		}
	}
	m.OnUpdate(m.fupdateDefault)
}

func (m *Joystick) Type() dock.ModuleType {
	return dock.Joystick
}

func (m *Joystick) OnChange(f func(valueX int, valueY int, button bool)) {
	m.fupdateUser = f
}

func (m *Joystick) GetValue() (x, y int, button bool) {
	if !m.Connected() {
		x, y = 500, 500
		return
	}
	x = m.valueX
	y = m.valueY
	button = m.button
	return
}

func (m *Joystick) GetDirection() (dir Direction, button bool) {
	if !m.Connected() {
		dir = DirNone
		return
	}
	if m.valueX < 250 {
		dir = dir | DirLeft
	} else if m.valueX > 750 {
		dir = dir | DirRight
	}

	if m.valueY < 250 {
		dir = dir | DirUp
	} else if m.valueY > 750 {
		dir = dir | DirDown
	}

	button = m.button
	return
}
