package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Dial struct {
	ModuleCommon
	value          int
	fupdateDefault UpdateFunc
	fupdateUser    func(value int)
}

func (m *Dial) Construct() {
	m.fupdateDefault = func(params []int) {
		m.value = Map(params[0], 0, 1023, 0, 1000)
		if m.fupdateUser != nil {
			m.fupdateUser(m.value)
		}
	}
	m.OnUpdate(m.fupdateDefault)
}

func (m *Dial) Type() dock.ModuleType {
	return dock.Dial
}

func (m *Dial) OnChange(f func(value int)) {
	m.fupdateUser = f
}

func (m *Dial) GetValue() int {
	return m.value
}
