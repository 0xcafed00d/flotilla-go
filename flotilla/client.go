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
	ports           []io.ReadWriteCloser
	docks           []*dock.Dock
	connecteModules map[ModuleAddress]Updateable
	modules         []Updateable
	eventChan       chan Event
}

func structMembersToInterfaces(moduleStructPtr interface{}) (res []interface{}) {

	typeof := reflect.TypeOf(moduleStructPtr)

	if typeof.Kind() != reflect.Ptr && typeof.Elem().Kind() != reflect.Struct {
		panic("modules supplied to Client.AquireModules not a struct pointer")
	}

	fields := typeof.Elem().NumField()
	for i := 0; i < fields; i++ {
		iface := reflect.ValueOf(moduleStructPtr).Elem().Field(i).Addr().Interface()
		res = append(res, iface)
	}
	return
}

func (c *Client) AquireModules(moduleStructPtr interface{}) {

	modules := structMembersToInterfaces(moduleStructPtr)

	for _, m := range modules {
		module := reflect.ValueOf(m).Elem().FieldByName("Module")
		if module.IsValid() {
			if mod, ok := module.Addr().Interface().(*Module); ok {
				mod.client = c
				mod.address = ModuleAddress{-1, -1}
				if u, ok := module.Addr().Interface().(Updateable); ok {
					c.modules = append(c.modules, u)
				}
			}
		}
	}
}

func (c *Client) Run() error {
	for {
		err := c.processEvent()
		if err != nil {
			return err
		}
	}
}

func (c *Client) processEvent() error {
	ev := <-c.eventChan
	if ev.EventType == dock.EventError {
		return ev.Error
	}

	addr := ModuleAddress{dock: ev.dockIndex, channel: ev.Channel}

	if m, ok := c.connecteModules[addr]; ok {
		m.Update(ev)
		if !m.Connected() {
			delete(c.connecteModules, addr)
		}
		return nil
	}

	if ev.EventType == dock.EventConnected {
		for _, m := range c.modules {
			if !m.Connected() {
				m.Update(ev)
				if m.Connected() {
					c.connecteModules[addr] = m
					break
				}
			}
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
	client.connecteModules = make(map[ModuleAddress]Updateable)
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
