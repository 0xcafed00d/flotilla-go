package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Slider struct {
	ModuleCommon
}

func (m *Slider) Type() dock.ModuleType {
	return dock.Slider
}

func (m *Slider) OnChange(f func(value int)) {
	m.OnUpdate(func(params []int) {
		f(Map(params[0], 0, 1023, 0, 1000))
	})
}
