package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Module interface {
	Update(ev Event)
	Connected() bool
	Type() dock.ModuleType
	Init(client *Client, t dock.ModuleType)
}

type Setable interface {
	Set(params ...int) error
}
