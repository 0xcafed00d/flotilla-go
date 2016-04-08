package dock

import (
	"io"
	"sync"
)

type dataBuffer struct {
	sync.Mutex
	condition *sync.Cond
	data      []byte
}

type Pipe struct {
	buf [2]dataBuffer
	sync.Mutex
	closed bool
}

func (p *Pipe) Close() error {
	p.Lock()
	defer p.Unlock()
	p.closed = true
	return nil
}

func (p *Pipe) Closed() bool {
	p.Lock()
	defer p.Unlock()
	return p.closed
}

type pipeEndpoint struct {
	pipe *Pipe
	rbuf *dataBuffer
	wbuf *dataBuffer
}

func NewPipe() (io.ReadWriteCloser, io.ReadWriteCloser, *Pipe) {
	p := &Pipe{}

	p.buf[0].condition = sync.NewCond(&p.buf[0].Mutex)
	p.buf[1].condition = sync.NewCond(&p.buf[1].Mutex)

	return &pipeEndpoint{p, &p.buf[0], &p.buf[1]}, &pipeEndpoint{p, &p.buf[1], &p.buf[0]}, p
}

func (pe *pipeEndpoint) Read(p []byte) (n int, err error) {
	pe.rbuf.Lock()
	defer pe.rbuf.Unlock()

	for len(pe.rbuf.data) == 0 {
		pe.rbuf.condition.Wait()
	}
	count := copy(p, pe.rbuf.data)
	pe.rbuf.data = pe.rbuf.data[count:]

	return count, nil
}

// Write never blocks
func (pe *pipeEndpoint) Write(p []byte) (n int, err error) {
	if pe.pipe.Closed() {
		return 0, io.EOF
	}

	pe.wbuf.Lock()
	defer pe.wbuf.Unlock()
	pe.wbuf.data = append(pe.wbuf.data, p...)
	pe.wbuf.condition.Signal()
	return len(p), nil
}

func (pe *pipeEndpoint) Close() error {
	pe.pipe.Lock()
	defer pe.pipe.Unlock()
	pe.pipe.closed = true
	return nil
}
