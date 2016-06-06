package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Dial struct {
	ModuleCommon
	value int
}

func (m *Dial) Construct() {
	m.OnUpdate(func(params []int) {
		m.value = Map(params[0], 0, 1023, 0, 1000)
	})
}

func (m *Dial) Type() dock.ModuleType {
	return dock.Dial
}

func (m *Dial) OnChange(f func(value int)) {
	m.OnUpdate(func(params []int) {
		m.value = Map(params[0], 0, 1023, 0, 1000)
		f(m.value)
	})
}

func (m *Dial) GetValue() int {
	return m.value
}
