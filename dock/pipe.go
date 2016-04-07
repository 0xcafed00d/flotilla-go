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

type PipeEndpoint struct {
	pipe *Pipe
	buf  *dataBuffer
}

func (p *Pipe) GetEndpoint1() io.ReadWriteCloser {
	p.buf[0].condition = sync.NewCond(&p.buf[0].Mutex)
	return &PipeEndpoint{p, &p.buf[0]}
}

func (p *Pipe) GetEndpoint2() io.ReadWriteCloser {
	return &PipeEndpoint{p, &p.buf[1]}
}

func (pe *PipeEndpoint) Read(p []byte) (n int, err error) {
	pe.buf.Lock()
	defer pe.buf.Unlock()

	for len(pe.buf.data) == 0 {
		pe.buf.condition.Wait()
	}
	count := copy(p, pe.buf.data)

	return count, nil
}

// Write never blocks
func (pe *PipeEndpoint) Write(p []byte) (n int, err error) {
	if pe.pipe.Closed() {
		return 0, io.EOF
	}

	pe.buf.Lock()
	defer pe.buf.Unlock()
	pe.buf.data = append(pe.buf.data, p...)
	pe.buf.condition.Signal()
	return len(p), nil
}

func (pe *PipeEndpoint) Close() error {
	pe.pipe.Lock()
	defer pe.pipe.Unlock()
	pe.pipe.closed = true
	return nil
}
