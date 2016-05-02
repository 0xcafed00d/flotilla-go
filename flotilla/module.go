package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type ModuleAddress struct {
	dock, channel int
}

type Module struct {
	address    *ModuleAddress
	moduleType dock.ModuleType
	client     *Client
}

func (m *Module) Update(ev Event) {
	if ev.ModuleType == m.moduleType {
		if ev.EventType == dock.EventConnected {
			m.address = &ModuleAddress{ev.dockIndex, ev.Channel}
		}
		if ev.EventType == dock.EventDisconnected {
			m.address = nil
		}
	}
}

func (m *Module) Connected() bool {
	return m.address != nil
}
