package flotilla

import (
	"testing"

	"github.com/simulatedsimian/assert"
)

type ev struct {
	button  int
	pressed bool
}

func TestUpdateFunc(t *testing.T) {
	assert := assert.Make(t)

	touch := Touch{}
	touch.Construct()

	var events []ev

	touch.OnChange(func(button int, pressed bool) {
		events = append(events, ev{button, pressed})
	})

	touch.fupdateDefault([]int{1, 0, 0, 0})
	assert(events).Equal([]ev{ev{0, true}})

	events = nil
	touch.fupdateDefault([]int{0, 1, 1, 0})
	assert(events).Equal([]ev{ev{0, false}, ev{1, true}, ev{2, true}})

	events = nil
	touch.fupdateDefault([]int{0, 0, 0, 1})
	assert(events).Equal([]ev{ev{1, false}, ev{2, false}, ev{3, true}})

}
