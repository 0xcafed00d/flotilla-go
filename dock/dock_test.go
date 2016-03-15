package dock

import (
	"testing"

	"github.com/simulatedsimian/assert"
)

func TestMsgSplitter(t *testing.T) {
	assert := assert.Make(t)

	msgSplit := makeMessageSplitter()

	assert(msgSplit([]byte("msg1\rmsg2\r"))).Equal([]string{"msg1", "msg2"})
	assert(msgSplit([]byte("msg3\rmsg"))).Equal([]string{"msg3"})
	assert(msgSplit([]byte("4\r"))).Equal([]string{"msg4"})
}
