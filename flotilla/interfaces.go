package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Updateable interface {
	Update(ev Event)
	Connected() bool
}

type Setable interface {
	Set(params ...int) error
}

type Typed interface {
	Type() dock.ModuleType
}
