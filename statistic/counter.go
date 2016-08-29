package statistic

import (
	"sync/atomic"
)

// Counter is a performance counter.
type Counter struct {
	OldClientQPS    int64
	OldErrLogTotal  int64
	OldSlowLogTotal int64

	ClientConns  int64
	ClientQPS    int64
	ErrLogTotal  int64
	SlowLogTotal int64
}

// IncrClientConns is to increase client conns.
func (c *Counter) IncrClientConns() {
	atomic.AddInt64(&c.ClientConns, 1)
}

// DecrClientConns is to decrease client conns.
func (c *Counter) DecrClientConns() {
	atomic.AddInt64(&c.ClientConns, -1)
}

// IncrClientQPS is to increase client qps.
func (c *Counter) IncrClientQPS() {
	atomic.AddInt64(&c.ClientQPS, 1)
}

// IncrErrLogTotal is to increase err log total.
func (c *Counter) IncrErrLogTotal() {
	atomic.AddInt64(&c.ErrLogTotal, 1)
}

// IncrSlowLogTotal is to increase slow log total.
func (c *Counter) IncrSlowLogTotal() {
	atomic.AddInt64(&c.SlowLogTotal, 1)
}

// FlushCounter is to flush the count per second
func (c *Counter) FlushCounter() {
	atomic.StoreInt64(&c.OldClientQPS, c.ClientQPS)
	atomic.StoreInt64(&c.OldErrLogTotal, c.ErrLogTotal)
	atomic.StoreInt64(&c.OldSlowLogTotal, c.SlowLogTotal)
	atomic.StoreInt64(&c.ClientQPS, 0)
}
