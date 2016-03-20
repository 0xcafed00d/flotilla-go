package dock

import "io"

type Dock struct {
	port   io.ReadWriter
	Events chan Event
}

func ConnectDock(port io.ReadWriter) *Dock {
	dock := Dock{
		port:   port,
		Events: make(chan Event),
	}

	go dock.reader()

	return &dock
}

func (d *Dock) reader() {
	spliter := makeMessageSplitter()
	buffer := make([]byte, 256)

	for {
		n, err := d.port.Read(buffer)
		if n > 0 {
			msgs := spliter(buffer[:n])
			for _, msg := range msgs {
				d.Events <- msgToEvent(msg)
			}
		}

		if err != nil {
			d.Events <- Event{EventType: Error, Error: err}
			return
		}
	}
}

func (d *Dock) SendDockCommand(command rune, params ...int) error {
	return nil
}

func (d *Dock) SendModuleCommand(command rune, port int, mtype ModuleType, params ...int) error {
	return nil
}
