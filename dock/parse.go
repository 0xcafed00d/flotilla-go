package dock

import "bytes"

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

func msgToEvent(msg string) Event {
	return Event{}
}
