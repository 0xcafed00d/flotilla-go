package flotilla

import "github.com/simulatedsimian/flotilla/dock"

type Updateable interface {
	Update(ev Event)
}

type Setable interface {
	Set(params ...int) error
}

type Typed interface {
	Type() dock.ModuleType
}
