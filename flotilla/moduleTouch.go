package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Touch struct {
	ModuleCommon
	state [4]bool
}

func (m *Touch) Type() dock.ModuleType {
	return dock.Touch
}

func (m *Touch) OnChange(f func(button int, pressed bool)) {
	m.OnUpdate(func(params []int) {
		for i := range params {
			if m.state[i] != (params[i] != 0) {
				m.state[i] = (params[i] != 0)
				f(i, (params[i] != 0))
			}
		}
	})
}

func (m *Touch) GetValue() [4]bool {
	return m.state
}
