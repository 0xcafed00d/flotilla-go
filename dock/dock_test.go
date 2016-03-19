package dock

import (
	"io"
	"io/ioutil"
	"strings"
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

	assert(msgToEvent("c 1/joystick")).Equal(Event{
		EventType:  Connected,
		ModuleType: Joystick,
		Port:       1,
	})

	assert(msgToEvent("d 1/joystick")).Equal(Event{
		EventType:  Disconnected,
		ModuleType: Joystick,
		Port:       1,
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

type Source struct {
	io.Reader
	io.Writer
}

func TestMsgRec(t *testing.T) {
	assert := assert.Make(t)
	s := Source{strings.NewReader("c 1/joystick\ru 1/joystick 1,234,874\rd 1/joystick\r"), ioutil.Discard}
	d := ConnectDock(s)

	assert(<-d.Events).Equal(Event{
		EventType:  Connected,
		ModuleType: Joystick,
		Port:       1,
	})

	assert(<-d.Events).Equal(Event{
		EventType:  Update,
		ModuleType: Joystick,
		Port:       1,
		Params:     []int{1, 234, 874},
	})

	assert(<-d.Events).Equal(Event{
		EventType:  Disconnected,
		ModuleType: Joystick,
		Port:       1,
	})

	assert(<-d.Events).Equal(Event{
		EventType: Error,
		Error:     io.EOF,
	})
}
