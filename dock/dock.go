package dock

import "io"

func messageSplitter() func(input []byte) []string {
	buffer := []byte{}

	return func(input []byte) []string {
		return nil
	}
}

type Event struct {
}

type Dock struct {
	port   io.ReadWriteCloser
	Events <-chan Event
}

func ConnectDock(port io.ReadWriteCloser) *Dock {
	return &FlotillaDock{
		port:   port,
		Events: make(chan Event),
	}
}
