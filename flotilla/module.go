package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type ModuleAddress struct {
	dock, channel int
}

type Module struct {
	ModuleAddress
	dock.ModuleType
	*Client
}

func (m *Module) Update(ev Event) {
	if ev.ModuleType == m.ModuleType {
		if ev.EventType == dock.EventConnected {
			m.ModuleAddress = ModuleAddress{ev.dockIndex, ev.Channel}
		}
		if ev.EventType == dock.EventDisconnected {
			m.ModuleAddress = ModuleAddress{-1, -1}
		}
	}
}

func (m *Module) Connected(ev Event) bool {
	return m.ModuleAddress.channel != -1
}
