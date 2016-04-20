package flotilla

import (
	"fmt"
	"io"
	"reflect"

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

func (c *Client) AquireModules(modules interface{}) {
	if reflect.TypeOf(modules).Kind() != reflect.Struct {
		panic("modules supplied to Client.AquireModules not a struct")
	}

	fields := reflect.TypeOf(modules).NumField()
	for i := 0; i < fields; i++ {
		fmt.Println(reflect.TypeOf(modules).Field(i).Type.Name())
	}
}

func (c *Client) Run() error {

	for {
		ev := <-c.eventChan
		if ev.EventType == dock.EventError {
			return ev.Error
		}
	}

	return nil
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
	if len(serialports) == 0 {
		return nil, fmt.Errorf("ConnectToDocks: No Serial Ports supplied")
	}

	ports := []io.ReadWriteCloser{}

	for i, s := range serialports {
		serialcfg := serial.Config{Name: s, Baud: 115200}
		port, err := serial.OpenPort(&serialcfg)
		if err != nil {
			for _, p := range ports {
				p.Close()
			}
			return nil, fmt.Errorf("Failed to connect to Dock %d (%s): %v", i, s, err)
		}
		ports = append(ports, port)
	}

	return ConnectToDocksRaw(ports...)
}

func ConnectToDocksRaw(ports ...io.ReadWriteCloser) (*Client, error) {
	client := makeClient()

	for _, port := range ports {
		client.ports = append(client.ports, port)
		client.docks = append(client.docks, dock.ConnectDock(port))
	}

	// create a go routine for each dock that reads the event and gives it to the
	// common client event chan along with source dock index
	for i, d := range client.docks {
		go func(d *dock.Dock, dockIndex int) {
			ev := <-d.Events
			client.eventChan <- Event{ev, dockIndex}
		}(d, i)
	}

	return client, nil
}

func FindDocks() (*Client, error) {
	return nil, fmt.Errorf("Find Docks Not Implemented Yet")
}
