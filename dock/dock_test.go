package dock

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/simulatedsimian/assert"
)

func TestJoin(t *testing.T) {
	assert := assert.Make(t)

	assert(join([]int{}, ",")).Equal("")
	assert(join([]int{1}, ",")).Equal("1")
	assert(join([]int{1, 2}, ",")).Equal("1,2")
	assert(join([]int{1, 2, 3}, ",")).Equal("1,2,3")
	assert(join([]int{1, 2, 3, 4}, ", ")).Equal("1, 2, 3, 4")
}

func TestMsgSplitter(t *testing.T) {
	assert := assert.Make(t)

	msgSplit := makeMessageSplitter()

	assert(msgSplit([]byte("msg1\r\nmsg2\r\n"))).Equal([]string{"msg1", "msg2"})
	assert(msgSplit([]byte("msg3\r\nmsg"))).Equal([]string{"msg3"})
	assert(msgSplit([]byte("4\r\n"))).Equal([]string{"msg4"})
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

	assert(msgToEvent("# this is a message")).Equal(Event{
		EventType: Message,
		Message:   "# this is a message",
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

type RW struct {
	io.Reader
	io.Writer
}

func TestMsgRec(t *testing.T) {
	assert := assert.Make(t)
	s := RW{strings.NewReader("c 1/joystick\r\nu 1/joystick 1,234,874\r\nd 1/joystick\r\n"), ioutil.Discard}
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

func TestMsgSend(t *testing.T) {
	assert := assert.Make(t)

	var out bytes.Buffer

	s := RW{strings.NewReader("c 1/joystick\r\nu 1/joystick 1,234,874\r\nd 1/joystick\r\n"), &out}
	d := ConnectDock(s)

	assert(d.SendDockCommand('e')).NoError()
	assert(d.SendDockCommand('p', 1)).NoError()

	assert(out.String()).Equal("e\rp 1\r")

	out.Reset()
	assert(d.SetModuleData(1, Rainbow, 255, 255, 255)).NoError()
	assert(out.String()).Equal("s 1 255,255,255\r")
}
