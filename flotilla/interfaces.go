package flotilla

type Updateable interface {
	Update(ev Event)
}
