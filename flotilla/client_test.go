package flotilla

import (
	"testing"

	"github.com/simulatedsimian/assert"
	"github.com/simulatedsimian/flotilla/dock"
)

type RequiredModules struct {
	Matrix
	Touch
	Number
	Dial
}

func TestAquire(t *testing.T) {
	assert := assert.Make(t)
	assert(true)

	e1, _ := dock.NewPipe().Endpoints()

	client, _ := ConnectToDocksRaw(e1)

	assert(client.structMembersToInterfaces(RequiredModules{})).Equal(
		[]interface{}{Matrix{}, Touch{}, Number{}, Dial{}},
	)
}
