package dock

import (
	"fmt"
	"io"
	"sync"
)

type Simulator struct {
	port    io.ReadWriteCloser
	modules [8]ModuleType
	sync.Mutex
	Requests chan Request
}

func NewSimulator(port io.ReadWriteCloser) *Simulator {
	sim := Simulator{port: port}
	sim.Requests = make(chan Request, 128)
	go sim.reader()
	return &sim
}

func (s *Simulator) Close() {
	s.port.Close()
}

func (s *Simulator) Type(channel int) ModuleType {
	return s.modules[channel-1]
}

func (s *Simulator) Connect(modType ModuleType, channel int) error {
	if s.modules[channel-1] != Unknown {
		err := s.Disconnect(channel)
		if err != nil {
			return err
		}
	}
	s.modules[channel-1] = modType
	_, err := fmt.Fprintf(s.port, "c %d/%s\r\n", channel, modType)
	return err
}

func (s *Simulator) Disconnect(channel int) error {
	if s.modules[channel-1] == Unknown {
		return nil
	}
	_, err := fmt.Fprintf(s.port, "d %d/%s\r\n", channel, s.modules[channel-1])
	s.modules[channel-1] = Unknown
	return err
}

func (s *Simulator) NotifyAttachedModules() error {
	for i := range s.modules {
		_, err := fmt.Fprintf(s.port, "c %d/%s\r\n", i+1, s.modules[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Simulator) NotifyUpdate(modType ModuleType, channel int, params ...int) error {
	_, err := fmt.Fprintf(s.port, "u %d/%s %s\r\n", channel, s.modules[channel], join(params, ","))
	return err
}

func (s *Simulator) reader() {
	splitter := makeMessageSplitter([]byte{0x0d}) // cr
	buffer := make([]byte, 128)

	for {
		n, err := s.port.Read(buffer)
		if n > 0 {
			msgs := splitter(buffer[:n])
			for _, msg := range msgs {
				s.Requests <- msgToRequest(msg)
			}
		}

		if err != nil {
			s.Requests <- Request{RequestType: ReqError, Error: err}
			return
		}
	}
}
