package flotilla

import "github.com/simulatedsimian/flotilla-go/dock"

type Module interface {
	Update(ev Event)
	Connected() bool
	Type() dock.ModuleType
	Init(client *Client, t dock.ModuleType)
}

type Constructable interface {
	Construct()
}

type Setable interface {
	Set(d *dock.Dock) error
}

type AnalogValue interface {
	GetValue() int
}

type AnalogControllable interface {
	SetValue(value int)
}
