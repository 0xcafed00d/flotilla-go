package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Updateable interface {
	Update(ev Event)
	Connected() bool
	Type() dock.ModuleType
}

type Setable interface {
	Set(params ...int) error
}
