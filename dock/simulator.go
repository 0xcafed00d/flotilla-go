package flotilla

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

func (s *Simulator) Connect(modType ModuleType, index int) error {
	if s.modules[index] != Unknown {
		err := s.Disconnect(index)
		if err != nil {
			return err
		}
	}
	s.modules[index] = modType
	_, err := fmt.Fprintf(s.port, "c %d/%s\r\n", index, modType)
	return err
}

func (s *Simulator) Disconnect(index int) error {
	if s.modules[index] == Unknown {
		return nil
	}
	_, err := fmt.Fprintf(s.port, "d %d/%s\r\n", index, s.modules[index])
	s.modules[index] = modType
	return err
}

func (s *Simulator) NotifyUpdate(modType ModuleType, index int, params ...int) {
	_, err := fmt.Fprintf(s.port, "u %d/%s %s\r\n", index, s.modules[index], join(params, ","))
	return err
}

func (s *Simulator) OnSet(f func(modType ModuleType, index int, params ...int)) {
}

func (s *Simulator) Tick() error {
	return nil
}
