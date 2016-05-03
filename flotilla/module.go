package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type ModuleAddress struct {
	dock, channel int
}

type Module struct {
	address    *ModuleAddress
	moduleType dock.ModuleType
	client     *Client
	updateFunc UpdateFunc
}

type UpdateFunc func(params []int)

func (m *Module) OnUpdate(f UpdateFunc) {
	m.updateFunc = f
}

func (m *Module) Update(ev Event) {
	if ev.ModuleType == m.moduleType {
		if ev.EventType == dock.EventConnected {
			m.address = &ModuleAddress{ev.dockIndex, ev.Channel}
		} else if ev.EventType == dock.EventDisconnected {
			m.address = nil
		} else {
			if m.updateFunc != nil {
				m.updateFunc(ev.Params)
			}
		}
	}
}

func (m *Module) Connected() bool {
	return m.address != nil
}
