package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type ModuleAddress struct {
	dock, index int
}

type Module struct {
	ModuleAddress
	dock.ModuleType
	*Client
}

func (m *Module) Update(ev Event) {
	if ev.ModuleType == m.ModuleType {
		if ev.EventType == dock.Connected {
			m.ModuleAddress = ModuleAddress{ev.dockIndex, ev.Port}
		}
		if ev.EventType == dock.Disconnected {
			m.ModuleAddress = ModuleAddress{-1, -1}
		}
	}
}

func (m *Module) Connected(ev Event) bool {
	return m.ModuleAddress.index != -1
}
