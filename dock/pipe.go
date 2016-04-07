package dock

import (
	"io"
	"sync"
)

type dataBuffer struct {
	sync.Mutex
	data []byte
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
	return &PipeEndpoint{p, &p.buf[0]}
}

func (p *Pipe) GetEndpoint2() io.ReadWriteCloser {
	return &PipeEndpoint{p, &p.buf[1]}
}

func (pe *PipeEndpoint) Read(p []byte) (n int, err error) {
	for {
		pe.buf.Lock()

		sync.Cond

		if len(pe.buf.data) == 0 {
			pe.buf.Unlock() // unlock buffer mutex so writes can occur
			pe.waitReadable()
			continue
		}
		count := copy(p, pe.buf.data)

		pe.buf.Unlock()
		return count, nil
	}
}

func (pe *PipeEndpoint) waitReadable() {

}

// Write never blocks
func (pe *PipeEndpoint) Write(p []byte) (n int, err error) {
	if pe.pipe.Closed() {
		return 0, io.EOF
	}

	pe.buf.Lock()
	defer pe.buf.Unlock()
	pe.buf.data = append(pe.buf.data, p...)
	return len(p), nil
}

func (pe *PipeEndpoint) Close() error {
	pe.pipe.Lock()
	defer pe.pipe.Unlock()
	pe.pipe.closed = true
	return nil
}
