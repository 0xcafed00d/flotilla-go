package dock

import (
	"bytes"
	"io"
)

func makeMessageSplitter() func(input []byte) []string {
	buffer := []byte{}

	return func(input []byte) []string {
		buffer = append(buffer, input...)
		msgs := []string{}

		for {
			i := bytes.IndexByte(buffer, '\r')
			if i == -1 {
				break
			}
			msgs = append(msgs, string(buffer[:i]))
			buffer = buffer[i+1:]
		}

		return msgs
	}
}

type EventType int

const (
	Conencted EventType = iota
	Disconnected
	Update
)

type Event struct {
	EventType
	ModuleType
	Port   int
	Params []int
}

type Dock struct {
	port   io.ReadWriteCloser
	Events <-chan Event
}

func ConnectDock(port io.ReadWriteCloser) *Dock {
	return &Dock{
		port:   port,
		Events: make(chan Event),
	}
}

func (d *Dock) SendCommand(port int, mtype ModuleType, params ...int) error {
	return nil
}
