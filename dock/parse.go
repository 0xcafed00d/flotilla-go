package dock

import (
	"bytes"
	"fmt"
)

func makeMessageSplitter() func(input []byte) []string {
	buffer := []byte{}

	return func(input []byte) []string {
		buffer = append(buffer, input...)
		msgs := []string{}

		for {
			i := bytes.IndexByte(buffer, '\r')
			if i == -1 {
				break
			}
			msgs = append(msgs, string(buffer[:i]))
			buffer = buffer[i+1:]
		}

		return msgs
	}
}

// u 1/module int,int,int,int....
func msgToEvent(msg string) Event {
	var evtype rune
	var port int
	var module string
	var values [8]int

	n, _ := fmt.Sscanf(msg, "%c %d/%s %d,%d,%d,%d,%d,%d,%d,%d",
		&evtype, &port, &module,
		&values[0], &values[1], &values[2], &values[3],
		&values[4], &values[5], &values[6], &values[7])

	//	fmt.Println(evtype, port, module, values, n, err)

	event := Event{}
	switch evtype {
	case 'c':
		event.EventType = Connected
	case 'd':
		event.EventType = Disconnected
	case 'u':
		event.EventType = Update
	default:
		event.EventType = Invalid
	}

	event.ModuleType = FromString(module)
	event.Port = port
	event.Params = values[:n-3]

	return event
}
