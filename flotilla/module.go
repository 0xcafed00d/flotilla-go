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
	addr := ModuleAddress{ev.dockIndex, ev.Channel}
	if ev.ModuleType == m.moduleType {
		switch ev.EventType {
		case dock.EventConnected:
			m.address = &addr
		case dock.EventDisconnected:
			if m.Connected() && *m.address == addr {
				m.address = nil
			}
		case dock.EventUpdate:
			if m.updateFunc != nil {
				m.updateFunc(ev.Params)
			}
		}
	}
}

func (m *Module) Connected() bool {
	return m.address != nil
}
