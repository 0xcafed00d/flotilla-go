package dock

import (
	"fmt"
	"io"
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

func TestValidation(t *testing.T) {
	assert := assert.Make(t)

	assert(validateParams(Dial, []int{1, 2, 3, 4})).HasError()
	assert(validateParams(Motor, []int{1})).NoError()

	assert(validateParams(Motor, []int{})).HasError()
	assert(validateParams(Motor, []int{1, 2})).HasError()

	assert(validateParams(Motor, []int{-64})).HasError()
	assert(validateParams(Motor, []int{64})).HasError()

	assert(validateParams(Number, []int{255, 255, 255, 255})).NoError()
	assert(validateParams(Number, []int{255, 255, 255, 255, 1})).NoError()
	assert(validateParams(Number, []int{255, 255, 255, 255, 1, 1})).NoError()
}

func TestMsgSplitter(t *testing.T) {
	assert := assert.Make(t)

	msgSplit := makeMessageSplitter([]byte{0x0d, 0x0a})

	assert(msgSplit([]byte("msg1\r\nmsg2\r\n"))).Equal([]string{"msg1", "msg2"})
	assert(msgSplit([]byte("msg3\r\nmsg"))).Equal([]string{"msg3"})
	assert(msgSplit([]byte("4\r\n"))).Equal([]string{"msg4"})
}

func TestMsgParser(t *testing.T) {
	assert := assert.Make(t)

	assert(msgToEvent("u 1/joystick 1,234,874")).Equal(Event{
		EventType:  EventUpdate,
		ModuleType: Joystick,
		Channel:    1,
		Params:     []int{1, 234, 874},
	})

	assert(msgToEvent("c 1/joystick")).Equal(Event{
		EventType:  EventConnected,
		ModuleType: Joystick,
		Channel:    1,
	})

	assert(msgToEvent("d 1/joystick")).Equal(Event{
		EventType:  EventDisconnected,
		ModuleType: Joystick,
		Channel:    1,
	})

	assert(msgToEvent("# this is a message")).Equal(Event{
		EventType: EventMessage,
		Message:   "# this is a message",
	})

	assert(msgToEvent("u 1/xxxx 1,234,874")).Equal(Event{
		EventType:  EventUpdate,
		ModuleType: Unknown,
		Channel:    1,
		Params:     []int{1, 234, 874},
	})

	assert(msgToEvent("p 1/xxxx 1,234,874")).Equal(Event{
		EventType:  EventInvalid,
		ModuleType: Unknown,
		Channel:    1,
		Params:     []int{1, 234, 874},
	})
}

func TestMsgRec(t *testing.T) {
	assert := assert.Make(t)

	e1, e2, _ := NewPipe()

	fmt.Fprintf(e1, "# message\r\nc 1/joystick\r\nu 1/joystick 1,234,874\r\nd 1/joystick\r\n")
	e1.Close()
	d := ConnectDock(e2)

	assert(<-d.Events).Equal(Event{
		EventType: EventMessage,
		Message:   "# message",
	})

	assert(<-d.Events).Equal(Event{
		EventType:  EventConnected,
		ModuleType: Joystick,
		Channel:    1,
	})

	assert(<-d.Events).Equal(Event{
		EventType:  EventUpdate,
		ModuleType: Joystick,
		Channel:    1,
		Params:     []int{1, 234, 874},
	})

	assert(<-d.Events).Equal(Event{
		EventType:  EventDisconnected,
		ModuleType: Joystick,
		Channel:    1,
	})

	assert(<-d.Events).Equal(Event{
		EventType: EventError,
		Error:     io.EOF,
	})
}

func TestMsgModuleType(t *testing.T) {
	assert := assert.Make(t)

	e1, e2, _ := NewPipe()

	fmt.Fprintf(e1, "c 1/joystick\r\n")
	d := ConnectDock(e2)

	<-d.Events
	assert(d.GetModuleType(2)).Equal(Unknown)
	assert(d.GetModuleType(1)).Equal(Joystick)

	fmt.Fprintf(e1, "d 1/joystick\r\n")

	<-d.Events
	assert(d.GetModuleType(1)).Equal(Unknown)
	e1.Close()
}

func TestMsgSend(t *testing.T) {
	assert := assert.Make(t)

	buffer := make([]byte, 128)

	e1, e2, _ := NewPipe()

	fmt.Fprintf(e1, "c 1/joystick\r\nu 1/joystick 1,234,874\r\nd 1/joystick\r\n")
	d := ConnectDock(e2)

	assert(d.SendDockCommand('e')).NoError()
	assert(d.SendDockCommand('p', 1)).NoError()

	n, err := e1.Read(buffer)
	assert(string(buffer[:n])).Equal("e\rp 1\r")
	assert(err).NoError()

	assert(d.SetModuleData(1, Rainbow, 255, 255, 255)).NoError()
	assert(d.SetModuleData(1, Rainbow, 255, 255, 255, 1)).HasError()

	n, err = e1.Read(buffer)
	assert(string(buffer[:n])).Equal("s 1 255,255,255\r")
	assert(err).NoError()
	e1.Close()
}
