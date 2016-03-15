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

type Event struct {
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
