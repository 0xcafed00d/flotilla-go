package flotilla

import (
	"fmt"
	"io"

	"github.com/simulatedsimian/flotilla/dock"
	"github.com/tarm/serial"
)

// Event is a wrapper around a dock.Event that contains an additional Dock index
type Event struct {
	dock.Event
	dockIndex int
}

type Client struct {
	ports     []io.ReadWriteCloser
	docks     []*dock.Dock
	eventChan chan Event
}

func (c *Client) Close() {
	for _, p := range c.ports {
		p.Close()
	}
	c.docks = nil
}

func makeClient() *Client {
	client := Client{}
	client.eventChan = make(chan Event, 100)
	return &client
}

func ConnectToDock(serialport string) (*Client, error) {
	return ConnectToDocks(serialport)
}

func ConnectToDocks(serialports ...string) (*Client, error) {
	client := makeClient()
	for i, s := range serialports {

		serialcfg := serial.Config{Name: s, Baud: 115200}
		port, err := serial.OpenPort(&serialcfg)
		if err != nil {
			client.Close()
			return nil, fmt.Errorf("Failed to connect to Dock %d (%s): %v", i, s, err)
		}

		client.ports = append(client.ports, port)
		client.docks = append(client.docks, dock.ConnectDock(port))
	}

	// create a go routine for each dock that reads the event and gives it to the
	// common client event chan along with source dock index
	for i, d := range client.docks {
		go func(d *dock.Dock, index int) {
			ev := <-d.Events
			client.eventChan <- Event{ev, index}
		}(d, i)
	}

	return client, nil
}

func FindDocks() (*Client, error) {
	return nil, fmt.Errorf("Find Docks Not Implemented Yet")
}
