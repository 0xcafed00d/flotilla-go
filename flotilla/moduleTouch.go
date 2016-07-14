package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Touch struct {
	ModuleCommon
	state     [4]bool
	prevState [4]bool

	fupdateDefault UpdateFunc
	fupdateUser    func(button int, pressed bool)
}

func (m *Touch) Construct() {
	m.fupdateDefault = func(params []int) {
		m.prevState = m.state
		for i := range params {
			m.state[i] = (params[i] != 0)
		}

		if m.fupdateUser != nil {
			for i := range params {
				if m.state[i] != m.prevState[i] {
					m.fupdateUser(i, m.state[i])
				}
			}
		}
	}
	m.OnUpdate(m.fupdateDefault)
}

func (m *Touch) Type() dock.ModuleType {
	return dock.Touch
}

func (m *Touch) OnChange(f func(button int, pressed bool)) {
	m.fupdateUser = f
}

func (m *Touch) GetValue() [4]bool {
	return m.state
}
