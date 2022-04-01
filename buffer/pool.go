package buffer

import "sync"

// Pool represents a buffer pool.
type Pool struct {
	pool sync.Pool
}

// Get retrieves a buffer from the pool, creating one if necessary.
func (p *Pool) Get() *Buffer {
	buf, ok := p.pool.Get().(*Buffer)
	if ok {
		return buf
	}
	return &Buffer{}
}

// Put places the given buffer to the pool. Once this method is called, the
// buffer should be disposed.
func (p *Pool) Put(buf *Buffer) {
	buf.Reset()
	p.pool.Put(buf)
}
