package flotilla

import (
	"fmt"
	"io"
	"reflect"
	"time"

	"github.com/simulatedsimian/flotilla/dock"
	"github.com/tarm/serial"
)

// Event is a wrapper around a dock.Event that contains an additional Dock index
type Event struct {
	dock.Event
	dockIndex int
}

type Client struct {
	ports            []io.ReadWriteCloser
	docks            []*dock.Dock
	connectedModules map[ModuleAddress]Module
	requestedModules []Module
	eventChan        chan Event
	ticker           *time.Ticker
	tickFunc         TickFunc
}

func structMembersToModules(moduleStructPtr interface{}) (res []Module) {

	typeof := reflect.TypeOf(moduleStructPtr)

	if typeof.Kind() != reflect.Ptr || typeof.Elem().Kind() != reflect.Struct {
		panic("modules supplied to Client.AquireModules not a struct pointer")
	}

	fields := typeof.Elem().NumField()
	for i := 0; i < fields; i++ {
		iface, ok := reflect.ValueOf(moduleStructPtr).Elem().Field(i).Addr().Interface().(Module)
		if ok {
			res = append(res, iface)
		}
	}
	return
}

func (c *Client) AquireModules(moduleStructPtr interface{}) {

	modules := structMembersToModules(moduleStructPtr)
	for _, m := range modules {
		m.Init(c, m.Type())
		c.requestedModules = append(c.requestedModules, m)
	}
}

func (c *Client) Run(tickTime time.Duration) error {
	time.Sleep(250 * time.Millisecond)
	for _, dock := range c.docks {
		dock.SendDockCommand('e')
	}
	time.Sleep(250 * time.Millisecond)

	c.ticker = time.NewTicker(tickTime)
	for {
		err := c.waitForEvent()
		if err != nil {
			return err
		}
	}
}

func (c *Client) waitForEvent() error {
	select {
	case ev := <-c.eventChan:
		return c.handleEvent(ev)
	case t := <-c.ticker.C:
		c.handleTick(t)
	}
	return nil
}

type TickFunc func(t time.Time)

func (c *Client) OnTick(tf TickFunc) {
	c.tickFunc = tf
}

func (c *Client) handleTick(t time.Time) {
	if c.tickFunc != nil {
		c.tickFunc(t)
	}

	for addr, mod := range c.connectedModules {
		if s, ok := mod.(Setable); ok {
			s.Set(c.docks[addr.dock])
		}
	}
}

func (c *Client) handleEvent(ev Event) error {

	if ev.EventType == dock.EventError {
		return ev.Error
	}

	addr := ModuleAddress{dock: ev.dockIndex, channel: ev.Channel}

	if m, ok := c.connectedModules[addr]; ok {
		m.Update(ev)
		if !m.Connected() {
			delete(c.connectedModules, addr)
		}
		return nil
	}

	if ev.EventType == dock.EventConnected {
		for _, m := range c.requestedModules {
			if !m.Connected() {
				m.Update(ev)
				if m.Connected() {
					c.connectedModules[addr] = m
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
	client.connectedModules = make(map[ModuleAddress]Module)
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
			for {
				ev := <-d.Events
				client.eventChan <- Event{ev, dockIndex}
				if ev.EventType == dock.EventError {
					return
				}
			}
		}(d, i)
	}

	return client, nil
}

func FindDocks() (*Client, error) {
	return nil, fmt.Errorf("Find Docks Not Implemented Yet")
}
