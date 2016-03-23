package dock

import (
	"fmt"
	"io"
	"sync"
)

type Dock struct {
	port   io.ReadWriter
	Events chan Event
	sync.RWMutex
	moduleTypes map[int]ModuleType
}

func ConnectDock(port io.ReadWriter) *Dock {
	dock := Dock{
		port:        port,
		Events:      make(chan Event, 128),
		moduleTypes: make(map[int]ModuleType),
	}

	go dock.reader()

	return &dock
}

func (d *Dock) handleEvent(ev Event) {
	if ev.EventType == Disconnected {
		d.RWMutex.Lock()
		d.moduleTypes[ev.Port] = Unknown
		d.RWMutex.Unlock()
	}

	if ev.EventType == Connected {
		d.RWMutex.Lock()
		d.moduleTypes[ev.Port] = ev.ModuleType
		d.RWMutex.Unlock()
	}

	d.Events <- ev
}

func (d *Dock) reader() {
	spliter := makeMessageSplitter()
	buffer := make([]byte, 128)

	for {
		n, err := d.port.Read(buffer)
		if n > 0 {
			msgs := spliter(buffer[:n])
			for _, msg := range msgs {
				d.handleEvent(msgToEvent(msg))
			}
		}

		if err != nil {
			d.handleEvent(Event{EventType: Error, Error: err})
			return
		}
	}
}

func (d *Dock) SendDockCommand(command rune, params ...int) error {

	if len(params) == 0 {
		_, err := fmt.Fprintf(d.port, "%c\r", command)
		return err
	}
	_, err := fmt.Fprintf(d.port, "%c %s\r", command, join(params, ","))
	return err
}

func (d *Dock) SetModuleData(port int, mtype ModuleType, params ...int) error {
	err := validateParams(mtype, params)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(d.port, "s %d %s\r", port, join(params, ","))
	return err
}

func (d *Dock) GetModuleType(port int) ModuleType {
	d.RWMutex.RLock()
	defer d.RWMutex.RUnlock()
	if mt, ok := d.moduleTypes[port]; ok {
		return mt
	}
	return Unknown
}
