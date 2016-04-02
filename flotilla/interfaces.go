package flotilla

type Updateable interface {
	Update(ev Event)
}

type Setable interface {
	Set(params ...int) error
}
