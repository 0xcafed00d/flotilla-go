package dock

import "fmt"

type EventType int

const (
	Invalid             = -1
	Connected EventType = iota
	Disconnected
	Update
	Message
	Error
)

func (e EventType) String() string {
	switch e {
	case Connected:
		return "Connected"
	case Disconnected:
		return "Disconnected"
	case Update:
		return "Update"
	case Message:
		return "Message"
	case Error:
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
	if e.EventType == Error {
		return fmt.Sprintf("Event: [%v, %v]", e.EventType, e.Error)
	}
	if e.EventType == Message {
		return fmt.Sprintf("Event: [%v, %v]", e.EventType, e.Message)
	}
	return fmt.Sprintf("Event: [%v, %v, %v, %v]", e.EventType, e.ModuleType, e.Channel, e.Params)
}
