package dock

import (
	"fmt"
	"io"
)

type Simulator struct {
	port    io.ReadWriteCloser
	modules [8]ModuleType
}

func MakeSimulator(port io.ReadWriteCloser) *Simulator {
	sim := Simulator{}
	return &sim
}

func (s *Simulator) Close() {
	s.port.Close()
}

func (s *Simulator) Type(index int) ModuleType {
	return s.modules[index]
}

func (s *Simulator) Connect(modType ModuleType, channel int) error {
	if s.modules[channel] != Unknown {
		err := s.Disconnect(channel)
		if err != nil {
			return err
		}
	}
	s.modules[channel] = modType
	_, err := fmt.Fprintf(s.port, "c %d/%s\r\n", channel, modType)
	return err
}

func (s *Simulator) Disconnect(channel int) error {
	if s.modules[channel] == Unknown {
		return nil
	}
	_, err := fmt.Fprintf(s.port, "d %d/%s\r\n", channel, s.modules[channel])
	s.modules[channel] = Unknown
	return err
}

func (s *Simulator) NotifyUpdate(modType ModuleType, channel int, params ...int) error {
	_, err := fmt.Fprintf(s.port, "u %d/%s %s\r\n", channel, s.modules[channel], join(params, ","))
	return err
}

func (s *Simulator) OnSet(f func(modType ModuleType, channel int, params ...int)) {
}

func (s *Simulator) Tick() error {
	return nil
}
