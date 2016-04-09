package dock

import (
	"fmt"
	"testing"
	"time"

	"github.com/simulatedsimian/assert"
)

func TestPipe1(t *testing.T) {
	assert := assert.Make(t)

	buffer := make([]byte, 128)

	e1, e2, _ := NewPipe()

	assert(fmt.Fprintf(e1, "hello")).Equal(5, nil)
	assert(fmt.Fprintf(e1, "world")).Equal(5, nil)

	assert(e2.Read(buffer)).Equal(10, nil)
	assert(string(buffer[:10])).Equal("helloworld")
}

func TestEcho(t *testing.T) {
	assert := assert.Make(t)

	buffer := make([]byte, 128)

	e1, e2, _ := NewPipe()

	go func() {
		buffer := make([]byte, 128)
		for {
			n, err := e2.Read(buffer)
			if err != nil {
				break
			}
			e2.Write(buffer[:n])
		}
	}()

	assert(fmt.Fprintf(e1, "hello")).Equal(5, nil)
	assert(fmt.Fprintf(e1, "world")).Equal(5, nil)

	assert(e1.Read(buffer)).Equal(10, nil)
	assert(string(buffer[:10])).Equal("helloworld")
}

// use -race to test for race conditions
func TestRace(t *testing.T) {
	assert := assert.Make(t)

	e1, e2, pipe := NewPipe()

	assert(fmt.Fprintf(e1, "hello")).Equal(5, nil)
	assert(fmt.Fprintf(e2, "WORLD")).Equal(5, nil)

	go func() {
		buffer := make([]byte, 4)
		for {
			n, err := e2.Read(buffer)
			if err != nil {
				break
			}
			e2.Write(buffer[:n])
		}
		fmt.Println("go func 1 closed")
	}()

	go func() {
		buffer := make([]byte, 3)
		for {
			n, err := e1.Read(buffer)
			if err != nil {
				break
			}
			e1.Write(buffer[:n])
		}
		fmt.Println("go func 2 closed")
	}()

	time.Sleep(3 * time.Second)
	pipe.Close()
	time.Sleep(time.Second)

}
