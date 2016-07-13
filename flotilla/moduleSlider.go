package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Slider struct {
	ModuleCommon
	value          int
	fupdateDefault UpdateFunc
	fupdateUser    func(value int)
}

func (m *Slider) Construct() {
	m.fupdateDefault = func(params []int) {
		m.value = Map(params[0], 0, 1023, 0, 1000)
		if m.fupdateUser != nil {
			m.fupdateUser(m.value)
		}
	}
	m.OnUpdate(m.fupdateDefault)
}

func (m *Slider) Type() dock.ModuleType {
	return dock.Slider
}

func (m *Slider) OnChange(f func(value int)) {
	m.fupdateUser = f
}

func (m *Slider) GetValue() int {
	return m.value
}
