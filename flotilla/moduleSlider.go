package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Slider struct {
	ModuleCommon
	value   int
	fupdate UpdateFunc
}

func (m *Slider) Construct() {
	m.fupdate = func(params []int) {
		m.value = Map(params[0], 0, 1023, 0, 1000)
	}
	m.OnUpdate(m.fupdate)
}

func (m *Slider) Type() dock.ModuleType {
	return dock.Slider
}

func (m *Slider) OnChange(f func(value int)) {
	m.OnUpdate(func(params []int) {
		m.fupdate(params)
		f(m.value)
	})
}

func (m *Slider) GetValue() int {
	return m.value
}
