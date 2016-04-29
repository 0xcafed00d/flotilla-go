package flotilla

import (
	"testing"

	"github.com/simulatedsimian/assert"
	"github.com/simulatedsimian/flotilla/dock"
)

type RequiredModules struct {
	M1 Matrix
	M2 Matrix
	Touch
	Number
	Dial
}

func TestAquire(t *testing.T) {
	mustPanic := assert.MustPanic
	assert := assert.Make(t)

	modules := RequiredModules{}

	assert(structMembersToInterfaces(&modules)).Equal(
		[]interface{}{&Matrix{}, &Matrix{}, &Touch{}, &Number{}, &Dial{}},
	)

	mustPanic(t, func(t *testing.T) {
		structMembersToInterfaces(modules)
	})

	mustPanic(t, func(t *testing.T) {
		structMembersToInterfaces(0)
	})
}

func TestConnectDisconnect(t *testing.T) {
	assert := assert.Make(t)

	e1, e2 := dock.NewPipe().Endpoints()

	client, _ := ConnectToDocksRaw(e1)
	sim := dock.NewSimulator(e2)

	var modules RequiredModules

	assert(modules.M1.Connected()).Equal(false)

	client.AquireModules(&modules)

	sim.Connect(dock.Matrix, 3)

	assert(modules.M1.Connected()).Equal(true)

	e1.Close()
}
