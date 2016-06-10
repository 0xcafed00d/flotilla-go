package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Weather struct {
	ModuleCommon
}

func (m *Weather) Type() dock.ModuleType {
	return dock.Weather
}
