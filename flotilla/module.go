package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Module struct {
	dock.ModuleType
	*dock.Dock
	index int
}

func (m *Module) Update(ev Event) {
	if ev.ModuleType == m.ModuleType {
		if ev.EventType == dock.Connected {
			m.index = ev.Port
		}
		if ev.EventType == dock.Disconnected {
			m.index = -1
		}
	}
}

func (m *Module) Connected(ev Event) bool {
	return m.index != -1
}
