package dock

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type splitterFunc func(input []byte) []string

func makeMessageSplitter(separator []byte) splitterFunc {
	buffer := []byte{}

	return func(input []byte) []string {
		buffer = append(buffer, input...)
		msgs := []string{}

		for {
			i := bytes.Index(buffer, separator)
			if i == -1 {
				break
			}
			msgs = append(msgs, string(buffer[:i]))
			buffer = buffer[i+2:]
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

	if evtype == '#' {
		return Event{EventType: Message, Message: msg}
	}

	event := Event{EventType: Invalid}

	if n >= 3 {
		switch evtype {
		case 'c':
			event.EventType = Connected
		case 'd':
			event.EventType = Disconnected
		case 'u':
			event.EventType = Update
		}

		event.ModuleType = FromString(module)
		event.Channel = port
		if n > 3 {
			event.Params = values[:n-3]
		}
	}
	return event
}

func msgToRequest(msg string) Request {

	parts := split(msg, " ,")
	if len(parts) == 0 {
		return Request{}
	}

	switch parts[0] {
	case "e":
		if len(parts) == 1 {
			return Request{RequestType: ReqEnquire}
		}
	case "r":
		if len(parts) == 1 {
			return Request{RequestType: ReqResetToBootloader}
		}
	case "d":
		if len(parts) == 1 {
			return Request{RequestType: ReqDebug}
		}
	case "v":
		if len(parts) == 1 {
			return Request{RequestType: ReqVersion}
		}

	case "p":
		if len(parts) == 2 {
			val, err := strconv.Atoi(parts[1])
			if err == nil && (val == 1 || val == 0) {
				return Request{RequestType: ReqPower, Params: []int{val}}
			}
		}
	case "":
		if len(parts) == 2 {
			val, err := strconv.Atoi(parts[1])
			if err == nil && (val == 1 || val == 0) {
				return Request{RequestType: ReqPower, Params: []int{val}}
			}
		}

	}

	return Request{}
}

func join(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return fmt.Sprint(a[0])
	}

	var b bytes.Buffer
	for i, v := range a {
		fmt.Fprint(&b, v)
		if i < (len(a) - 1) {
			fmt.Fprint(&b, sep)
		}
	}

	return b.String()
}

func choose(b bool, vt interface{}, vf interface{}) interface{} {
	if b {
		return vt
	}
	return vf
}

func split(s, charset string) []string {
	res := []string{}
	tokenStart := -1

	for i, r := range s {
		if strings.ContainsRune(charset, r) {
			if tokenStart != -1 {
				res = append(res, s[tokenStart:i])
				tokenStart = -1
			}
		} else {
			if tokenStart == -1 {
				tokenStart = i
			}
		}
	}
	if tokenStart != -1 {
		res = append(res, s[tokenStart:])
	}
	return res
}
