package dock

import "fmt"

type EventType int

const (
	EventInvalid EventType = iota
	EventConnected
	EventDisconnected
	EventUpdate
	EventMessage
	EventError
)

func (e EventType) String() string {
	switch e {
	case EventConnected:
		return "Connected"
	case EventDisconnected:
		return "Disconnected"
	case EventUpdate:
		return "Update"
	case EventMessage:
		return "Message"
	case EventError:
		return "Error"
	}
	return "invalid EventType"
}

type Event struct {
	EventType
	ModuleType
	Channel int
	Params  []int
	Message string
	Error   error
}

func (e Event) String() string {
	if e.EventType == EventError {
		return fmt.Sprintf("Event: [%v, %v]", e.EventType, e.Error)
	}
	if e.EventType == EventMessage {
		return fmt.Sprintf("Event: [%v, %v]", e.EventType, e.Message)
	}
	return fmt.Sprintf("Event: [%v, %v, %v, %v]", e.EventType, e.ModuleType, e.Channel, e.Params)
}
