package dock

import (
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/simulatedsimian/assert"
)

func TestParseRequest(t *testing.T) {
	assert := assert.Make(t)

	assert(msgToRequest("e")).Equal(Request{RequestType: ReqEnquire})
	assert(msgToRequest("r")).Equal(Request{RequestType: ReqResetToBootloader})
	assert(msgToRequest("v")).Equal(Request{RequestType: ReqVersion})
	assert(msgToRequest("d")).Equal(Request{RequestType: ReqDebug})
	assert(msgToRequest("p 0")).Equal(Request{RequestType: ReqPower, Params: []int{0}})
	assert(msgToRequest("p 1")).Equal(Request{RequestType: ReqPower, Params: []int{1}})
	assert(msgToRequest("p")).Equal(Request{})
	assert(msgToRequest("p 3")).Equal(Request{})

	assert(msgToRequest("n d hello")).Equal(Request{RequestType: ReqName, Params: []int{'d'}, ParamStr: "hello"})
	assert(msgToRequest("n u world")).Equal(Request{RequestType: ReqName, Params: []int{'u'}, ParamStr: "world"})

	assert(msgToRequest("s 5 1,2,3")).Equal(Request{RequestType: ReqSet, Channel: 5, Params: []int{1, 2, 3}})
}

func TestSimulatorRequest(t *testing.T) {
	assert := assert.Make(t)

	e1, e2 := NewPipe().Endpoints()
	sim := NewSimulator(e1)

	fmt.Fprint(e2, "e\r")
	assert(<-sim.Requests).Equal(Request{RequestType: ReqEnquire})
	e2.Close()
	assert(<-sim.Requests).Equal(Request{RequestType: ReqError, Error: io.EOF})

}

func TestSimulatorConnectDisconnect(t *testing.T) {
	assert := assert.Make(t)

	buffer := make([]byte, 128)
	e1, e2 := NewPipe().Endpoints()
	sim := NewSimulator(e1)

	assert(sim.modules[2]).Equal(Unknown)
	sim.Connect(Matrix, 3)
	assert(sim.modules[2]).Equal(Matrix)
	time.Sleep(100 * time.Millisecond)
	n, _ := e2.Read(buffer)
	assert(string(buffer[:n])).Equal("c 3/matrix\r\n")

	sim.Disconnect(3)
	assert(sim.modules[2]).Equal(Unknown)
	time.Sleep(100 * time.Millisecond)
	n, _ = e2.Read(buffer)
	assert(string(buffer[:n])).Equal("d 3\r\n")

	e2.Close()
}
