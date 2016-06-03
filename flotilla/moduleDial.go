package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Dial struct {
	ModuleCommon
}

func (m *Dial) Type() dock.ModuleType {
	return dock.Dial
}

func (m *Dial) OnChange(f func(value int)) {
	m.OnUpdate(func(params []int) {
		f(Map(params[0], 0, 1023, 0, 1000))
	})
}
