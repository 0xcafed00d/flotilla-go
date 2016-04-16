package dock

import (
	"testing"

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
