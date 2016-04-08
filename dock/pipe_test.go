package dock

import (
	"fmt"
	"testing"

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
