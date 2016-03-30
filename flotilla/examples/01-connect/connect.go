package main

import "github.com/simulatedsimian/flotilla/flotilla"

func main() {
	client, err := flotilla.ConnectToDock("/dev/ttyACM0")
	flotilla.ExitOnError(err)
	client.Close()
}
