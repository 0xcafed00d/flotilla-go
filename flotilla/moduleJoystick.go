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
	x = m.valueX
	y = m.valueY
	button = m.button
	return
}
