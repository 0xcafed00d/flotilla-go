package dock

import (
	"testing"

	"github.com/simulatedsimian/assert"
)

func TestMsgSplitter(t *testing.T) {
	assert := assert.Make(t)

	msgSplit := makeMessageSplitter()

	assert(msgSplit([]byte("msg1\rmsg2\r"))).Equal([]string{"msg1", "msg2"})
	assert(msgSplit([]byte("msg3\rmsg"))).Equal([]string{"msg3"})
	assert(msgSplit([]byte("4\r"))).Equal([]string{"msg4"})
}

func TestMsgParser(t *testing.T) {
	assert := assert.Make(t)

	assert(msgToEvent("u 1/joystick 1,234,874")).Equal(Event{
		EventType:  Update,
		ModuleType: Joystick,
		Port:       1,
		Params:     []int{1, 234, 874},
	})

	assert(msgToEvent("u 1/xxxx 1,234,874")).Equal(Event{
		EventType:  Update,
		ModuleType: Unknown,
		Port:       1,
		Params:     []int{1, 234, 874},
	})

	assert(msgToEvent("p 1/xxxx 1,234,874")).Equal(Event{
		EventType:  Invalid,
		ModuleType: Unknown,
		Port:       1,
		Params:     []int{1, 234, 874},
	})
}
