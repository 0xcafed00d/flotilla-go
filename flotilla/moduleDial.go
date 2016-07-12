package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Dial struct {
	ModuleCommon
	value   int
	fupdate UpdateFunc
}

func (m *Dial) Construct() {
	m.fupdate = func(params []int) {
		m.value = Map(params[0], 0, 1023, 0, 1000)
	}
	m.OnUpdate(m.fupdate)
}

func (m *Dial) Type() dock.ModuleType {
	return dock.Dial
}

func (m *Dial) OnChange(f func(value int)) {
	m.OnUpdate(func(params []int) {
		m.fupdate(params)
		f(m.value)
	})
}

func (m *Dial) GetValue() int {
	return m.value
}
