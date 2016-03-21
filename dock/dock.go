package dock

import (
	"fmt"
	"io"
)

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

	if len(params) == 0 {
		_, err := fmt.Fprintf(d.port, "%c\r", command)
		return err
	}
	_, err := fmt.Fprintf(d.port, "%c %s\r", command, join(params, ","))
	return err
}

func (d *Dock) SetModuleData(port int, mtype ModuleType, params ...int) error {
	_, err := fmt.Fprintf(d.port, "s %d %s\r", port, join(params, ","))
	return err
}
