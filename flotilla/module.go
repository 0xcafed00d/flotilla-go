package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type ModuleAddress struct {
	dock, channel int
}

type ModuleCommon struct {
	address    *ModuleAddress
	moduleType dock.ModuleType
	client     *Client
	updateFunc UpdateFunc
}

func (m *ModuleCommon) Init(client *Client, t dock.ModuleType) {
	m.client = client
	m.moduleType = t
}

type UpdateFunc func(params []int)

func (m *ModuleCommon) OnUpdate(f UpdateFunc) {
	m.updateFunc = f
}

func (m *ModuleCommon) Update(ev Event) {
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

func (m *ModuleCommon) Connected() bool {
	return m.address != nil
}
